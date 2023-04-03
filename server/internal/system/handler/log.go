package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/response"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/internal/system/service"
	"github.com/limeschool/easy-admin/server/internal/system/types"
)

func LoginLog(ctx *gin.Context) {
	// 检验参数
	in := types.LoginLogRequest{}
	if ctx.ShouldBind(&in) != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}

	if list, total, err := service.PageLoginLog(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.List(ctx, in.Page, len(list), int(total), list)
	}
}
