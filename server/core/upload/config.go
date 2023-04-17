package upload

import (
	"github.com/limeschool/easy-admin/server/config"
	"github.com/limeschool/easy-admin/server/global"
)

var Config config.Upload

func Init() {
	Config = global.Config.Upload
}
