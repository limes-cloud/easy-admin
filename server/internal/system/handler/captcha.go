package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/response"
	"github.com/limeschool/easy-admin/server/internal/system/service"
)

func Captcha(ctx *gin.Context) {
	if resp, err := service.ImageCaptcha(ctx); err != nil {
		response.Error(ctx, err)
	} else {
		response.Data(ctx, resp)
	}
}

func EmailCaptcha(ctx *gin.Context) {
	if resp, err := service.EmailCaptcha(ctx); err != nil {
		response.Error(ctx, err)
	} else {
		response.Data(ctx, resp)
	}
}
