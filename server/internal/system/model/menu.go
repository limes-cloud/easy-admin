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
	ParentID   int64   `json:"parent_id"  gorm:"not null;size:32;comment:父菜单id"`
	Title      string  `json:"title" gorm:"not null;size:128;comment:菜单标题"`
	Icon       string  `json:"icon" gorm:"size:32;comment:菜单图标"`
	Path       string  `json:"path"  gorm:"type:varbinary(128);comment:菜单路径"`
	Name       string  `json:"name" gorm:"type:varbinary(32);comment:菜单唯一标志符"`
	Type       string  `json:"type" gorm:"not null;size:32;comment:菜单类型"`
	Permission string  `json:"permission" gorm:"size:32;comment:菜单指令"`
	Method     string  `json:"method" gorm:"size:32;comment:接口方法"`
	Component  string  `json:"component" gorm:"size:128;comment:组件地址"`
	Redirect   *string `json:"redirect" gorm:"size:128;comment:重定向地址"`
	Weight     *int    `json:"weight" gorm:"default:0;size:16;comment:菜单权重"`
	IsHidden   *bool   `json:"is_hidden" gorm:"default:false;comment:是否隐藏"`
	IsCache    *bool   `json:"is_cache" gorm:"default:false;comment:是否缓存"`
	IsHome     *bool   `json:"is_home" gorm:"->" gorm:"default:false;comment:是否为首页"` // todo change
	Operator   string  `json:"operator" gorm:"size:128;comment:操作人员名称"`
	OperatorID int64   `json:"operator_id" gorm:"size:32;comment:操作人员id"`
	Children   []*Menu `json:"children,omitempty" gorm:"-"`
	types.BaseModel
}

const (
	RedisBaseApiKey = "sysBaseApi"
)

func (m *Menu) ID() int64 {
	return m.BaseModel.ID
}

func (m *Menu) Parent() int64 {
	return m.ParentID
}

func (m *Menu) AppendChildren(child any) {
	menu := child.(*Menu)
	m.Children = append(m.Children, menu)
}

func (m *Menu) ChildrenNode() []tree.Tree {
	var list []tree.Tree
	for _, item := range m.Children {
		list = append(list, item)
	}
	return list
}

func (m *Menu) TableName() string {
	return "tb_system_menu"
}

// Create 创建菜单
func (m *Menu) Create(ctx *core.Context) error {
	md := ctx.Metadata()
	if md == nil {
		return errors.MetadataError
	}

	m.Operator = md.Username
	m.OperatorID = md.UserID

	if m.Type == consts.MenuBA {
		tools.DelayDelCache(ctx.Redis().GetRedis(consts.Cache), RedisBaseApiKey)
	}
	// 创建菜单
	return transferErr(database(ctx).Create(&m).Error)
}

// OneByID 通过id查询指定菜单
func (m *Menu) OneByID(ctx *core.Context, id int64) error {
	return transferErr(database(ctx).First(m, id).Error)
}

// OneByName 通过name条件查询指定菜单
func (m *Menu) OneByName(ctx *core.Context, name string) error {
	return transferErr(database(ctx).First(m, "name=?", name).Error)
}

// GetBaseApiPath 获取基础菜单api列表
func (m *Menu) GetBaseApiPath(ctx *core.Context) map[string]bool {
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
		list, err := m.All(ctx, "type = ?", consts.MenuBA)
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
func (m *Menu) All(ctx *core.Context, cond ...interface{}) ([]*Menu, error) {
	var list []*Menu
	return list, transferErr(database(ctx).Order("weight desc").Find(&list, cond...).Error)
}

// Tree 获取菜单树
func (m *Menu) Tree(ctx *core.Context, cond ...interface{}) (tree.Tree, error) {
	list, err := m.All(ctx, cond...)
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
func (m *Menu) Update(ctx *core.Context) error {
	md := ctx.Metadata()
	if md == nil {
		return errors.MetadataError
	}

	m.Operator = md.Username
	m.OperatorID = md.UserID

	if m.Type == consts.MenuBA {
		tools.DelayDelCache(ctx.Redis().GetRedis(consts.Cache), RedisBaseApiKey)
	}
	return transferErr(database(ctx).Updates(m).Error)
}

// UpdateHome 更新菜单首页
func (m *Menu) UpdateHome(ctx *core.Context, menuID int64) error {
	err := database(ctx).Table(m.TableName()).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id != ?", menuID).Update("is_home", false).Error; err != nil {
			return err
		}
		return tx.Where("id=?", menuID).Update("is_home", true).Error
	})
	return transferErr(err)
}

// DeleteByIds 通过条件删除菜单
func (m *Menu) DeleteByIds(ctx *core.Context, ids []int64) error {
	if err := database(ctx).First(m, "id in ?", ids).Error; err != nil {
		return transferErr(err)
	}
	// 删除基础api缓存
	if m.Type == consts.MenuBA {
		tools.DelayDelCache(ctx.Redis().GetRedis(consts.Cache), RedisBaseApiKey)
	}
	return transferErr(database(ctx).Delete(m).Error)
}
