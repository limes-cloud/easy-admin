package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/internal/system/service"
	"github.com/limeschool/easy-admin/server/internal/system/types"
)

func AllMenu(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	if resp, err := service.AllMenu(ctx); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

func AddMenu(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.AddMenuRequest{}
	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.AddMenu(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func UpdateMenu(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.UpdateMenuRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.UpdateMenu(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func DeleteMenu(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.DeleteMenuRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.DeleteMenu(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}
