package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core"
	"github.com/unrolled/secure"
)

func Ssl() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := core.New(c)
		defer ctx.Release()

		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":443",
		})
		err := middleware.Process(ctx.Writer, ctx.Request)
		if err != nil {
			ctx.RespError(err)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
