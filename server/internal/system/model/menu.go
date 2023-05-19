package model

import (
	"encoding/json"
	"github.com/limeschool/easy-admin/server/consts"
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/tools"
	"github.com/limeschool/easy-admin/server/tools/lock"
	"github.com/limeschool/easy-admin/server/tools/proto"
	"github.com/limeschool/easy-admin/server/tools/tree"
	"github.com/limeschool/easy-admin/server/types"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Menu struct {
	ParentID   int64   `json:"parent_id"  gorm:"not nil;size:32;comment:父菜单id"`
	Title      string  `json:"title" gorm:"not nil;size:128;comment:菜单标题"`
	Icon       string  `json:"icon" gorm:"size:32;comment:菜单图标"`
	Path       string  `json:"path"  gorm:"type:varbinary(128);comment:菜单路径"`
	Name       string  `json:"name" gorm:"type:varbinary(32);comment:菜单唯一标志符"`
	Type       string  `json:"type" gorm:"not nil;size:32;comment:菜单类型"`
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

// IsBaseApiPath 获取基础菜单api
func (m *Menu) IsBaseApiPath(ctx *core.Context, method, path string) bool {
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
		return false
	}

	return data[strings.ToLower(method+":"+path)]
}

// All 获取全部的菜单列表
func (m *Menu) All(ctx *core.Context, cond ...interface{}) ([]*Menu, error) {
	var list []*Menu
	//Order("weight desc").
	return list, transferErr(database(ctx).Find(&list, cond...).Error)
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

func (m *Menu) InitData(ctx *core.Context) error {
	ins := []Menu{

		{
			ParentID: 0, Title: "根菜单", Icon: "", Path: "/", Name: "", Type: "R", Permission: "root", Method: "", Component: "", Redirect: nil, Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        1,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 1, Title: "系统管理", Icon: "settings", Path: "/system", Name: "System", Type: "M", Permission: "", Method: "", Component: "Layout", Redirect: proto.String("/system/user"), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(true), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        2,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 1, Title: "基本接口", Icon: "apps", Path: "/baseApi", Name: "baseApi", Type: "BA", Permission: "", Method: "", Component: "", Redirect: proto.String(""), Weight: proto.Int(1), IsHidden: proto.Bool(true), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        4,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 4, Title: "系统管理基础接口", Icon: "apps", Path: "/baseApi", Name: "baseApi", Type: "M", Permission: "", Method: "", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(true), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        5,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 5, Title: "获取当前用户信息", Icon: "", Path: "/api/system/user", Name: "", Type: "BA", Permission: "", Method: "GET", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        6,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 5, Title: "获取当前用户菜单", Icon: "", Path: "/api/system/user/menus", Name: "", Type: "BA", Permission: "", Method: "GET", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        7,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 5, Title: "获取系统部门信息", Icon: "", Path: "/api/system/teams", Name: "", Type: "BA", Permission: "", Method: "GET", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        8,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 2, Title: "菜单管理", Icon: "menu", Path: "/system/menu", Name: "systemMenu", Type: "M", Permission: "", Method: "", Component: "system/menu/index", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(true), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        9,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 9, Title: "查看菜单", Icon: "", Path: "/api/system/menus", Name: "", Type: "A", Permission: "system:menu:query", Method: "GET", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        10,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 9, Title: "新增菜单", Icon: "", Path: "/api/system/menu", Name: "", Type: "A", Permission: "system:menu:add", Method: "POST", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        11,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 9, Title: "修改菜单", Icon: "", Path: "/api/system/menu", Name: "", Type: "A", Permission: "system:menu:update", Method: "PUT", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        12,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 9, Title: "删除菜单", Icon: "", Path: "/api/system/menu", Name: "", Type: "A", Permission: "system:menu:delete", Method: "DELETE", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        13,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 2, Title: "部门管理", Icon: "user-group", Path: "/system/team", Name: "sysTeam", Type: "M", Permission: "", Method: "", Component: "system/team/index", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        14,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 14, Title: "新增部门", Icon: "", Path: "/api/system/team", Name: "", Type: "A", Permission: "system:team:add", Method: "POST", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        15,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 14, Title: "修改部门", Icon: "", Path: "/api/system/team", Name: "", Type: "A", Permission: "system:team:update", Method: "PUT", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        16,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 14, Title: "删除部门", Icon: "", Path: "/api/system/team", Name: "", Type: "A", Permission: "system:team:delete", Method: "DELETE", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        17,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 2, Title: "角色管理", Icon: "safe", Path: "/system/role", Name: "sysRole", Type: "M", Permission: "", Method: "", Component: "system/role/index", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        18,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 18, Title: "查看角色", Icon: "", Path: "/api/system/roles", Name: "", Type: "A", Permission: "system:role:query", Method: "GET", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        19,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 18, Title: "新增角色", Icon: "", Path: "/api/system/role", Name: "", Type: "A", Permission: "system:role:add", Method: "POST", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        20,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 18, Title: "修改角色", Icon: "", Path: "/api/system/role", Name: "", Type: "A", Permission: "system:role:update", Method: "PUT", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        21,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 18, Title: "删除角色", Icon: "", Path: "/api/system/role", Name: "", Type: "A", Permission: "system:role:delete", Method: "DELETE", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        22,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 2, Title: "用户管理", Icon: "user", Path: "/system/user", Name: "sysUser", Type: "M", Permission: "", Method: "", Component: "system/user/index", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        23,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 23, Title: "查看用户", Icon: "", Path: "/api/system/users", Name: "", Type: "A", Permission: "system:user:query", Method: "GET", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        24,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 23, Title: "新增用户", Icon: "", Path: "/api/system/user", Name: "", Type: "A", Permission: "system:user:add", Method: "POST", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        25,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 23, Title: "修改用户", Icon: "", Path: "/api/system/user", Name: "", Type: "A", Permission: "system:user:update", Method: "PUT", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        26,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 23, Title: "删除用户", Icon: "", Path: "/api/system/user", Name: "", Type: "A", Permission: "system:user:delete", Method: "DELETE", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        27,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 2, Title: "登陆日志", Icon: "history", Path: "/system/login_log", Name: "sysLoginLog", Type: "M", Permission: "", Method: "", Component: "system/login_log/index", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        28,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 28, Title: "查看登陆日志", Icon: "", Path: "/api/system/login/log", Name: "", Type: "A", Permission: "system:login:log:query", Method: "GET", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        29,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 18, Title: "修改角色菜单", Icon: "", Path: "", Name: "", Type: "G", Permission: "system:role:menu", Method: "", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        30,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 30, Title: "获取角色的菜单id", Icon: "", Path: "/api/system/role/menu_ids", Name: "", Type: "A", Permission: "system:role:menu:query", Method: "GET", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        31,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 30, Title: "修改角色菜单", Icon: "", Path: "/api/system/role/menu", Name: "", Type: "A", Permission: "system:role:menu:update", Method: "PUT", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        32,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 1, Title: "数据展板", Icon: "dashboard", Path: "/dashboard", Name: "Dashboard", Type: "M", Permission: "", Method: "", Component: "Layout", Redirect: proto.String("/dashboard/workplace"), Weight: proto.Int(1), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(true), BaseModel: types.BaseModel{
				ID:        33,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 33, Title: "系统数据", Icon: "dashboard", Path: "/dashboard/workplace", Name: "Workplace", Type: "M", Permission: "", Method: "", Component: "dashboard/workplace/index", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(true), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        34,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 5, Title: "刷新用户token", Icon: "", Path: "/api/system/token/refresh", Name: "", Type: "BA", Permission: "", Method: "POST", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        35,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 2, Title: "通知管理", Icon: "notification", Path: "/system/notice", Name: "Notice", Type: "M", Permission: "", Method: "", Component: "system/notice/index", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        36,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 36, Title: "新增系统通知", Icon: "", Path: "/api/system/notice", Name: "", Type: "A", Permission: "system:notice:add", Method: "POST", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        37,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 36, Title: "修改系统通知", Icon: "", Path: "/api/system/notice", Name: "", Type: "A", Permission: "system:notice:update", Method: "PUT", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        38,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 36, Title: "删除系统通知", Icon: "", Path: "/api/system/notice", Name: "", Type: "A", Permission: "system:notice:delete", Method: "DELETE", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        39,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 5, Title: "获取系统通知列表", Icon: "", Path: "/api/system/notices", Name: "", Type: "BA", Permission: "", Method: "GET", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        40,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 5, Title: "查看系统通知内容", Icon: "", Path: "/api/system/notice", Name: "", Type: "BA", Permission: "", Method: "GET", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        41,
				CreatedAt: time.Now().Unix(),
			},
		}, {
			ParentID: 5, Title: "获取用户通知未读数量", Icon: "", Path: "/api/system/notice/unread_num", Name: "", Type: "BA", Permission: "", Method: "GET", Component: "", Redirect: proto.String(""), Weight: proto.Int(0), IsHidden: proto.Bool(false), IsCache: proto.Bool(false), IsHome: proto.Bool(false), BaseModel: types.BaseModel{
				ID:        42,
				CreatedAt: time.Now().Unix(),
			},
		},
	}
	return database(ctx).Create(&ins).Error
}
