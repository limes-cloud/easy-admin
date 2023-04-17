package model

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/orm"
)

type LoginLog struct {
	orm.CreateModel
	Phone       string `json:"phone"`
	IP          string `json:"ip"`
	Address     string `json:"address"`
	Browser     string `json:"browser"`
	Device      string `json:"device"`
	Status      bool   `json:"status"`
	Code        int    `json:"code"`
	Description string `json:"description"`
}

func (u LoginLog) TableName() string {
	return "tb_system_login_log"
}

func (u *LoginLog) Create(ctx *gin.Context) error {
	return transferErr(database(ctx).Create(u).Error)
}

func (u *LoginLog) Page(ctx *gin.Context, options orm.PageOptions) ([]LoginLog, int64, error) {
	list, total := make([]LoginLog, 0), int64(0)

	db := database(ctx).Model(u)

	if options.Model != nil {
		db = orm.GormWhere(db, u.TableName(), options.Model)
	}

	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	db = db.Order("created_at desc").Offset((options.Page - 1) * options.PageSize).Limit(options.PageSize)

	return list, total, db.Find(&list).Error
}
