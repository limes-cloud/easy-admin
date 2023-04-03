package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/global"
	"github.com/limeschool/easy-admin/server/internal/system/model"
	"github.com/limeschool/easy-admin/server/internal/system/types"
	"github.com/limeschool/easy-admin/server/tools/tree"
)

// AllMenu 获取菜单树
func AllMenu(ctx *gin.Context) (tree.Tree, error) {
	menu := model.Menu{}
	return menu.Tree(ctx)
}

// AddMenu 新增菜单
func AddMenu(ctx *gin.Context, in *types.AddMenuRequest) error {
	menu := model.Menu{}
	if copier.Copy(&menu, in) != nil {
		return errors.AssignError
	}

	// 判断name值是否重复
	if in.Name != "" && menu.OneByName(ctx, in.Name) == nil {
		return errors.DulMenuNameError
	}

	return menu.Create(ctx)
}

// UpdateMenu 更新菜单
func UpdateMenu(ctx *gin.Context, in *types.UpdateMenuRequest) error {
	inMenu := model.Menu{}
	if copier.Copy(&inMenu, in) != nil {
		return errors.AssignError
	}

	menu := model.Menu{}
	if err := menu.OneByID(ctx, in.ID); err != nil {
		return err
	}

	if in.ParentID != 0 && in.ID == in.ParentID {
		return errors.MenuParentIdError
	}

	if menu.Name != in.Name && menu.OneByName(ctx, in.Name) == nil {
		return errors.DulMenuNameError
	}

	// 之前为接口，现在修改类型不为接口，则删除之前的rbac数据
	if menu.Type == "A" && in.Type != "A" {
		_, _ = global.Casbin.RemoveFilteredPolicy(1, menu.Path, menu.Method)
	}

	// 之前和现在都为接口，且存在方法或者路径变更时则更新rbac数据
	if menu.Type == "A" && in.Type == "A" && (menu.Method != in.Method || menu.Path != in.Path) {
		oldPolices := global.Casbin.GetFilteredPolicy(1, menu.Path, menu.Method)
		if len(oldPolices) != 0 {
			var newPolices [][]string
			for _, val := range oldPolices {
				newPolices = append(newPolices, []string{val[0], in.Path, in.Method})
			}
			_, _ = global.Casbin.UpdatePolicies(oldPolices, newPolices)
		}
	}

	// 当之前不是接口，现在是接口的情况下，则进行新增
	if menu.Type != "A" && in.Type == "A" {
		// 获取选中当前菜单的角色
		roleMenu := model.RoleMenu{}
		roleMenus, _ := roleMenu.RoleMenus(ctx, in.ID)
		if len(roleMenus) != 0 {
			var roleIds []int64
			for _, item := range roleMenus {
				roleIds = append(roleIds, item.RoleID)
			}

			// 获取当前菜单的全部角色信息
			role := model.Role{}
			roles, _ := role.All(ctx, "id in ?", roleIds)

			// 添加菜单到rbac权限表
			var newPolices [][]string
			for _, val := range roles {
				newPolices = append(newPolices, []string{val.Keyword, in.Path, in.Method})
			}
			_, _ = global.Casbin.AddPolicies(newPolices)
		}
	}

	return inMenu.Update(ctx)
}

// DeleteMenu 删除菜单
func DeleteMenu(ctx *gin.Context, in *types.DeleteMenuRequest) error {
	if in.ID == 1 {
		return errors.DeleteRootMenuError
	}
	// 获取指定id为根节点的菜单树
	menu := model.Menu{}
	list, _ := menu.All(ctx)
	var treeList []tree.Tree
	for _, item := range list {
		treeList = append(treeList, item)
	}
	t := tree.BuildTreeByID(treeList, in.ID)

	// 获取菜单树下的所有菜单ID
	ids := tree.GetTreeID(t)

	// 删除当前id中的类型为api的rbac权限表
	apiList, _ := menu.All(ctx, "id in ? and type='A'", ids)
	for _, item := range apiList {
		_, _ = global.Casbin.RemoveFilteredPolicy(1, item.Path, item.Method)
	}

	// 从数据库删除菜单
	return menu.DeleteByIds(ctx, ids)
}
