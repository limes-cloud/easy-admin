package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/consts"
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/internal/system/service"
	"github.com/limeschool/easy-admin/server/internal/system/types"
	"github.com/limeschool/easy-admin/server/tools"
)

func AllRole(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	if resp, err := service.AllRole(ctx); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

func AddRole(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.AddRoleRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}

	if !tools.InList([]string{consts.ALLTEAM, consts.DOWNTEAM, consts.CURTEAM, consts.CUSTOM}, in.DataScope) {
		ctx.RespError(errors.ParamsError)
		return
	}

	if err := service.AddRole(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func UpdateRole(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.UpdateRoleRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}

	if !tools.InList([]string{consts.ALLTEAM, consts.DOWNTEAM, consts.CURTEAM, consts.CUSTOM}, in.DataScope) {
		ctx.RespError(errors.ParamsError)
		return
	}

	if err := service.UpdateRole(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func DeleteRole(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.DeleteRoleRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.DeleteRole(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}
