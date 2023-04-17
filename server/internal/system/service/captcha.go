package service

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/captcha"
	"github.com/limeschool/easy-admin/server/core/metadata"
	"github.com/limeschool/easy-admin/server/internal/system/model"
	"time"
)

func ImageCaptcha(ctx *gin.Context) (map[string]any, error) {
	t := 30
	id, baseStr, err := captcha.NewImageCaptcha(ctx, 4, time.Duration(t)*time.Second)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"base64": baseStr,
		"expire": t,
		"id":     id,
	}, err
}

func EmailCaptcha(ctx *gin.Context) (any, error) {
	md, err := metadata.GetFormContext(ctx)
	if err != nil {
		return nil, err
	}

	// 获取用户邮箱信息
	user := model.User{}
	if err = user.OneByID(ctx, md.UserID); err != nil {
		return nil, err
	}

	// 发送邮箱验证码
	t := 3 * time.Minute
	id, err := captcha.NewEmailCaptcha(ctx, user.Email, 6, t)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"id":     id,
		"expire": t.Seconds(),
	}, nil

}
