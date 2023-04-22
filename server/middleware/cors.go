package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core"
)

// Cors 跨域相关的
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := core.New(c)
		defer ctx.Release()

		conf := ctx.Config().Middleware.Cors
		// 允许 Origin 字段中的域发送请求
		if conf.AllowOrigin == "*" {
			c.Writer.Header().Add("Access-Control-Allow-Origin", c.Request.Header.Get("origin"))
		} else {
			c.Writer.Header().Add("Access-Control-Allow-Origin", conf.AllowOrigin)
		}
		// 设置预验请求有效期为 86400 秒
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		// 设置允许请求的方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", conf.AllowMethod)
		// 设置允许请求的 Header
		c.Writer.Header().Set("Access-Control-Allow-Headers", conf.AllowHeader)
		// 设置拿到除基本字段外的其他字段
		c.Writer.Header().Set("Access-Control-Expose-Headers", conf.ExposeHeader)
		// 配置是否可以带认证信息
		c.Writer.Header().Set("Access-Control-Allow-Credentials", fmt.Sprint(conf.Credentials))
		// OPTIONS请求返回200
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
