package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/response"
	"github.com/limeschool/easy-admin/server/errors"
)

func Resp404() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response.Error(ctx, errors.New("此接口不存在"))
	}
}
