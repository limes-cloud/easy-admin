package trace

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/global"
	"go.uber.org/zap"
)

const (
	Key = "trace-id"
)

// Logger 获取链路日志器
func Logger(ctx *gin.Context) *zap.Logger {
	traceId, _ := ctx.Value(Key).(string)
	return global.Logger.With(zap.Any(Key, traceId))
}

// GetID 获取链路日志id
func GetID(ctx *gin.Context) string {
	traceId, _ := ctx.Value(Key).(string)
	return traceId
}

// SetID 设置链路日志id
func SetID(ctx *gin.Context, id string) {
	ctx.Set(Key, id)
}

func Context(c *gin.Context) context.Context {
	ctx := context.TODO()
	for key, val := range c.Keys {
		ctx = context.WithValue(ctx, key, val)
	}
	return ctx
}
