package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/global"
	"go.uber.org/ratelimit"
)

// RateLimit 服务限流 单位时间内 limit/s
func RateLimit() gin.HandlerFunc {
	rl := ratelimit.New(global.Config.Middleware.RateLimit.Limit)
	return func(ctx *gin.Context) {
		rl.Take()
	}
}
