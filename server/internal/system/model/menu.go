package model

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/metadata"
	"github.com/limeschool/easy-admin/server/core/orm"
	"github.com/limeschool/easy-admin/server/global"
	"github.com/limeschool/easy-admin/server/tools/lock"
	"github.com/limeschool/easy-admin/server/tools/tree"
	"strings"
	"time"
)

type Menu struct {
	ParentID   int64   `json:"parent_id"`
	Title      string  `json:"title"`
	Icon       string  `json:"icon"`
	Path       string  `json:"path"`
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	Permission string  `json:"permission"`
	Method     string  `json:"method"`
	Component  string  `json:"component"`
	Redirect   *string `json:"redirect"`
	Weight     *int    `json:"weight"`
	IsHidden   *bool   `json:"is_hidden"`
	IsCache    *bool   `json:"is_cache"`
	Operator   string  `json:"operator"`
	OperatorID int64   `json:"operator_id"`
	Children   []*Menu `json:"children,omitempty" gorm:"-"`
	orm.BaseModel
}

const (
	RedisBaseApiKey = "sysBaseApi"
)

func (u *Menu) ID() int64 {
	return u.BaseModel.ID
}

func (u *Menu) Parent() int64 {
	return u.ParentID
}

func (u *Menu) AppendChildren(child any) {
	menu := child.(*Menu)
	u.Children = append(u.Children, menu)
}

func (u *Menu) ChildrenNode() []tree.Tree {
	var list []tree.Tree
	for _, item := range u.Children {
		list = append(list, item)
	}
	return list
}

func (u *Menu) TableName() string {
	return "tb_system_menu"
}

// Create 创建菜单
func (u *Menu) Create(ctx *gin.Context) error {
	md, err := metadata.GetFormContext(ctx)
	if err != nil {
		return err
	}
	u.Operator = md.Username
	u.OperatorID = md.UserID

	if u.Permission == global.BaseApi {
		orm.DelayDelCache(global.Redis, RedisBaseApiKey)
	}
	// 创建菜单
	return transferErr(database(ctx).Create(&u).Error)
}

// OneByID 通过id查询指定菜单
func (u *Menu) OneByID(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).First(u, id).Error)
}

// OneByName 通过name条件查询指定菜单
func (u *Menu) OneByName(ctx *gin.Context, name string) error {
	return transferErr(database(ctx).First(u, "name=?", name).Error)
}

// GetBaseApiPath 获取基础菜单api列表
func (u *Menu) GetBaseApiPath(ctx *gin.Context) map[string]bool {
	lockKey := RedisBaseApiKey + "_lock"

	lc := lock.NewLockWithDuration(global.Redis, lockKey, 10*time.Second)
	data, err := lc.AcquireFunc(func() (any, error) {
		// 从缓存中获取数据
		m := map[string]bool{}
		str, err := global.Redis.Get(ctx, RedisBaseApiKey).Result()
		if err != nil {
			return nil, err
		}
		return m, json.Unmarshal([]byte(str), &m)
	}, func() (any, error) {

		// 从数据库中读取
		m := map[string]bool{}
		list, err := u.All(ctx, "permission = ? and type = 'A'", global.BaseApi)
		if err != nil {
			return nil, err
		}

		// 转换格式
		for _, val := range list {
			m[strings.ToLower(val.Method+":"+val.Path)] = true
		}

		// 将数据库中的数据存入缓存
		b, _ := json.Marshal(m)
		global.Redis.Set(ctx, RedisBaseApiKey, string(b), 1*time.Hour)
		return m, nil
	})

	if err != nil {
		return nil
	}
	return data.(map[string]bool)
}

// All 获取全部的菜单列表
func (u *Menu) All(ctx *gin.Context, cond ...interface{}) ([]*Menu, error) {
	var list []*Menu
	return list, transferErr(database(ctx).Order("weight desc").Find(&list, cond...).Error)
}

// Tree 获取菜单树
func (u *Menu) Tree(ctx *gin.Context, cond ...interface{}) (tree.Tree, error) {
	list, err := u.All(ctx, cond...)
	if err != nil {
		return nil, err
	}
	var treeList []tree.Tree
	for _, item := range list {
		treeList = append(treeList, item)
	}
	return tree.BuildTree(treeList), nil
}

// Update 更新菜单
func (u *Menu) Update(ctx *gin.Context) error {
	md, err := metadata.GetFormContext(ctx)
	if err != nil {
		return err
	}
	u.Operator = md.Username
	u.OperatorID = md.UserID

	if u.Permission == global.BaseApi {
		orm.DelayDelCache(global.Redis, RedisBaseApiKey)
	}
	return transferErr(database(ctx).Updates(u).Error)
}

// DeleteByIds 通过条件删除菜单
func (u *Menu) DeleteByIds(ctx *gin.Context, ids []int64) error {
	if err := database(ctx).First(u, "id in ?", ids).Error; err != nil {
		return transferErr(err)
	}
	// 删除基础api缓存
	if u.Permission == global.BaseApi {
		orm.DelayDelCache(global.Redis, RedisBaseApiKey)
	}
	return transferErr(database(ctx).Delete(u).Error)
}
