package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/internal/system/service"
	"github.com/limeschool/easy-admin/server/internal/system/types"
)

func UploadFile(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.UploadRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}

	if data, err := service.UploadFile(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(data)
	}
}
