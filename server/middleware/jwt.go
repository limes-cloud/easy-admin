package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/jwt"
	"github.com/limeschool/easy-admin/server/core/metadata"
	"github.com/limeschool/easy-admin/server/core/response"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/global"
)

func JwtAuth() gin.HandlerFunc {
	conf := global.Config.Middleware.Jwt

	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get(conf.Header)
		if token == "" {
			return
		}

		// 解析token
		claims, err := jwt.Parse(conf.Secret, token)
		if err != nil {
			response.Error(ctx, err)
			ctx.Abort()
			return
		}

		// 解析元数据
		md, err := metadata.Parse(claims)
		if err != nil {
			response.Error(ctx, err)
			ctx.Abort()
			return
		}

		// 是否开启唯一设备登陆
		if conf.Unique {
			if !jwt.Compare(ctx, md.UserID, token) {
				response.Error(ctx, errors.DulDeviceLoginError)
				ctx.Abort()
				return
			}
		}

		// 设置到上下文
		md.SetToContext(ctx)
	}
}
