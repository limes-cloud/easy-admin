package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/casbin"
	"github.com/limeschool/easy-admin/server/core/metadata"
	"github.com/limeschool/easy-admin/server/core/response"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/global"
	"strings"
)

func Casbin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.FullPath()
		method := ctx.Request.Method

		// 白名单直接放行
		if casbin.IsWhiteList(strings.ToLower(method + ":" + path)) {
			return
		}

		// 获取元数据
		md, err := metadata.GetFormContext(ctx)
		if err != nil {
			response.Error(ctx, err)
			ctx.Abort()
			return
		}

		// 白名单或者超管直接放行
		if md.RoleKey == global.JwtSuperAdmin || casbin.IsBaseApi(ctx, method, path) {
			return
		}

		// 权限判断
		if is, _ := global.Casbin.Enforce(md.RoleKey, path, method); !is {
			response.Error(ctx, errors.NotResourcePowerError)
			ctx.Abort()
			return
		}
	}
}
