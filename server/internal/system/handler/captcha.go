package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/internal/system/service"
	"github.com/limeschool/easy-admin/server/internal/system/types"
)

func Captcha(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.CaptchaRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}

	if resp, err := service.ImageCaptcha(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

func EmailCaptcha(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.CaptchaRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}

	if resp, err := service.EmailCaptcha(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}
