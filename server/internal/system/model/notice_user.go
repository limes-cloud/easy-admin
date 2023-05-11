package model

import "github.com/limeschool/easy-admin/server/core"

type NoticeUser struct {
	NoticeID int64  `json:"notice_id" gorm:"not null;size:32;comment:通知id"`
	UserID   int64  `json:"user_id" gorm:"not null;size:32;comment:人员id"`
	ReadAt   int64  `json:"read_at" gorm:"not null;size:32;comment:阅读时间"`
	User     User   `gorm:"->"`
	Notice   Notice `gorm:"->"`
}

func (u NoticeUser) TableName() string {
	return "tb_system_notice_user"
}

// Create 创建阅读信息
func (u *NoticeUser) Create(ctx *core.Context) error {
	return transferErr(database(ctx).Create(u).Error)
}
