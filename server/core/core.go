package core

import (
	"fmt"
	"github.com/limeschool/easy-admin/server/config"
	"github.com/limeschool/easy-admin/server/core/captcha"
	"github.com/limeschool/easy-admin/server/core/cert"
	"github.com/limeschool/easy-admin/server/core/email"
	"github.com/limeschool/easy-admin/server/core/enforcer"
	"github.com/limeschool/easy-admin/server/core/logger"
	"github.com/limeschool/easy-admin/server/core/orm"
	"github.com/limeschool/easy-admin/server/core/redis"
)

func Init() {
	// 初始化配置实例
	conf := config.New()
	// 初始化全局实例
	initGlobalInstance(conf)
	// 监听配置变更重新初始化实例
	conf.Watch(func(c *config.Config) {
		watchGlobalInstance(c)
	})
}

func initGlobalInstance(conf *config.Config) {
	// 新建配置实例
	logIns := logger.New(conf.Log, conf.Service.Name)
	// 新建数据库实例
	ormIns := orm.New(conf.Orm, logIns)
	// 新建缓存实例
	redisIns := redis.New(conf.Redis)
	// 新建鉴权实例
	efIns := enforcer.New(conf.Enforcer, ormIns.GetDB(conf.Enforcer.DB))
	// 新建证书实例
	certIns := cert.New(conf.Cert)
	// 新建邮件实例
	emailIns := email.New(conf.Email)
	// 新建验证码实例
	captchaIns := captcha.New(conf.Captcha, redisIns, emailIns)
	// 全局配置初始化
	InitGlobal(conf,
		WithLogger(logIns),
		WithOrm(ormIns),
		WithRedis(redisIns),
		WithEnforcer(efIns),
		WithCert(certIns),
		WithEmail(emailIns),
		WithCaptcha(captchaIns),
	)
}

func watchGlobalInstance(conf *config.Config) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("配置变更失败:%v", err)
		}
	}()

	UpdateGlobal(conf,
		WithLogger(logger.New(conf.Log, conf.Service.Name)),
		WithCert(cert.New(conf.Cert)),
		WithEmail(email.New(conf.Email)),
		WithCaptcha(captcha.New(conf.Captcha, g.redis, g.email)),
	)
}
