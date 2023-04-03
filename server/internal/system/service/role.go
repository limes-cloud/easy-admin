package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/limeschool/easy-admin/server/core/metadata"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/global"
	"github.com/limeschool/easy-admin/server/internal/system/model"
	"github.com/limeschool/easy-admin/server/internal/system/types"
	"github.com/limeschool/easy-admin/server/tools/tree"
)

// AllRole 返回所有的角色
func AllRole(ctx *gin.Context) (tree.Tree, error) {
	// 获取当前用户的角色
	md, err := metadata.GetFormContext(ctx)
	if err != nil {
		return nil, err
	}

	role := model.Role{}
	return role.Tree(ctx, md.RoleID)
}

// AddRole 新增角色
func AddRole(ctx *gin.Context, in *types.AddRoleRequest) error {
	role := model.Role{}
	if copier.Copy(&role, in) != nil {
		return errors.AssignError
	}
	return role.Create(ctx)
}

// UpdateRole 更新角色信息
func UpdateRole(ctx *gin.Context, in *types.UpdateRoleRequest) error {
	// 系统创建的角色不能修改
	if in.ID == 1 {
		return errors.SuperAdminEditError
	}

	// 获取当前用户信息
	md, err := metadata.GetFormContext(ctx)
	if err != nil {
		return err
	}

	// 不能自己禁用自己角色
	if in.Status != nil && !*in.Status {
		if in.ID == md.RoleID {
			return errors.DisableCurRoleError
		}
	}

	// 提交修改
	role := model.Role{}
	if copier.Copy(&role, in) != nil {
		return errors.AssignError
	}

	return role.Update(ctx)
}

// DeleteRole 删除角色信息
func DeleteRole(ctx *gin.Context, in *types.DeleteRoleRequest) error {
	// // 系统创建的角色不能删除
	if in.ID == 1 {
		return errors.SuperAdminDelError
	}

	// 删除角色时需要删除rbac权限表
	role := model.Role{}
	if err := role.OneByID(ctx, in.ID); err != nil {
		return err
	}
	_, _ = global.Casbin.RemoveFilteredPolicy(0, role.Keyword)

	return role.DeleteByID(ctx, in.ID)
}
