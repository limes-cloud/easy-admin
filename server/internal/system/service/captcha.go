package service

import (
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/internal/system/model"
	"github.com/limeschool/easy-admin/server/internal/system/types"
	"go.uber.org/zap"
)

func ImageCaptcha(ctx *core.Context, in *types.CaptchaRequest) (any, error) {
	return ctx.ImageCaptcha(in.Name).New()
}

func EmailCaptcha(ctx *core.Context, in *types.CaptchaRequest) (any, error) {
	md := ctx.Metadata()
	if md == nil {
		return nil, errors.MetadataError
	}

	// 获取用户邮箱信息
	user := model.User{}
	if err := user.OneByID(ctx, md.UserID); err != nil {
		return nil, err
	}

	// 发送邮箱验证码
	res, err := ctx.EmailCaptcha(in.Name).New(user.Email)
	if err != nil {
		ctx.Logger().Error("验证码发送失败", zap.Error(err))
		return nil, errors.CaptchaSendError
	}
	return res, nil
}
