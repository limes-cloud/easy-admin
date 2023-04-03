package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/response"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/internal/system/service"
	"github.com/limeschool/easy-admin/server/internal/system/types"
)

func UploadFile(ctx *gin.Context) {
	// 检验参数
	in := types.UploadRequest{}
	if ctx.ShouldBind(&in) != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}

	if data, err := service.UploadFile(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.Data(ctx, data)
	}
}
