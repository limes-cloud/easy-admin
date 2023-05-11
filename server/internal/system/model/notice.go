package model

import (
	"fmt"
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/types"
	"time"
)

type Notice struct {
	types.BaseModel
	Type       string `json:"type" gorm:"not null;size:128;comment:通知类型"`
	Title      string `json:"title" gorm:"not null;size:256;comment:通知标题"`
	Status     *bool  `json:"status" gorm:"not null;comment:通知状态"`
	Content    string `json:"content,omitempty" gorm:"not null;comment:通知内容"`
	Operator   string `json:"operator" gorm:"not null;size:128;comment:操作人员名称"`
	OperatorID int64  `json:"operator_id" gorm:"not null;size:32;comment:操作人员id"`
	ReadAt     int64  `json:"read_at" gorm:"-"`
}

func (u Notice) TableName() string {
	return "tb_system_notice"
}

// OneByID 通过id查询通知详细信息
func (u *Notice) OneByID(ctx *core.Context, id int64) error {
	md := ctx.Metadata()
	if md == nil {
		return errors.MetadataError
	}

	nu := NoticeUser{
		UserID:   md.UserID,
		NoticeID: id,
		ReadAt:   time.Now().Unix(),
	}
	_ = nu.Create(ctx)

	return transferErr(database(ctx).First(u, id).Error)
}

// UnReadNum 通过id查询通知详细信息
func (u *Notice) UnReadNum(ctx *core.Context) (int64, error) {
	md := ctx.Metadata()
	if md == nil {
		return 0, errors.MetadataError
	}

	nu := NoticeUser{}
	db := database(ctx).Model(u)

	join := fmt.Sprintf("left join %s on %s.notice_id=%s.id and %s.user_id=%d",
		nu.TableName(),
		nu.TableName(),
		u.TableName(),
		nu.TableName(),
		md.UserID,
	)

	total := int64(0)
	return total, db.Joins(join).Where(fmt.Sprintf("status=true and %v.user_id is null", nu.TableName())).Count(&total).Error
}

// Page 查询分页数据
func (u *Notice) Page(ctx *core.Context, options types.PageOptions) ([]*Notice, int64, error) {
	list, total := make([]*Notice, 0), int64(0)

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
	db = db.Select("id,type,title,status,operator,operator_id,created_at,updated_at,read_at")
	db = db.Offset((options.Page - 1) * options.PageSize).Limit(options.PageSize)

	return list, total, db.Find(&list).Error
}

// Create 创建通知信息
func (u *Notice) Create(ctx *core.Context) error {
	md := ctx.Metadata()
	if md == nil {
		return errors.MetadataError
	}

	u.Operator = md.Username
	u.OperatorID = md.UserID

	u.UpdatedAt = time.Now().Unix()
	return transferErr(database(ctx).Create(u).Error)
}

// Update 更新通知信息
func (u *Notice) Update(ctx *core.Context) error {
	md := ctx.Metadata()
	if md == nil {
		return errors.MetadataError
	}

	u.Operator = md.Username
	u.OperatorID = md.UserID
	// 执行更新
	return transferErr(database(ctx).Updates(&u).Error)
}

// DeleteByID 通过id删除通知信息
func (u *Notice) DeleteByID(ctx *core.Context, id int64) error {
	return transferErr(database(ctx).Delete(u, id).Error)
}
