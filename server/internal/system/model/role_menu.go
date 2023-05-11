package model

import (
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/types"
	"gorm.io/gorm"
)

type RoleMenu struct {
	types.BaseModel
	RoleID     int64  `json:"role_id" gorm:"not null;size:32;comment:角色id"`
	MenuID     int64  `json:"menu_id" gorm:"not null;size:32;comment:菜单id"`
	Operator   string `json:"operator" gorm:"not null;size:128;comment:操作人员名称"`
	OperatorID int64  `json:"operator_id" gorm:"not null;size:32;comment:操作人员id"`
	Role       Role   `json:"role" gorm:"->;constraint:OnDelete:cascade"`
	Menu       Menu   `json:"menu" gorm:"->;constraint:OnDelete:cascade"`
}

func (RoleMenu) TableName() string {
	return "tb_system_role_menu"
}

// Update 批量更新角色所属菜单
func (rm *RoleMenu) Update(ctx *core.Context, roleId int64, menuIds []int64) error {
	// 操作者信息
	md := ctx.Metadata()
	if md == nil {
		return errors.MetadataError
	}

	rm.Operator = md.Username
	rm.OperatorID = md.UserID

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
		if err := tx.Where("role_id=?", roleId).Delete(rm).Error; err != nil {
			return err
		}
		return tx.Create(&list).Error
	})

	return transferErr(err)
}

// RoleMenus 通过角色ID获取角色菜单
func (rm *RoleMenu) RoleMenus(ctx *core.Context, roleId int64) ([]*RoleMenu, error) {
	var list []*RoleMenu
	return list, transferErr(database(ctx).Find(&list, "role_id=?", roleId).Error)
}

// MenuRoles 通过菜单ID获取角色菜单列表
func (rm *RoleMenu) MenuRoles(ctx *core.Context, menuId int64) ([]*RoleMenu, error) {
	var list []*RoleMenu
	return list, transferErr(database(ctx).Find(&list, "menu_id=?", menuId).Error)
}

// DeleteByRoleID 通过角色id删除 角色所属菜单
func (rm *RoleMenu) DeleteByRoleID(ctx *core.Context, roleId int64) error {
	return transferErr(database(ctx).Delete(rm, "role_id=?", roleId).Error)
}
