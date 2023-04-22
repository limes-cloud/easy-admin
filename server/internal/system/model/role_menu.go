package model

import (
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/types"
	"gorm.io/gorm"
)

type RoleMenu struct {
	types.BaseModel
	RoleID     int64  `json:"role_id"`
	MenuID     int64  `json:"menu_id"`
	Operator   string `json:"operator"`
	OperatorID int64  `json:"operator_id"`
}

func (RoleMenu) TableName() string {
	return "tb_system_role_menu"
}

// Update 批量更新角色所属菜单
func (r *RoleMenu) Update(ctx *core.Context, roleId int64, menuIds []int64) error {
	// 操作者信息
	md := ctx.Metadata()
	if md == nil {
		return errors.MetadataError
	}

	r.Operator = md.Username
	r.OperatorID = md.UserID

	// 组装新的菜单数据
	list := make([]RoleMenu, 0)
	for _, menuId := range menuIds {
		list = append(list, RoleMenu{
			RoleID:     roleId,
			MenuID:     menuId,
			OperatorID: md.UserID,
			Operator:   md.Username,
		})
	}

	// 删除之后再重新创建
	err := database(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("role_id=?", roleId).Delete(r).Error; err != nil {
			return err
		}
		return tx.Create(&list).Error
	})

	return transferErr(err)
}

// RoleMenus 通过角色ID获取角色菜单
func (r *RoleMenu) RoleMenus(ctx *core.Context, roleId int64) ([]*RoleMenu, error) {
	var list []*RoleMenu
	return list, transferErr(database(ctx).Find(&list, "role_id=?", roleId).Error)
}

// MenuRoles 通过菜单ID获取角色菜单列表
func (r *RoleMenu) MenuRoles(ctx *core.Context, menuId int64) ([]*RoleMenu, error) {
	var list []*RoleMenu
	return list, transferErr(database(ctx).Find(&list, "menu_id=?", menuId).Error)
}

// DeleteByRoleID 通过角色id删除 角色所属菜单
func (r *RoleMenu) DeleteByRoleID(ctx *core.Context, roleId int64) error {
	return transferErr(database(ctx).Delete(r, "role_id=?", roleId).Error)
}
