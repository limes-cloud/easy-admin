package model

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/metadata"
	"github.com/limeschool/easy-admin/server/core/orm"
	"github.com/limeschool/easy-admin/server/tools/tree"
)

type Role struct {
	orm.BaseModel
	ParentID    int64   `json:"parent_id"`
	Name        string  `json:"name" `
	Keyword     string  `json:"keyword"`
	Status      *bool   `json:"status,omitempty" `
	Weight      *int    `json:"weight"`
	Description *string `json:"description,omitempty"`
	TeamIds     *string `json:"team_ids,omitempty"`
	DataScope   string  `json:"data_scope,omitempty"`
	Operator    string  `json:"operator"`
	OperatorID  int64   `json:"operator_id"`
	Children    []*Role `json:"children"  gorm:"-"`
}

func (r *Role) ID() int64 {
	return r.BaseModel.ID
}

func (r *Role) Parent() int64 {
	return r.ParentID
}

func (r *Role) AppendChildren(child any) {
	menu := child.(*Role)
	r.Children = append(r.Children, menu)
}

func (r *Role) ChildrenNode() []tree.Tree {
	var list []tree.Tree
	for _, item := range r.Children {
		list = append(list, item)
	}
	return list
}

func (r *Role) TableName() string {
	return "tb_system_role"
}

// Create 创建角色信息
func (r *Role) Create(ctx *gin.Context) error {
	md, err := metadata.GetFormContext(ctx)
	if err != nil {
		return err
	}
	r.Operator = md.Username
	r.OperatorID = md.UserID

	return transferErr(database(ctx).Create(&r).Error)
}

// OneByID 通过ID查询角色信息
func (r *Role) OneByID(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).First(r, "id = ?", id).Error)
}

// All 查询全部角色信息
func (r *Role) All(ctx *gin.Context, cond ...any) ([]*Role, error) {
	var list []*Role
	return list, transferErr(database(ctx).Order("weight desc").Find(&list, cond...).Error)
}

// RoleStatus 获取角色状态
func (r *Role) RoleStatus(ctx *gin.Context, roleId int64) bool {
	team, err := r.Tree(ctx, 1)
	if err != nil {
		return false
	}
	res := false
	dfsRoleStatus(team.(*Role), roleId, true, &res)
	return res
}

func dfsRoleStatus(role *Role, roleId int64, status bool, res *bool) bool {
	if roleId == role.BaseModel.ID {
		is := *role.Status && status
		*res = is
	}

	for _, item := range role.Children {
		dfsRoleStatus(item, roleId, status && *item.Status, res)
	}

	return status
}

// Tree 查询指定角色为根节点的角色树
func (r *Role) Tree(ctx *gin.Context, roleId int64) (tree.Tree, error) {
	list, err := r.All(ctx)
	if err != nil {
		return nil, err
	}
	var treeList []tree.Tree
	for _, item := range list {
		treeList = append(treeList, item)
	}
	return tree.BuildTreeByID(treeList, roleId), nil
}

// Update 更新角色信息
func (r *Role) Update(ctx *gin.Context) error {
	md, err := metadata.GetFormContext(ctx)
	if err != nil {
		return err
	}
	r.Operator = md.Username
	r.OperatorID = md.UserID

	return transferErr(database(ctx).Updates(r).Error)
}

// DeleteByID 通过ID删除角色信息
func (r *Role) DeleteByID(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Where("id = ?", id).Delete(&r).Error)
}
