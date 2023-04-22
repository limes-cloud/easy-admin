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
	"sync"
)

var (
	g    = new(global)
	once = sync.Once{}
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

// InitGlobal 初始化global实例
func InitGlobal(config *config.Config, opts ...options) {
	g.config = config
	once.Do(func() {
		for _, opt := range opts {
			opt(g)
		}
	})
}

// UpdateGlobal 更新实例信息
func UpdateGlobal(config *config.Config, opts ...options) {
	g.config = config
	for _, opt := range opts {
		opt(g)
	}
}

func GlobalConfig() *config.Config {
	return g.config
}
