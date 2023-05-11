package model

import (
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/types"
)

type LoginLog struct {
	types.CreateModel
	Phone       string `json:"phone" gorm:"not null;type:varbinary(32);comment:手机号"`
	IP          string `json:"ip" gorm:"not null;type:varbinary(64);comment:登陆IP"`
	Address     string `json:"address" gorm:"not null;size:128;comment:登陆地址"`
	Browser     string `json:"browser" gorm:"not null;size:128;comment:登陆浏览器"`
	Device      string `json:"device" gorm:"not null;size:128;comment:登录设备"`
	Status      bool   `json:"status" gorm:"not null;comment:登录状态"`
	Code        int    `json:"code" gorm:"not null;size:32;comment:错误码"`
	Description string `json:"description" gorm:"not null;size:256;comment:登录备注"`
}

func (u LoginLog) TableName() string {
	return "tb_system_login_log"
}

func (u *LoginLog) Create(ctx *core.Context) error {
	return transferErr(database(ctx).Create(u).Error)
}

func (u *LoginLog) Page(ctx *core.Context, options types.PageOptions) ([]LoginLog, int64, error) {
	list, total := make([]LoginLog, 0), int64(0)

	db := database(ctx).Model(u)

	if options.Model != nil {
		db = ctx.Orm().GormWhere(db, u.TableName(), options.Model)
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
