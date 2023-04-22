package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/errors"
	"go.uber.org/zap"
)

// Recovery 全局panic 重启
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			ctx := core.New(c)
			defer ctx.Release()

			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				ctx.Logger().WithOptions(zap.AddCallerSkip(2)).Error(message,
					zap.Any("path", ctx.Request.URL.Path),
					zap.Any("method", ctx.Request.Method),
					zap.Any("user_agent", ctx.Request.Header.Get("User-Agent")),
					zap.Any("panic", PanicErr()),
					zap.Any("request", RequestInfo(ctx.Gin())),
				)
				ctx.RespError(errors.ServerError)
				ctx.Abort()
			}
		}()

		c.Next()
	}
}
