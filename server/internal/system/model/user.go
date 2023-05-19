package model

import (
	"encoding/json"
	"github.com/limeschool/easy-admin/server/consts"
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/tools"
	"github.com/limeschool/easy-admin/server/tools/proto"
	"github.com/limeschool/easy-admin/server/tools/tree"
	"github.com/limeschool/easy-admin/server/types"
	"time"
)

type User struct {
	types.BaseModel
	TeamID      int64   `json:"team_id" gorm:"not null;size:32;comment:部门id"`
	RoleID      int64   `json:"role_id" gorm:"not null;size:32;comment:角色id"`
	Name        string  `json:"name" gorm:"not null;size:32;comment:用户姓名"`
	Nickname    string  `json:"nickname" gorm:"not null;size:128;comment:用户昵称"`
	Sex         *bool   `json:"sex,omitempty" gorm:"not null;comment:用户性别"`
	Phone       string  `json:"phone" gorm:"not null;size:32;comment:用户电话"`
	Password    string  `json:"password,omitempty" gorm:"not null;->:false;<-:create,update;comment:用户密码"`
	Avatar      string  `json:"avatar" gorm:"not null;size:128;comment:用户头像"`
	Email       string  `json:"email,omitempty" gorm:"not null;type:varbinary(128);comment:用户邮箱"`
	Status      *bool   `json:"status,omitempty" gorm:"not null;comment:用户状态"`
	DisableDesc *string `json:"disable_desc" gorm:"not null;size:128;comment:禁用原因"`
	LastLogin   int64   `json:"last_login" gorm:"comment:最后登陆时间"`
	Operator    string  `json:"operator" gorm:"not null;size:128;comment:操作人员名称"`
	OperatorID  int64   `json:"operator_id" gorm:"not null;size:32;comment:操作人员id"`
	Role        Role    `json:"role" gorm:"->"`
	Team        Team    `json:"team" gorm:"->"`
}

func (u *User) TableName() string {
	return "tb_system_user"
}

// OneByID 通过id查询用户信息
func (u *User) OneByID(ctx *core.Context, id int64) error {
	db := database(ctx).Preload("Role").Preload("Team")
	return transferErr(db.First(u, id).Error)
}

// OneByPhone 通过phone查询用户信息
func (u *User) OneByPhone(ctx *core.Context, phone string) error {
	db := database(ctx).Preload("Role").Preload("Team")
	return transferErr(db.First(u, "phone=?", phone).Error)
}

// PasswordByPhone 查询全部字段信息包括密码
func (u *User) PasswordByPhone(ctx *core.Context, phone string) (string, error) {
	m := map[string]any{}
	if err := database(ctx).First(u, "phone = ?", phone).Scan(&m).Error; err != nil {
		return "", transferErr(err)
	}
	return m["password"].(string), nil
}

// Page 查询分页数据
func (u *User) Page(ctx *core.Context, options types.PageOptions) ([]*User, int64, error) {
	list, total := make([]*User, 0), int64(0)

	db := database(ctx).Model(u)

	if options.Model != nil {
		db = ctx.Orm().GormWhere(db, u.TableName(), options.Model)
	}

	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	db = db.Preload("Role").Preload("Team")
	db = db.Offset((options.Page - 1) * options.PageSize).Limit(options.PageSize)

	return list, total, db.Find(&list).Error
}

// Create 创建用户信息
func (u *User) Create(ctx *core.Context) error {
	md := ctx.Metadata()
	if md == nil {
		return errors.MetadataError
	}
	u.Operator = md.Username
	u.OperatorID = md.UserID

	u.UpdatedAt = time.Now().Unix()
	u.Password = tools.ParsePwd(u.Password)
	return transferErr(database(ctx).Create(u).Error)
}

func (u *User) UpdateLastLogin(ctx *core.Context, t int64) error {
	return transferErr(database(ctx).Model(u).Where("id", u.ID).Update("last_login", t).Error)
}

// Update 更新用户信息
func (u *User) Update(ctx *core.Context) error {
	md := ctx.Metadata()
	if md == nil {
		return errors.MetadataError
	}
	u.Operator = md.Username
	u.OperatorID = md.UserID

	if u.Password != "" {
		u.Password = tools.ParsePwd(u.Password)
	}

	// 执行更新
	return transferErr(database(ctx).Updates(&u).Error)
}

// DeleteByID 通过id删除用户信息
func (u *User) DeleteByID(ctx *core.Context, id int64) error {
	return transferErr(database(ctx).Delete(u, id).Error)
}

// GetAdminTeamIdByUserId 通过用户id获取用户所管理的部门id
func (u *User) GetAdminTeamIdByUserId(ctx *core.Context, userId int64) ([]int64, error) {
	// 操作者信息
	user := User{}
	if err := user.OneByID(ctx, userId); err != nil {
		return nil, err
	}

	// 查询角色信息
	role := Role{}
	if err := role.OneByID(ctx, user.RoleID); err != nil {
		return nil, err
	}

	// 当用户权限是当前部门时，直接返回当前部门的id
	if role.DataScope == consts.CURTEAM {
		return []int64{user.TeamID}, nil
	}

	ids := make([]int64, 0)
	// 当用户权限是自定义部门时，直接返回自定义部门id
	if role.DataScope == consts.CUSTOM {
		return ids, json.Unmarshal([]byte(*role.TeamIds), &ids)
	}

	// 以当前部门为根节点构造部门树
	team := Team{}
	teamList, _ := team.All(ctx)
	var treeList []tree.Tree
	for _, item := range teamList {
		treeList = append(treeList, item)
	}
	teamTree := tree.BuildTreeByID(treeList, user.TeamID)

	// 根据部门树取值
	switch role.DataScope {
	case consts.ALLTEAM:
		// 全部数据权限时返回所有部门id
		ids = tree.GetTreeID(teamTree)
	case consts.DOWNTEAM:
		// 下级部门权限时，排除当前部门id
		ids = tree.GetTreeID(teamTree)
		if len(ids) > 2 {
			ids = ids[1:]
		} else {
			ids = []int64{}
		}
	}
	return ids, nil
}

func (u *User) InitData(ctx *core.Context) error {
	ins := User{
		BaseModel:   types.BaseModel{ID: 1, CreatedAt: time.Now().Unix()},
		TeamID:      1,
		RoleID:      1,
		Name:        "超级管理员",
		Nickname:    "superAdmin",
		Sex:         proto.Bool(true),
		Phone:       "18888888888",
		Password:    tools.ParsePwd("123456"),
		Email:       "18888888888@qq.com",
		Status:      proto.Bool(true),
		DisableDesc: proto.String(""),
	}
	return database(ctx).Create(&ins).Error

}
