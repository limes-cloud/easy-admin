package file

import (
	"errors"
	"github.com/limeschool/easy-admin/server/config"
	"io"

	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
)

type huawei struct {
	config config.File
	client *obs.ObsClient
}

func NewHuawei(conf *config.File) (Store, error) {
	if conf.Endpoint == "" || conf.Location == "" || conf.Bucket == "" || conf.AccessKey == "" || conf.SecretKey == "" {
		return nil, errors.New("upload config error")
	}

	client, err := obs.New(conf.AccessKey, conf.SecretKey, conf.Endpoint)
	if err != nil {
		return nil, err
	}
	return &huawei{
		client: client,
	}, nil
}

func (s *huawei) Put(key string, r io.Reader) error {
	input := &obs.PutObjectInput{}
	input.Bucket = s.config.Bucket
	input.Key = key
	input.Body = r

	_, err := s.client.PutObject(input)

	return err
}

func (s *huawei) PutFromLocal(key string, localPath string) error {
	input := &obs.PutFileInput{}
	input.Bucket = s.config.Bucket
	input.Key = key
	input.SourceFile = localPath

	_, err := s.client.PutFile(input)

	return err
}

func (s *huawei) Get(key string) (io.ReadCloser, error) {
	input := &obs.GetObjectInput{}
	input.Bucket = s.config.Bucket
	input.Key = key

	output, err := s.client.GetObject(input)
	if err != nil {
		return nil, err
	}

	return output.Body, err
}

func (s *huawei) Delete(key string) error {
	input := &obs.DeleteObjectInput{}
	input.Bucket = s.config.Bucket
	input.Key = key

	_, err := s.client.DeleteObject(input)

	return err
}

func (s *huawei) meta(key string) (*obs.GetObjectMetadataOutput, error) {
	input := &obs.GetObjectMetadataInput{}
	input.Bucket = s.config.Bucket
	input.Key = key

	output, err := s.client.GetObjectMetadata(input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (s *huawei) Size(key string) (int64, error) {
	input := &obs.GetObjectMetadataInput{}
	input.Bucket = s.config.Bucket
	input.Key = key

	output, err := s.client.GetObjectMetadata(input)
	if err != nil {
		return 0, err
	}

	return output.ContentLength, nil
}

func (s *huawei) Exists(key string) (bool, error) {
	_, err := s.meta(key)
	if err != nil {
		return false, err
	}

	return true, nil
}
