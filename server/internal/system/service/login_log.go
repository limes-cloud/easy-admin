package service

import (
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/internal/system/model"
	"github.com/limeschool/easy-admin/server/internal/system/types"
	"github.com/limeschool/easy-admin/server/tools/address"
	"github.com/limeschool/easy-admin/server/tools/ua"
	tps "github.com/limeschool/easy-admin/server/types"
)

func AddLoginLog(ctx *core.Context, phone string, err error) error {
	ip := ctx.ClientIP()
	userAgent := ctx.Request.Header.Get("User-Agent")
	info := ua.Parse(userAgent)
	desc := ""
	code := 0

	if err != nil {
		customErr, _ := err.(*tps.Response)
		code = customErr.Code
		desc = customErr.Msg
	}

	log := model.LoginLog{
		Phone:       phone,
		IP:          ip,
		Address:     address.New(ip).GetAddress(),
		Browser:     info.Name,
		Status:      err == nil,
		Description: desc,
		Code:        code,
		Device:      info.OS + " " + info.OSVersion,
	}
	return log.Create(ctx)
}

func PageLoginLog(ctx *core.Context, in *types.LoginLogRequest) ([]model.LoginLog, int64, error) {
	log := model.LoginLog{}
	return log.Page(ctx, tps.PageOptions{
		Page:     in.Page,
		PageSize: in.PageSize,
		Model:    in,
	})
}
