package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/internal/system/service"
	"github.com/limeschool/easy-admin/server/internal/system/types"
)

func AllFile(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.AllFileRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}

	if list, err := service.AllFile(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(list)
	}
}

func AddFile(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.AddFileRequest{}
	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	// 调用实现
	if err := service.AddFile(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func UpdateFile(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.UpdateFileRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	// 调用实现
	if err := service.UpdateFile(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func DeleteFile(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.DeleteFileRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	// 调用实现
	if err := service.DeleteFile(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}
