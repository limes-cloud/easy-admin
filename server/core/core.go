package core

import (
	"github.com/limeschool/easy-admin/server/core/casbin"
	"github.com/limeschool/easy-admin/server/core/cert"
	"github.com/limeschool/easy-admin/server/core/config"
	"github.com/limeschool/easy-admin/server/core/email"
	"github.com/limeschool/easy-admin/server/core/logger"
	"github.com/limeschool/easy-admin/server/core/orm"
	"github.com/limeschool/easy-admin/server/core/redis"
)

func Init() {
	// 初始化配置
	config.Init()
	// 初始化证书
	cert.Init()
	// 日志初始化
	logger.Init()
	// 数据库初始化
	orm.Init()
	// redis 初始化
	redis.Init()
	// casbin初始化
	casbin.Init()
	// email 初始化
	email.Init()
}
