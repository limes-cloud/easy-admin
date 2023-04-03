package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/response"
	"github.com/unrolled/secure"
)

func Ssl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":443",
		})
		err := middleware.Process(ctx.Writer, ctx.Request)
		if err != nil {
			response.Error(ctx, err)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
