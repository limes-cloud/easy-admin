package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/response"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/global"
	"github.com/limeschool/easy-admin/server/internal/system/service"
	"github.com/limeschool/easy-admin/server/internal/system/types"
	"github.com/limeschool/easy-admin/server/tools"
)

func AllRole(ctx *gin.Context) {
	if resp, err := service.AllRole(ctx); err != nil {
		response.Error(ctx, err)
	} else {
		response.Data(ctx, resp)
	}
}

func AddRole(ctx *gin.Context) {
	in := types.AddRoleRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}

	if !tools.InList([]string{global.ALLTEAM, global.DOWNTEAM, global.CURTEAM, global.CUSTOM}, in.DataScope) {
		response.Error(ctx, errors.ParamsError)
		return
	}

	if err := service.AddRole(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.Success(ctx)
	}
}

func UpdateRole(ctx *gin.Context) {
	in := types.UpdateRoleRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}

	if !tools.InList([]string{global.ALLTEAM, global.DOWNTEAM, global.CURTEAM, global.CUSTOM}, in.DataScope) {
		response.Error(ctx, errors.ParamsError)
		return
	}

	if err := service.UpdateRole(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.Success(ctx)
	}
}

func DeleteRole(ctx *gin.Context) {
	in := types.DeleteRoleRequest{}
	if ctx.ShouldBind(&in) != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}
	if err := service.DeleteRole(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.Success(ctx)
	}
}
