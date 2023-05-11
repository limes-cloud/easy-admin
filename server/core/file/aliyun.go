package file

import (
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/limeschool/easy-admin/server/config"
	"io"
	"strconv"
)

type aliyun struct {
	Bucket *oss.Bucket
}

func NewAliyun(conf *config.File) (Store, error) {
	if conf.Endpoint == "" || conf.AccessKeyID == "" || conf.AccessKeySecret == "" {
		return nil, errors.New("upload config error")
	}

	client, err := oss.New(conf.Endpoint, conf.AccessKeyID, conf.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	bucket, err := client.Bucket(conf.Bucket)
	if err != nil {
		return nil, err
	}
	return &aliyun{
		Bucket: bucket,
	}, nil
}

func (s *aliyun) Put(key string, r io.Reader) error {
	return s.Bucket.PutObject(key, r)
}

func (s *aliyun) PutFromLocal(key string, localPath string) error {
	return s.Bucket.PutObjectFromFile(key, localPath)
}

func (s *aliyun) Get(key string) (io.ReadCloser, error) {
	return s.Bucket.GetObject(key)
}

func (s *aliyun) Delete(key string) error {
	return s.Bucket.DeleteObject(key)
}

func (s *aliyun) Size(key string) (int64, error) {
	header, err := s.Bucket.GetObjectDetailedMeta(key)
	if err != nil {
		return 0, err
	}

	length := header.Get("Content-Length")
	return strconv.ParseInt(length, 10, 64)
}

func (s *aliyun) Exists(key string) (bool, error) {
	return s.Bucket.IsObjectExist(key)
}
