package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogo/protobuf/proto"
	"github.com/jinzhu/copier"
	"github.com/limeschool/easy-admin/server/core/metadata"
	"github.com/limeschool/easy-admin/server/core/orm"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/internal/system/model"
	"github.com/limeschool/easy-admin/server/internal/system/types"
	"gorm.io/gorm"
)

// PageNotice 获取通知分页信息
func PageNotice(ctx *gin.Context, in *types.PageNoticeRequest) ([]*model.Notice, int64, error) {
	md, err := metadata.GetFormContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	notice := model.Notice{}
	noticeUser := model.NoticeUser{}

	// 返回数据
	return notice.Page(ctx, orm.PageOptions{
		Page:     in.Page,
		PageSize: in.PageSize,
		Model:    in,
		Scopes: func(db *gorm.DB) *gorm.DB {
			join := fmt.Sprintf("left join %s on %s.notice_id=%s.id and %s.user_id=%d",
				noticeUser.TableName(),
				noticeUser.TableName(),
				notice.TableName(),
				noticeUser.TableName(),
				md.UserID,
			)
			db = db.Joins(join)

			if in.IsRead == nil {
				return db
			}

			if *in.IsRead == true {
				return db.Where(fmt.Sprintf("%s.user_id is not null", noticeUser.TableName()))
			} else {
				return db.Where(fmt.Sprintf("%s.user_id is null", noticeUser.TableName()))
			}
		},
	})
}

// GetNotice 查询通知信息
func GetNotice(ctx *gin.Context, in *types.GetNoticeRequest) (*model.Notice, error) {
	notice := &model.Notice{}
	return notice, notice.OneByID(ctx, in.ID)
}

// GetUnReadNoticeNum 获取未读消息数量
func GetUnReadNoticeNum(ctx *gin.Context) (int64, error) {
	notice := &model.Notice{}
	return notice.UnReadNum(ctx)
}

// AddNotice 新增通知信息
func AddNotice(ctx *gin.Context, in *types.AddNoticeRequest) error {
	notice := model.Notice{
		Status: proto.Bool(false),
	}
	if copier.Copy(&notice, in) != nil {
		return errors.AssignError
	}
	return notice.Create(ctx)
}

// UpdateNotice 更新通知信息
func UpdateNotice(ctx *gin.Context, in *types.UpdateNoticeRequest) error {
	notice := model.Notice{}
	if copier.Copy(&notice, in) != nil {
		return errors.AssignError
	}
	return notice.Update(ctx)
}

// DeleteNotice 删除通知信息
func DeleteNotice(ctx *gin.Context, in *types.DeleteNoticeRequest) error {
	notice := model.Notice{}
	return notice.DeleteByID(ctx, in.ID)
}
