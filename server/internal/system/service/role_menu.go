package service

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/global"
	"github.com/limeschool/easy-admin/server/internal/system/model"
	"github.com/limeschool/easy-admin/server/internal/system/types"
	"github.com/limeschool/easy-admin/server/tools/tree"
)

// UpdateRoleMenu 修改角色所属菜单
func UpdateRoleMenu(ctx *gin.Context, in *types.AddRoleMenuRequest) error {
	// 超级管理员不存在菜单权限，自动获取全部菜单，所以禁止修改
	if in.RoleID == 1 {
		return errors.SuperAdminEditError
	}

	// 获取当前role的数据
	role := model.Role{}
	if err := role.OneByID(ctx, in.RoleID); err != nil {
		return err
	}

	// 进行菜单修改
	rm := model.RoleMenu{}
	if err := rm.Update(ctx, in.RoleID, in.MenuIds); err != nil {
		return err
	}

	// 删除当前用户的全部rbac权限
	_, _ = global.Casbin.RemoveFilteredPolicy(0, role.Keyword)

	// 获取当前修改菜单的信息
	menu := model.Menu{}
	var policies [][]string
	apiList, _ := menu.All(ctx, "id in ? and type = 'A'", in.MenuIds)
	for _, item := range apiList {
		policies = append(policies, []string{role.Keyword, item.Path, item.Method})
	}

	// 将新的策略的策略写入rbac
	_, _ = global.Casbin.AddPolicies(policies)

	return nil
}

// RoleMenuIds 获取角色菜单的所有id
func RoleMenuIds(ctx *gin.Context, in *types.RoleMenuIdsRequest) ([]int64, error) {

	// 获取当前角色的所有菜单
	rm := model.RoleMenu{}
	rmList, err := rm.RoleMenus(ctx, in.RoleID)
	if err != nil {
		return nil, err
	}

	// 组装所有的菜单id
	var ids []int64
	for _, item := range rmList {
		ids = append(ids, item.MenuID)
	}

	return ids, nil
}

func RoleMenu(ctx *gin.Context, in *types.RoleMenuRequest) (tree.Tree, error) {
	// 查询角色信息
	role := model.Role{}
	if err := role.OneByID(ctx, in.RoleID); err != nil {
		return nil, err
	}

	var menus []*model.Menu

	var menu model.Menu

	if role.Keyword == global.JwtSuperAdmin {
		menus, _ = menu.All(ctx, "permission!=?", global.BaseApi)
	} else {
		// 查询角色所属菜单
		rm := model.RoleMenu{}
		rmList, err := rm.RoleMenus(ctx, in.RoleID)
		if err != nil {
			return nil, err
		}

		// 获取菜单的所有id
		var ids []int64
		for _, item := range rmList {
			ids = append(ids, item.MenuID)
		}
		if len(ids) == 0 {
			return nil, nil
		}

		// 获取指定id的所有菜单
		menus, _ = menu.All(ctx, "id in ? and permission!=?", ids, global.BaseApi)
	}

	var listTree []tree.Tree
	for _, item := range menus {
		listTree = append(listTree, item)
	}

	return tree.BuildTree(listTree), nil
}
