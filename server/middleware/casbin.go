package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/consts"
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/errors"
)

func Casbin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := core.New(c)
		defer ctx.Release()

		path := ctx.FullPath()
		method := ctx.Request.Method

		// 白名单直接放行
		if ctx.Enforcer().IsWhitelist(method, path) {
			return
		}

		// 获取元数据,不存在则跳过权限验证
		md := ctx.Metadata()
		if md == nil {
			return
		}

		// 白名单或者超管直接放行
		// || ctx.Enforcer().IsBaseApi(method, path)  todo
		if md.RoleKey == consts.JwtSuperAdmin {
			return
		}

		// 权限判断
		if is, _ := ctx.Enforcer().Instance().Enforce(md.RoleKey, path, method); !is {
			ctx.RespError(errors.NotResourcePowerError)
			ctx.Abort()
			return
		}
	}
}
