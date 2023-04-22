package file

import (
	"context"
	"errors"
	"github.com/limeschool/easy-admin/server/config"
	"io"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type tencent struct {
	client *cos.Client
}

func NewTencent(conf config.File) (Store, error) {
	if conf.Endpoint == "" || conf.SecretID == "" || conf.SecretKey == "" {
		return nil, errors.New("upload config error")
	}
	u, err := url.Parse(conf.Endpoint)
	if err != nil {
		return nil, err
	}

	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  conf.SecretID,
			SecretKey: conf.SecretKey,
		},
	})

	return &tencent{client}, nil
}

func (s *tencent) Put(key string, r io.Reader) error {
	response, err := s.client.Object.Put(context.Background(), key, r, nil)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return httpError(response)
	}

	return nil
}

func (s *tencent) PutFromLocal(key string, localPath string) error {
	response, err := s.client.Object.PutFromFile(context.Background(), key, localPath, nil)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return httpError(response)
	}

	return nil
}

func (s *tencent) Get(key string) (io.ReadCloser, error) {
	resp, err := s.client.Object.Get(context.Background(), key, nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func (s *tencent) Delete(key string) error {
	_, err := s.client.Object.Delete(context.Background(), key)

	return err
}

func (s *tencent) Size(key string) (int64, error) {
	resp, err := s.client.Object.Head(context.Background(), key, nil)
	if err != nil {
		return 0, err
	}

	return resp.ContentLength, nil
}

func (s *tencent) Exists(key string) (bool, error) {
	return s.client.Object.IsExist(context.Background(), key)
}

func httpError(response *cos.Response) error {
	bytes, err := io.ReadAll(response.Body)
	defer func() {
		err = response.Body.Close()
	}()
	if err != nil {
		return err
	}

	return errors.New(string(bytes))
}
