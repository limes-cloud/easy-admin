package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/consts"
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/internal/system/model"
)

func Enforce() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := core.New(c)
		defer ctx.Release()

		path := ctx.FullPath()
		method := ctx.Request.Method

		// 获取元数据,不存在则跳过权限验证
		md := ctx.Metadata()
		if md == nil {
			return
		}

		// 超级管理放行
		if md.RoleKey == consts.JwtSuperAdmin {
			return
		}

		// 基础api放行
		menu := model.Menu{}
		if menu.IsBaseApiPath(ctx, method, path) {
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
