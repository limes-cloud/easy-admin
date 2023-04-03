package service

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/internal/system/types"
	"github.com/limeschool/easy-admin/server/tools"
	"github.com/limeschool/easy-admin/server/tools/upload"
	"github.com/limeschool/easy-admin/server/tools/upload/aliyun"
	"github.com/limeschool/easy-admin/server/tools/upload/huawei"
	"github.com/limeschool/easy-admin/server/tools/upload/minio"
	"github.com/limeschool/easy-admin/server/tools/upload/qiniu"
	"github.com/limeschool/easy-admin/server/tools/upload/tencent"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

// UploadFile 上传文件
func UploadFile(ctx *gin.Context, in *types.UploadRequest) (any, error) {

	form, err := ctx.MultipartForm()
	if err != nil {
		return nil, errors.InitUploadError
	}

	resp := make(map[string]any)

	// 提前循环检查是否存在不可上传的文件类型
	if len(upload.Config.AcceptTypes) != 0 {
		for _, files := range form.File {
			for _, file := range files {
				index := strings.LastIndex(file.Filename, ".")
				if index == -1 {
					return nil, errors.UploadTypeError
				}
				if !tools.InList(upload.Config.AcceptTypes, file.Filename[index+1:]) {
					return nil, errors.UploadTypeError
				}
			}
		}
	}

	// 上传文件
	for key, files := range form.File {
		var fileNames []string
		for _, file := range files {
			fileName, err := uploadFile(file, in.Dir)
			if err != nil {
				return nil, err
			}
			fileNames = append(fileNames, fileName)
		}

		if len(fileNames) == 1 {
			resp[key] = fileNames[0]
		} else if len(fileNames) > 1 {
			resp[key] = fileNames
		}
	}
	return resp, nil
}

func uploadFile(fileHeader *multipart.FileHeader, dir string) (string, error) {
	if upload.Config.MaxSize > 0 && int(fileHeader.Size/1024) > upload.Config.MaxSize {
		return "", errors.FileLimitMaxSizeError
	}

	tempFile, err := fileHeader.Open()
	if err != nil {
		return "", errors.OpenFileError
	}

	defer tempFile.Close()

	fileName := fileHeader.Filename
	// 判断是否对文件进行重命名
	if upload.Config.Rename {
		uid := fmt.Sprintf("%x", md5.Sum([]byte(uuid.New().String())))
		fileType := fileName[strings.LastIndex(fileName, "."):]
		fileName = uid + fileType
	}

	// 处理本地上传
	if upload.Config.DriveType == "local" {
		return uploadByLocal(upload.Config.LocalDir+"/"+dir, fileName, tempFile)
	}

	// 处理云端上传
	var drive upload.Driver
	switch upload.Config.DriveType {
	case "tencent":
		drive = tencent.NewDriver()
	case "qiniu":
		drive = qiniu.NewDriver()
	case "aliyun":
		drive = aliyun.NewDriver()
	case "huawei":
		drive = huawei.NewDriver()
	case "minio":
		drive = minio.NewDriver()
	default:
		return "", errors.UploadTypeNotSupportError
	}
	store, err := drive.Storage()
	if err != nil {
		return "", errors.NewF("初始化对象存储实例失败：%v", err.Error())
	}
	if err = store.Put(dir+"/"+fileName, tempFile); err != nil {
		return "", errors.NewF("上传文件到对象存储失败：%v", err.Error())
	}
	return fileName, nil
}

// uploadByLocal 上传文件到本地
func uploadByLocal(dir string, fileName string, file multipart.File) (string, error) {
	if is, err := tools.PathExists(dir); !is {
		if err != nil {
			return "", errors.NewF("获取文件上传目录信息失败：%v", err)
		}
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			return "", errors.NewF("创建文件上传目录失败：%v", err)
		}
	}

	path := dir + "/" + fileName
	saveFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return "", errors.NewF("创建上传文件失败：%v", err)
	}
	defer saveFile.Close()
	if _, err = io.Copy(saveFile, file); err != nil {
		return "", errors.NewF("保存上传文件失败：%v", err)
	}
	return path, nil
}
