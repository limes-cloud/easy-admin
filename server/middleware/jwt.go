package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/jwt"
	"github.com/limeschool/easy-admin/server/core/metadata"
	"github.com/limeschool/easy-admin/server/core/response"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/global"
	"strings"
)

func JwtAuth() gin.HandlerFunc {
	conf := global.Config.Middleware.Jwt

	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get(conf.Header)
		if token == "" {
			return
		}

		// 解析token
		claims, parseErr := jwt.Parse(conf.Secret, token)
		if claims == nil && parseErr != nil {
			response.Error(ctx, parseErr)
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

		// 设置到上下文
		md.SetToContext(ctx)

		// 检查是否过期，白名单内跳过
		url := strings.ToLower(ctx.Request.Method + ":" + ctx.FullPath())
		if errors.Is(parseErr, errors.TokenExpiredError) {
			if conf.Whitelist[url] {
				return
			} else {
				response.Error(ctx, parseErr)
				ctx.Abort()
				return
			}
		}

		// 是否开启唯一设备登陆
		if conf.Unique {
			if !jwt.Compare(ctx, md.UserID, token) {
				response.Error(ctx, errors.DulDeviceLoginError)
				ctx.Abort()
				return
			}
		}

	}
}
