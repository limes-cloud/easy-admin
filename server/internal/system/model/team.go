package model

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/metadata"
	"github.com/limeschool/easy-admin/server/core/orm"
	"github.com/limeschool/easy-admin/server/tools/tree"
)

type Team struct {
	orm.BaseModel
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	ParentID    int64   `json:"parent_id"`
	Operator    string  `json:"operator"`
	OperatorID  int64   `json:"operator_id"`
	Children    []*Team `json:"children,omitempty" gorm:"-"`
}

func (t *Team) ID() int64 {
	return t.BaseModel.ID
}

func (t *Team) Parent() int64 {
	return t.ParentID
}

func (t *Team) AppendChildren(child any) {
	team := child.(*Team)
	t.Children = append(t.Children, team)
}

func (t *Team) ChildrenNode() []tree.Tree {
	var list []tree.Tree
	for _, item := range t.Children {
		list = append(list, item)
	}
	return list
}

func (t *Team) TableName() string {
	return "tb_system_team"
}

// Create 创建部门
func (t *Team) Create(ctx *gin.Context) error {
	md, err := metadata.GetFormContext(ctx)
	if err != nil {
		return err
	}
	t.Operator = md.Username
	t.OperatorID = md.UserID

	return transferErr(database(ctx).Create(&t).Error)
}

// Tree 获取部门树
func (t *Team) Tree(ctx *gin.Context) (tree.Tree, error) {
	// 获取部门列表
	list := make([]*Team, 0)
	if err := database(ctx).Find(&list).Error; err != nil {
		return nil, err
	}

	// 根据列表构建部门树
	var trees []tree.Tree
	for _, item := range list {
		trees = append(trees, item)
	}
	return tree.BuildTree(trees), nil
}

// All 获取全部部门
func (t *Team) All(ctx *gin.Context) ([]*Team, error) {
	list := make([]*Team, 0)
	if err := database(ctx).Find(&list).Error; err != nil {
		return nil, transferErr(err)
	}
	return list, nil
}

// Update 更新部门信息
func (t *Team) Update(ctx *gin.Context) error {
	md, err := metadata.GetFormContext(ctx)
	if err != nil {
		return err
	}
	t.Operator = md.Username
	t.OperatorID = md.UserID

	return transferErr(database(ctx).Updates(t).Error)
}

// DeleteByID 通过id删除指定部门 以及部门下的部门
func (t *Team) DeleteByID(ctx *gin.Context, id int64) error {
	list, err := t.All(ctx)
	if err != nil {
		return err
	}

	var treeList []tree.Tree
	for _, item := range list {
		treeList = append(treeList, item)
	}
	team := tree.BuildTreeByID(treeList, id)
	ids := tree.GetTreeID(team)

	// 进行数据删除
	return transferErr(database(ctx).Where("id in ?", ids).Delete(&t).Error)
}
