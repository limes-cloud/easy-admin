package core

import (
	"github.com/limeschool/easy-admin/server/config"
	"github.com/limeschool/easy-admin/server/core/captcha"
	"github.com/limeschool/easy-admin/server/core/cert"
	"github.com/limeschool/easy-admin/server/core/email"
	"github.com/limeschool/easy-admin/server/core/enforcer"
	"github.com/limeschool/easy-admin/server/core/logger"
	"github.com/limeschool/easy-admin/server/core/orm"
	"github.com/limeschool/easy-admin/server/core/redis"
)

var (
	g = new(global)
)

type global struct {
	config   *config.Config
	logger   logger.Logger
	orm      orm.Orm
	redis    redis.Redis
	enforcer enforcer.Enforcer
	cert     cert.Cert
	email    email.Email
	captcha  captcha.Captcha
}

func initGlobal(config *config.Config, opts ...option) {
	g.config = config
	for _, opt := range opts {
		opt(g)
	}
}

func GlobalConfig() *config.Config {
	return g.config
}

func GlobalOrm() orm.Orm {
	return g.orm
}

func GlobalLogger() logger.Logger {
	return g.logger
}

func GlobalRedis() redis.Redis {
	return g.redis
}
