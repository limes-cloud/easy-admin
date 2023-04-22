package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/errors"
)

// IpLimit ip限流器 窗口时间内访问次数限制
func IpLimit() gin.HandlerFunc {
	prefix := "limit_ip_"

	return func(c *gin.Context) {
		ctx := core.New(c)
		defer ctx.Release()

		conf := ctx.Config().Middleware.IpLimit
		key := prefix + ctx.ClientIP()
		redis := ctx.Redis().GetRedis(conf.Cache)

		count, _ := redis.Get(ctx, key).Int()
		if count == 0 {
			redis.SetNX(ctx, key, 1, conf.Window)
		} else {
			if count > conf.Limit {
				ctx.RespError(errors.IpLimitError)
				ctx.Abort()
			} else {
				redis.Incr(ctx, key)
			}
		}
	}
}
