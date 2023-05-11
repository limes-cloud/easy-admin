package model

import (
	"github.com/gogo/protobuf/proto"
	"github.com/limeschool/easy-admin/server/types"
	"gorm.io/gorm"
	"time"
)

// InitData 初始化角色数据
func (r *Role) InitData(db *gorm.DB) error {
	ins := []Role{
		{
			BaseModel: types.BaseModel{
				ID:        1,
				CreatedAt: time.Now().Unix(),
			}, ParentID: 0, Name: "超级管理员", Keyword: "superAdmin", Status: proto.Bool(true),
		},
	}
	return db.Create(ins).Error
}

func (m *Menu) InitData(db *gorm.DB) error {
	//ins := []Menu{
	//	{
	//		BaseModel: types.BaseModel{
	//			ID:        1,
	//			CreatedAt: time.Now().Unix(),
	//		}, ParentID: 0, Title: "超级管理员", Icon: "superAdmin", Path: "", Name: "", Type: "", Permission: "", Method: "", Component: "", Redirect: nil, Weight: nil, IsHidden: nil, IsCache: nil, IsHome: nil,
	//	},
	//}
	return nil

}

func (rm *RoleMenu) InitData(db *gorm.DB) error {
	return nil
}

func (t *Team) InitData(db *gorm.DB) error {
	return nil

}

func (u *User) InitData(db *gorm.DB) error {
	return nil

}

func (u *Notice) InitData(db *gorm.DB) error {
	return nil

}

func (u *NoticeUser) InitData(db *gorm.DB) error {
	return nil

}

func (u *LoginLog) InitData(db *gorm.DB) error {
	return nil

}
