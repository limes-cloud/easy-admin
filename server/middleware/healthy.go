package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core"
)

// Healthy 健康检查
func Healthy() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := core.New(c)
		defer ctx.Release()

		ctx.RespSuccess()
	}
}
