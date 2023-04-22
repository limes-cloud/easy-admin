package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/internal/system/service"
	"github.com/limeschool/easy-admin/server/internal/system/types"
)

func AllTeam(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	if resp, err := service.AllTeam(ctx); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

func AddTeam(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.AddTeamRequest{}
	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.AddTeam(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func UpdateTeam(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.UpdateTeamRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.UpdateTeam(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func DeleteTeam(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.DeleteTeamRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	if err := service.DeleteTeam(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}
