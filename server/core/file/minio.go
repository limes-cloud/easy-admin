package file

import (
	"bytes"
	"context"
	"errors"
	"github.com/limeschool/easy-admin/server/config"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"net/http"
	"os"

	m "github.com/minio/minio-go/v7"
)

type minio struct {
	client *m.Client
	config *config.File
}

func NewMinio(conf *config.File) (Store, error) {
	if conf.Bucket == "" || conf.AccessKey == "" || conf.SecretKey == "" || conf.Endpoint == "" {
		return nil, errors.New("upload config error")
	}

	minioClient, err := m.New(conf.Endpoint, &m.Options{
		Creds:  credentials.NewStaticV4(conf.AccessKey, conf.SecretKey, ""),
		Secure: conf.UseSsl,
	})
	if err != nil {
		return nil, err
	}
	return &minio{
		client: minioClient,
		config: conf,
	}, nil
}

func (s *minio) Put(key string, r io.Reader) error {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r)
	if err != nil {
		return err
	}

	_, err = s.client.PutObject(context.Background(), s.config.Bucket, key, buf, int64(buf.Len()), m.PutObjectOptions{ContentType: "application/octet-stream"})

	return err
}

func (s *minio) putFile(key string, f *os.File) error {
	fi, err := f.Stat()
	if err != nil {
		return err
	}
	size := fi.Size()

	_, err = s.client.PutObject(context.Background(), s.config.Bucket, key, f, size, m.PutObjectOptions{ContentType: "application/octet-stream"})

	return err
}

func (s *minio) PutFromLocal(key string, localPath string) error {
	f, err := os.Open(localPath)
	if err != nil {
		return err
	}

	return s.putFile(key, f)
}

func (s *minio) getObject(key string) (*m.Object, error) {
	return s.client.GetObject(context.Background(), s.config.Bucket, key, m.GetObjectOptions{})
}

func (s *minio) Get(key string) (io.ReadCloser, error) {
	object, err := s.getObject(key)
	if err != nil {
		return nil, err
	}

	return object, nil
}

func (s *minio) stat(key string) (m.ObjectInfo, error) {
	return s.client.StatObject(context.Background(), s.config.Bucket, key, m.StatObjectOptions{})
}

func (s *minio) Size(key string) (int64, error) {
	info, err := s.stat(key)
	if err != nil {
		return 0, err
	}

	return info.Size, nil
}

func (s *minio) Exists(key string) (bool, error) {
	_, err := s.stat(key)
	if err != nil {
		if errResponse, ok := err.(m.ErrorResponse); ok {
			if errResponse.StatusCode == http.StatusNotFound {
				return false, nil
			}
		}

		return false, err
	}

	return true, nil
}

func (s *minio) Delete(key string) error {
	return s.client.RemoveObject(context.Background(), s.config.Bucket, key, m.RemoveObjectOptions{})
}
