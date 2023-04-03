package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/response"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/global"
	"github.com/limeschool/easy-admin/server/tools"
)

// IpLimit ip限流器 窗口时间内访问次数限制
func IpLimit() gin.HandlerFunc {
	conf := global.Config.Middleware.IpLimit
	prefix := "limit_ip_"
	redis := global.Redis

	return func(ctx *gin.Context) {
		key := prefix + tools.ClientIP(ctx)
		count, _ := redis.Get(ctx, key).Int()
		if count == 0 {
			redis.SetNX(ctx, key, 1, conf.Window)
		} else {
			if count > conf.Limit {
				response.Error(ctx, errors.IpLimitError)
			} else {
				redis.Incr(ctx, key)
			}
		}
	}
}
