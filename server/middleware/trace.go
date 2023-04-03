package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/limeschool/easy-admin/server/core/trace"
	"github.com/limeschool/easy-admin/server/global"
)

// Trace 设置链路id
func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.GetHeader(global.Config.Log.Header)
		if id == "" {
			id = uuid.New().String()
		}
		trace.SetID(ctx, id)
	}
}
