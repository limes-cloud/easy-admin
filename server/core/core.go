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
	initInstance(conf)

	// 监听配置变更重新初始化实例
	conf.Watch(func(c *config.Config) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("配置变更失败:%v", err)
			}
		}()
		initInstance(c)
	})
}

func initInstance(conf *config.Config) {
	// 日志
	loggerIns := logger.New(conf.Log, conf.Service.Name)
	// 数据库
	ormIns := orm.New(conf.Orm, loggerIns)
	// redis
	redisIns := redis.New(conf.Redis)
	// 权限校验器
	enforcerIns := enforcer.New(conf.Enforcer, ormIns)
	// 邮箱
	emailIns := email.New(conf.Email)
	// 验证码
	captchaIns := captcha.New(conf.Captcha, redisIns, emailIns)
	// 证书
	certIns := cert.New(conf.Cert)

	// 实例化到全局对象
	initGlobal(conf,
		withLogger(loggerIns),
		withOrm(ormIns),
		withRedis(redisIns),
		withEnforcer(enforcerIns),
		withEmail(emailIns),
		withCaptcha(captchaIns),
		withCert(certIns),
	)
}
