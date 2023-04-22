package model

import (
	"encoding/json"
	"github.com/limeschool/easy-admin/server/consts"
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/tools"
	"github.com/limeschool/easy-admin/server/tools/lock"
	"github.com/limeschool/easy-admin/server/tools/tree"
	"github.com/limeschool/easy-admin/server/types"
	"gorm.io/gorm"
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
	IsHome     bool    `json:"is_home" gorm:"->"`
	types.BaseModel
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
func (u *Menu) Create(ctx *core.Context) error {
	md := ctx.Metadata()
	if md == nil {
		return errors.MetadataError
	}

	u.Operator = md.Username
	u.OperatorID = md.UserID

	if u.Permission == consts.BaseApi {
		tools.DelayDelCache(ctx.Redis().GetRedis(consts.Cache), RedisBaseApiKey)
	}
	// 创建菜单
	return transferErr(database(ctx).Create(&u).Error)
}

// OneByID 通过id查询指定菜单
func (u *Menu) OneByID(ctx *core.Context, id int64) error {
	return transferErr(database(ctx).First(u, id).Error)
}

// OneByName 通过name条件查询指定菜单
func (u *Menu) OneByName(ctx *core.Context, name string) error {
	return transferErr(database(ctx).First(u, "name=?", name).Error)
}

// GetBaseApiPath 获取基础菜单api列表
func (u *Menu) GetBaseApiPath(ctx *core.Context) map[string]bool {
	lockKey := RedisBaseApiKey + "_lock"

	redis := ctx.Redis().GetRedis(consts.Cache)

	lc := lock.NewLockWithDuration(redis, lockKey, 10*time.Second)
	defer lc.Release()

	data := map[string]bool{}
	err := lc.AcquireFunc(func() error {
		if str, err := redis.Get(ctx, RedisBaseApiKey).Result(); err != nil {
			return err
		} else {
			return json.Unmarshal([]byte(str), &data)
		}
	}, func() error {
		list, err := u.All(ctx, "permission = ? and type = 'A'", consts.BaseApi)
		if err != nil {
			return err
		}

		// 转换格式
		for _, val := range list {
			data[strings.ToLower(val.Method+":"+val.Path)] = true
		}

		// 将数据库中的数据存入缓存
		b, _ := json.Marshal(data)
		redis.Set(ctx, RedisBaseApiKey, string(b), 1*time.Hour)
		return nil
	})

	if err != nil {
		return nil
	}

	return data
}

// All 获取全部的菜单列表
func (u *Menu) All(ctx *core.Context, cond ...interface{}) ([]*Menu, error) {
	var list []*Menu
	return list, transferErr(database(ctx).Order("weight desc").Find(&list, cond...).Error)
}

// Tree 获取菜单树
func (u *Menu) Tree(ctx *core.Context, cond ...interface{}) (tree.Tree, error) {
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
func (u *Menu) Update(ctx *core.Context) error {
	md := ctx.Metadata()
	if md == nil {
		return errors.MetadataError
	}

	u.Operator = md.Username
	u.OperatorID = md.UserID

	if u.Permission == consts.BaseApi {
		tools.DelayDelCache(ctx.Redis().GetRedis(consts.Cache), RedisBaseApiKey)
	}
	return transferErr(database(ctx).Updates(u).Error)
}

// UpdateHome 更新菜单首页
func (u *Menu) UpdateHome(ctx *core.Context, menuID int64) error {
	err := database(ctx).Table(u.TableName()).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id != ?", menuID).Update("is_home", false).Error; err != nil {
			return err
		}
		return tx.Where("id=?", menuID).Update("is_home", true).Error
	})
	return transferErr(err)
}

// DeleteByIds 通过条件删除菜单
func (u *Menu) DeleteByIds(ctx *core.Context, ids []int64) error {
	if err := database(ctx).First(u, "id in ?", ids).Error; err != nil {
		return transferErr(err)
	}
	// 删除基础api缓存
	if u.Permission == consts.BaseApi {
		tools.DelayDelCache(ctx.Redis().GetRedis(consts.Cache), RedisBaseApiKey)
	}
	return transferErr(database(ctx).Delete(u).Error)
}
