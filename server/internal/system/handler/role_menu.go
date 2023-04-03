package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/response"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/internal/system/service"
	"github.com/limeschool/easy-admin/server/internal/system/types"
)

func UpdateRoleMenu(ctx *gin.Context) {
	in := types.AddRoleMenuRequest{}
	if err := ctx.ShouldBindJSON(&in); err != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}
	if err := service.UpdateRoleMenu(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.Success(ctx)
	}
}

func RoleMenuIds(ctx *gin.Context) {
	in := types.RoleMenuIdsRequest{}
	if err := ctx.ShouldBind(&in); err != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}
	if ids, err := service.RoleMenuIds(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.Data(ctx, ids)
	}
}

func RoleMenu(ctx *gin.Context) {
	in := types.RoleMenuRequest{}
	if err := ctx.ShouldBind(&in); err != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}
	if ids, err := service.RoleMenu(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.Data(ctx, ids)
	}
}
