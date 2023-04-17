package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/metadata"
	"github.com/limeschool/easy-admin/server/core/orm"
	"time"
)

type Notice struct {
	orm.BaseModel
	Type       string `json:"type"`
	Title      string `json:"title"`
	Status     *bool  `json:"status"`
	Content    string `json:"content,omitempty"`
	Operator   string `json:"operator"`
	OperatorID int64  `json:"operator_id"`
	ReadAt     int64  `json:"read_at" gorm:"->"`
}

func (u Notice) TableName() string {
	return "tb_system_notice"
}

// OneByID 通过id查询通知详细信息
func (u *Notice) OneByID(ctx *gin.Context, id int64) error {
	md, err := metadata.GetFormContext(ctx)
	if err != nil {
		return err
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
func (u *Notice) UnReadNum(ctx *gin.Context) (int64, error) {
	md, err := metadata.GetFormContext(ctx)
	if err != nil {
		return 0, err
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
func (u *Notice) Page(ctx *gin.Context, options orm.PageOptions) ([]*Notice, int64, error) {
	list, total := make([]*Notice, 0), int64(0)

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
	db = db.Select("id,type,title,status,operator,operator_id,created_at,updated_at,read_at")
	db = db.Offset((options.Page - 1) * options.PageSize).Limit(options.PageSize)

	return list, total, db.Find(&list).Error
}

// Create 创建通知信息
func (u *Notice) Create(ctx *gin.Context) error {
	md, err := metadata.GetFormContext(ctx)
	if err != nil {
		return err
	}
	u.Operator = md.Username
	u.OperatorID = md.UserID

	u.UpdatedAt = time.Now().Unix()
	return transferErr(database(ctx).Create(u).Error)
}

// Update 更新通知信息
func (u *Notice) Update(ctx *gin.Context) error {
	md, err := metadata.GetFormContext(ctx)
	if err != nil {
		return err
	}
	u.Operator = md.Username
	u.OperatorID = md.UserID
	// 执行更新
	return transferErr(database(ctx).Updates(&u).Error)
}

// DeleteByID 通过id删除通知信息
func (u *Notice) DeleteByID(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Delete(u, id).Error)
}
