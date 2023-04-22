package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/limeschool/easy-admin/server/core"
)

// Trace 设置链路id
func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := core.New(c)
		defer ctx.Release()

		id := ctx.GetHeader(ctx.Config().Log.Header)
		if id == "" {
			id = uuid.New().String()
		}
		ctx.SetTraceID(id)
	}
}
