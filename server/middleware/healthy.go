package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/response"
	"time"
)

// Healthy 健康检查
func Healthy() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		time.Sleep(20 * time.Second)
		response.Success(ctx)
	}
}
