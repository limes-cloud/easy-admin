package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/response"
	"github.com/limeschool/easy-admin/server/core/runtime"
	"github.com/limeschool/easy-admin/server/core/trace"
	"github.com/limeschool/easy-admin/server/errors"
	"go.uber.org/zap"
)

// Recovery 全局panic 重启
func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				trace.Logger(ctx).WithOptions(zap.AddCallerSkip(2)).Error(message,
					zap.Any("panic", runtime.PanicErr()),
					zap.Any("request", runtime.RequestInfo(ctx)),
				)
				response.Error(ctx, errors.ServerError)
				ctx.Abort()
			}
		}()

		ctx.Next()
	}
}
