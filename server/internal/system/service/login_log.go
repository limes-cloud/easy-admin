package service

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/orm"
	"github.com/limeschool/easy-admin/server/internal/system/model"
	"github.com/limeschool/easy-admin/server/internal/system/types"
	"github.com/limeschool/easy-admin/server/tools"
	"github.com/limeschool/easy-admin/server/tools/address"
	"github.com/limeschool/easy-admin/server/tools/ua"
	tps "github.com/limeschool/easy-admin/server/types"
)

func AddLoginLog(ctx *gin.Context, phone string, err error) error {
	ip := tools.ClientIP(ctx)
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
		Address:     address.GetAddress(ip),
		Browser:     info.Name,
		Status:      err == nil,
		Description: desc,
		Code:        code,
		Device:      info.OS + " " + info.OSVersion,
	}
	return log.Create(ctx)
}

func PageLoginLog(ctx *gin.Context, in *types.LoginLogRequest) ([]model.LoginLog, int64, error) {
	log := model.LoginLog{}
	return log.Page(ctx, orm.PageOptions{
		Page:     in.Page,
		PageSize: in.PageSize,
		Model:    in,
	})
}
