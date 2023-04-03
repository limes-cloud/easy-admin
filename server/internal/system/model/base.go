package model

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/orm"
	"github.com/limeschool/easy-admin/server/core/trace"
	"github.com/limeschool/easy-admin/server/global"
	"gorm.io/gorm"
)

const (
	_Orm = "system"
)

// dataMap 数据字典
var dataMap = map[string]string{
	"phone":   "手机号码",
	"email":   "电子邮箱",
	"keyword": "标志",
	"name":    "名称",
}

// Database 进行数据库选择
func database(ctx *gin.Context) *gorm.DB {
	return global.Orm[_Orm].WithContext(trace.Context(ctx))
}

func transferErr(err error) error {
	return orm.TransferErr(dataMap, err)
}
