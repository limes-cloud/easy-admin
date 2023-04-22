package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/internal/system/service"
	"github.com/limeschool/easy-admin/server/internal/system/types"
)

func PageNotice(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.PageNoticeRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}

	if list, total, err := service.PageNotice(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespList(total, list)
	}
}

func GetNotice(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.GetNoticeRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}

	if data, err := service.GetNotice(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(data)
	}
}

// GetUnReadNoticeNum 获取未读消息数量
func GetUnReadNoticeNum(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	if data, err := service.GetUnReadNoticeNum(ctx); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(data)
	}
}

func AddNotice(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.AddNoticeRequest{}
	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	// 调用实现
	if err := service.AddNotice(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func UpdateNotice(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.UpdateNoticeRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	// 调用实现
	if err := service.UpdateNotice(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func DeleteNotice(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.DeleteNoticeRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	// 调用实现
	if err := service.DeleteNotice(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}
