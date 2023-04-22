/**
 * @Author: 1280291001@qq.com
 * @Description:
 * @File: option
 * @Version: 1.0.0
 * @Date: 2023/4/19 00:29
 */

package core

import (
	"github.com/limeschool/easy-admin/server/core/captcha"
	"github.com/limeschool/easy-admin/server/core/cert"
	"github.com/limeschool/easy-admin/server/core/email"
	"github.com/limeschool/easy-admin/server/core/enforcer"
	"github.com/limeschool/easy-admin/server/core/logger"
	"github.com/limeschool/easy-admin/server/core/orm"
	"github.com/limeschool/easy-admin/server/core/redis"
)

type options func(*global)

func WithLogger(log logger.Logger) options {
	return func(g2 *global) {
		g2.logger = log
	}
}

func WithOrm(orm orm.Orm) options {
	return func(g2 *global) {
		g2.orm = orm
	}
}

func WithRedis(redis redis.Redis) options {
	return func(g2 *global) {
		g2.redis = redis
	}
}

func WithEnforcer(enforcer enforcer.Enforcer) options {
	return func(g2 *global) {
		g2.enforcer = enforcer
	}
}

func WithCert(cert cert.Cert) options {
	return func(g2 *global) {
		g2.cert = cert
	}
}

func WithEmail(email email.Email) options {
	return func(g2 *global) {
		g2.email = email
	}
}

func WithCaptcha(captcha captcha.Captcha) options {
	return func(g2 *global) {
		g2.captcha = captcha
	}
}
