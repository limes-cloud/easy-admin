package huawei

import (
	core "github.com/limeschool/easy-admin/server/tools/upload"
)

type config struct {
	Endpoint  string
	Location  string
	Bucket    string
	AccessKey string
	SecretKey string
}

func getConfig() *config {
	return &config{
		Endpoint:  core.Config.Endpoint,
		Location:  core.Config.Location,
		Bucket:    core.Config.Bucket,
		AccessKey: core.Config.AccessKey,
		SecretKey: core.Config.SecretKey,
	}
}
