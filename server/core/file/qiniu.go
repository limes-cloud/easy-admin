package file

import (
	"bytes"
	"context"
	"errors"
	"github.com/limeschool/easy-admin/server/config"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	oss "github.com/qiniu/go-sdk/v7/storage"
)

type qiniu struct {
	config        config.File
	bucketManager *oss.BucketManager
}

func NewQiniu(conf config.File) (Store, error) {
	if conf.Bucket == "" || conf.AccessKey == "" || conf.SecretKey == "" {
		return nil, errors.New("upload config error")
	}
	mac := qbox.NewMac(conf.AccessKey, conf.SecretKey)
	return &qiniu{
		config: conf,
		bucketManager: oss.NewBucketManager(mac, &oss.Config{
			UseHTTPS: conf.UseSsl,
		}),
	}, nil
}

func (s *qiniu) Put(key string, r io.Reader) error {
	upToken := s.upToken()
	formUploader := s.uploader()
	ret := oss.PutRet{}

	buf := new(bytes.Buffer)
	size, err := io.Copy(buf, r)
	if err != nil {
		return err
	}

	err = formUploader.Put(context.Background(), &ret, upToken, key, buf, size, nil)
	if err != nil {
		return err
	}

	return err
}

func (s *qiniu) upToken() string {
	putPolicy := oss.PutPolicy{
		Scope: s.config.Bucket,
	}

	mac := qbox.NewMac(s.config.AccessKey, s.config.SecretKey)

	return putPolicy.UploadToken(mac)
}

func (s *qiniu) uploader() *oss.FormUploader {
	cfg := oss.Config{}
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	return oss.NewFormUploader(&cfg)
}

func (s *qiniu) PutFromLocal(key string, localPath string) error {
	upToken := s.upToken()
	formUploader := s.uploader()
	ret := oss.PutRet{}

	return formUploader.PutFile(context.Background(), &ret, upToken, key, localPath, nil)
}

func (s *qiniu) Get(key string) (io.ReadCloser, error) {
	url := s.getDownloadUrl(key)

	return s.getUrlContent(url)
}

func (s *qiniu) getDownloadUrl(key string) string {
	var url string

	if s.config.Private {
		mac := qbox.NewMac(s.config.AccessKey, s.config.SecretKey)
		deadline := time.Now().Add(time.Second * 3600).Unix() // 1小时有效期
		url = oss.MakePrivateURL(mac, s.config.Domain, key, deadline)
	} else {
		url = oss.MakePublicURL(s.config.Domain, key)
	}

	return url
}

func (s *qiniu) getUrlContent(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func (s *qiniu) meta(key string) (http.Header, error) {
	fi, err := s.bucketManager.Stat(s.config.Bucket, key)
	if err != nil {
		return nil, err
	}

	header := http.Header{}
	header.Set("Content-Length", strconv.FormatInt(fi.Fsize, 10))

	return header, nil
}

func (s *qiniu) Size(key string) (int64, error) {
	fi, err := s.bucketManager.Stat(s.config.Bucket, key)
	if err != nil {
		return 0, err
	}

	return fi.Fsize, nil
}

func (s *qiniu) Exists(key string) (bool, error) {
	_, err := s.meta(key)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *qiniu) Delete(key string) error {
	return s.bucketManager.Delete(s.config.Bucket, key)
}
