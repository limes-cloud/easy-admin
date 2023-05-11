package model

import (
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/tools"
	"gorm.io/gorm"
)

const (
	_orm = "system"
)

func DBName() string {
	return _orm
}

// dataMap 数据字典
var dataMap = map[string]string{
	"phone":   "手机号码",
	"email":   "电子邮箱",
	"keyword": "标志",
	"name":    "名称",
}

// Database 进行数据库选择
func database(ctx *core.Context) *gorm.DB {
	return ctx.Orm().GetDB(_orm).WithContext(ctx.SourceCtx())
}

func transferErr(err error) error {
	return tools.TransferErr(dataMap, err)
}
