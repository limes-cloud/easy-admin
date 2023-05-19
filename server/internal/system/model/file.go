package model

import (
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/types"
)

type File struct {
	types.DeleteModel
	UserID   int64  `json:"user_id" gorm:"not null;size:32;comment:用户id"`
	Src      string `json:"src" gorm:"not null;size:256;comment:文件地址"`
	Name     string `json:"name" gorm:"not null;size:256;comment:文件名称"`
	IsDir    *bool  `json:"is_dir" gorm:"not null;comment:是否为文件夹"`
	ParentID int64  `json:"parent_id" gorm:"not null;size:32;comment:父文件夹id"`
	Size     int64  `json:"size" gorm:"not null;size:32;comment:文件大小"`
	Suffix   string `json:"suffix" gorm:"not null;size:32;comment:文件后缀"`
	Mime     string `json:"mime" gorm:"not null;size:256;comment:文件mime"`
	User     User   `json:"user" gorm:"->"`
}

func (u *File) TableName() string {
	return "tb_system_file"
}

// All 查询分页数据
func (u *File) All(ctx *core.Context, options types.AllOptions) ([]*File, error) {
	list := make([]*File, 0)

	db := database(ctx).Model(u)

	if options.Model != nil {
		db = ctx.Orm().GormWhere(db, u.TableName(), options.Model)
	}

	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}

	return list, db.Find(&list).Error
}

// Create 创建通知信息
func (u *File) Create(ctx *core.Context) error {
	md := ctx.Metadata()
	if md == nil {
		return errors.MetadataError
	}
	return transferErr(database(ctx).Create(u).Error)
}

// Update 更新文件信息
func (u *File) Update(ctx *core.Context) error {
	md := ctx.Metadata()
	if md == nil {
		return errors.MetadataError
	}
	return transferErr(database(ctx).Updates(&u).Error)
}

// DeleteByID 通过id删除
func (u *File) DeleteByID(ctx *core.Context, id int64) error {
	return transferErr(database(ctx).Delete(u, id).Error)
}
