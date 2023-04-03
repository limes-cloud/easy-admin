package tencent

import (
	core "github.com/limeschool/easy-admin/server/tools/upload"
)

type config struct {
	Url       string
	SecretId  string
	SecretKey string
}

func getConfig() *config {
	return &config{
		Url:       core.Config.Endpoint,
		SecretId:  core.Config.SecretID,
		SecretKey: core.Config.SecretKey,
	}
}
