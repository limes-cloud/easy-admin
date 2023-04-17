package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/response"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/internal/system/service"
	"github.com/limeschool/easy-admin/server/internal/system/types"
)

func PageNotice(ctx *gin.Context) {
	// 检验参数
	in := types.PageNoticeRequest{}
	if ctx.ShouldBind(&in) != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}

	if list, total, err := service.PageNotice(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.List(ctx, in.Page, len(list), int(total), list)
	}
}

func GetNotice(ctx *gin.Context) {
	// 检验参数
	in := types.GetNoticeRequest{}
	if ctx.ShouldBind(&in) != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}

	if data, err := service.GetNotice(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.Data(ctx, data)
	}
}

// GetUnReadNoticeNum 获取未读消息数量
func GetUnReadNoticeNum(ctx *gin.Context) {
	if data, err := service.GetUnReadNoticeNum(ctx); err != nil {
		response.Error(ctx, err)
	} else {
		response.Data(ctx, data)
	}
}

func AddNotice(ctx *gin.Context) {
	// 检验参数
	in := types.AddNoticeRequest{}
	if err := ctx.ShouldBindJSON(&in); err != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}
	// 调用实现
	if err := service.AddNotice(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.Success(ctx)
	}
}

func UpdateNotice(ctx *gin.Context) {
	// 检验参数
	in := types.UpdateNoticeRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}
	// 调用实现
	if err := service.UpdateNotice(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.Success(ctx)
	}
}

func DeleteNotice(ctx *gin.Context) {
	// 检验参数
	in := types.DeleteNoticeRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}
	// 调用实现
	if err := service.DeleteNotice(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.Success(ctx)
	}
}
