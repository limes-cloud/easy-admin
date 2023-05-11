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

type option func(*global)

func withLogger(log logger.Logger) option {
	return func(g2 *global) {
		g2.logger = log
	}
}

func withOrm(orm orm.Orm) option {
	return func(g2 *global) {
		g2.orm = orm
	}
}

func withRedis(redis redis.Redis) option {
	return func(g2 *global) {
		g2.redis = redis
	}
}

func withEnforcer(enforcer enforcer.Enforcer) option {
	return func(g2 *global) {
		g2.enforcer = enforcer
	}
}

func withCert(cert cert.Cert) option {
	return func(g2 *global) {
		g2.cert = cert
	}
}

func withEmail(email email.Email) option {
	return func(g2 *global) {
		g2.email = email
	}
}

func withCaptcha(captcha captcha.Captcha) option {
	return func(g2 *global) {
		g2.captcha = captcha
	}
}
