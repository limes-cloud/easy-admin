package aliyun

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	core "github.com/limeschool/easy-admin/server/tools/upload"
)

type Driver struct{}

func NewDriver() core.Driver {
	return &Driver{}
}

func (d *Driver) Storage() (core.Storage, error) {
	bucket, err := ossBucket()
	if err != nil {
		return nil, err
	}

	store := Store{
		Bucket: bucket,
	}

	return core.NewStorage(&store), nil
}

func ossBucket() (*oss.Bucket, error) {
	conf := getConfig()

	if conf.Endpoint == "" || conf.AccessKeyID == "" || conf.AccessKeySecret == "" {
		return nil, core.ErrorConfigEmpty
	}

	client, err := oss.New(conf.Endpoint, conf.AccessKeyID, conf.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	return client.Bucket(conf.Bucket)
}

func (d Driver) Name() string {
	return "aliyun"
}
