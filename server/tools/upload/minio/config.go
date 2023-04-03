package minio

import (
	core "github.com/limeschool/easy-admin/server/tools/upload"
)

type config struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	UseSSL    bool
	Bucket    string
}

func getConfig() *config {
	return &config{
		Endpoint:  core.Config.Endpoint,
		AccessKey: core.Config.AccessKey,
		SecretKey: core.Config.SecretKey,
		UseSSL:    core.Config.UseSsl,
		Bucket:    core.Config.Bucket,
	}
}
