package service

import (
	"github.com/limeschool/easy-admin/server/core"
	"go.uber.org/zap"

	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/internal/system/types"
)

// UploadFile 上传文件
func UploadFile(ctx *core.Context, in *types.UploadRequest) (any, error) {
	// 检验是否是存在上传的dir
	file, err := ctx.File(in.Dir)
	if err != nil {
		ctx.Logger().Error("文件上传失败", zap.Error(err))
		return nil, errors.InitUploadError
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return nil, errors.InitUploadError
	}

	resp := make(map[string]any)

	// 提前循环检查是否存在不可上传的文件类型
	for _, files := range form.File {
		for _, item := range files {
			if err = file.CheckType(item.Filename); err != nil {
				return nil, errors.UploadTypeNotSupportError
			}
			if err = file.CheckSize(item.Size); err != nil {
				return nil, errors.UploadLimitMaxSizeError
			}
		}
	}

	// 上传文件
	for key, files := range form.File {
		var fileNames []string
		for _, item := range files {
			temp, err := item.Open()
			if err != nil {
				return nil, errors.OpenFileError
			}
			fileName, err := file.Upload(item.Filename, temp)
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
