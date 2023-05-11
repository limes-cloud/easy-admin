package service

import (
	"github.com/limeschool/easy-admin/server/config"
	"github.com/limeschool/easy-admin/server/core"
)

// Config 获取系统配置信息
func Config(ctx *core.Context) *config.Service {
	// 检查系统数据库信息
	return ctx.Config().Service

}
