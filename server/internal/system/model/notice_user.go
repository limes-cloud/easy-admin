package model

import (
	"github.com/gin-gonic/gin"
)

type NoticeUser struct {
	NoticeID int64 `json:"notice_id"`
	UserID   int64 `json:"user_id"`
	ReadAt   int64 `json:"read_at"`
}

func (u NoticeUser) TableName() string {
	return "tb_system_notice_user"
}

// Create 创建阅读信息
func (u *NoticeUser) Create(ctx *gin.Context) error {
	return transferErr(database(ctx).Create(u).Error)
}
