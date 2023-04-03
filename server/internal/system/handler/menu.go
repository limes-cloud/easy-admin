package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/response"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/internal/system/service"
	"github.com/limeschool/easy-admin/server/internal/system/types"
)

func AllMenu(ctx *gin.Context) {
	if resp, err := service.AllMenu(ctx); err != nil {
		response.Error(ctx, err)
	} else {
		response.Data(ctx, resp)
	}
}

func AddMenu(ctx *gin.Context) {
	in := types.AddMenuRequest{}
	if err := ctx.ShouldBindJSON(&in); err != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}
	if err := service.AddMenu(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.Success(ctx)
	}
}

func UpdateMenu(ctx *gin.Context) {
	in := types.UpdateMenuRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}
	if err := service.UpdateMenu(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.Success(ctx)
	}
}

func DeleteMenu(ctx *gin.Context) {
	in := types.DeleteMenuRequest{}
	if ctx.ShouldBind(&in) != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}
	if err := service.DeleteMenu(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.Success(ctx)
	}
}
