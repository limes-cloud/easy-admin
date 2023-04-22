package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/internal/system/service"
	"github.com/limeschool/easy-admin/server/internal/system/types"
)

func UpdateRoleMenu(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.AddRoleMenuRequest{}
	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.UpdateRoleMenu(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func RoleMenuIds(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.RoleMenuIdsRequest{}
	if err := ctx.ShouldBind(&in); err != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if ids, err := service.RoleMenuIds(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(ids)
	}
}

func RoleMenu(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.RoleMenuRequest{}
	if err := ctx.ShouldBind(&in); err != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if ids, err := service.RoleMenu(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(ids)
	}
}
