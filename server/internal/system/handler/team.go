package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/response"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/internal/system/service"
	"github.com/limeschool/easy-admin/server/internal/system/types"
)

func AllTeam(ctx *gin.Context) {
	if resp, err := service.AllTeam(ctx); err != nil {
		response.Error(ctx, err)
	} else {
		response.Data(ctx, resp)
	}
}

func AddTeam(ctx *gin.Context) {
	in := types.AddTeamRequest{}
	if err := ctx.ShouldBindJSON(&in); err != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}
	if err := service.AddTeam(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.Success(ctx)
	}
}

func UpdateTeam(ctx *gin.Context) {
	in := types.UpdateTeamRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}
	if err := service.UpdateTeam(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.Success(ctx)
	}
}

func DeleteTeam(ctx *gin.Context) {
	in := types.DeleteTeamRequest{}
	if ctx.ShouldBind(&in) != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}
	if err := service.DeleteTeam(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.Success(ctx)
	}
}
