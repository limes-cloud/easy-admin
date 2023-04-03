package model

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/metadata"
	"github.com/limeschool/easy-admin/server/core/orm"
	"github.com/limeschool/easy-admin/server/global"
	"github.com/limeschool/easy-admin/server/tools"
	"github.com/limeschool/easy-admin/server/tools/tree"
	"time"
)

type User struct {
	orm.BaseModel
	TeamID      int64   `json:"team_id"`
	RoleID      int64   `json:"role_id"`
	Name        string  `json:"name"`
	Nickname    string  `json:"nickname"`
	Sex         *bool   `json:"sex,omitempty"`
	Phone       string  `json:"phone"`
	Password    string  `json:"password,omitempty"  gorm:"->:false;<-:create,update"`
	Avatar      string  `json:"avatar"`
	Email       string  `json:"email,omitempty"`
	Status      *bool   `json:"status,omitempty"`
	DisableDesc *string `json:"disable_desc"`
	LastLogin   int64   `json:"last_login"`
	Operator    string  `json:"operator"`
	OperatorID  int64   `json:"operator_id"`
	Role        Role    `json:"role" gorm:"->"`
	Team        Team    `json:"team" gorm:"->"`
}

func (u User) TableName() string {
	return "tb_system_user"
}

// OneByID 通过id查询用户信息
func (u *User) OneByID(ctx *gin.Context, id int64) error {
	db := database(ctx).Preload("Role").Preload("Team")
	return transferErr(db.First(u, id).Error)
}

// OneByPhone 通过phone查询用户信息
func (u *User) OneByPhone(ctx *gin.Context, phone string) error {
	db := database(ctx).Preload("Role").Preload("Team")
	return transferErr(db.First(u, "phone=?", phone).Error)
}

// PasswordByPhone 查询全部字段信息包括密码
func (u *User) PasswordByPhone(ctx *gin.Context, phone string) (string, error) {
	m := map[string]any{}
	if err := database(ctx).First(u, "phone = ?", phone).Scan(&m).Error; err != nil {
		return "", transferErr(err)
	}
	return m["password"].(string), nil
}

// Page 查询分页数据
func (u *User) Page(ctx *gin.Context, options orm.PageOptions) ([]*User, int64, error) {
	list, total := make([]*User, 0), int64(0)

	db := database(ctx).Model(u)

	if options.Model != nil {
		db = orm.GormWhere(db, u.TableName(), options.Model)
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
func (u *User) Create(ctx *gin.Context) error {
	md, err := metadata.GetFormContext(ctx)
	if err != nil {
		return err
	}
	u.Operator = md.Username
	u.OperatorID = md.UserID

	u.UpdatedAt = time.Now().Unix()
	u.Password, _ = tools.ParsePwd(u.Password)
	return transferErr(database(ctx).Create(u).Error)
}

func (u *User) UpdateLastLogin(ctx *gin.Context, t int64) error {
	return transferErr(database(ctx).Model(u).Where("id", u.ID).Update("last_login", t).Error)
}

// Update 更新用户信息
func (u *User) Update(ctx *gin.Context) error {
	md, err := metadata.GetFormContext(ctx)
	if err != nil {
		return err
	}
	u.Operator = md.Username
	u.OperatorID = md.UserID

	if u.Password != "" {
		u.Password, _ = tools.ParsePwd(u.Password)
	}

	// 执行更新
	return transferErr(database(ctx).Updates(&u).Error)
}

// DeleteByID 通过id删除用户信息
func (u *User) DeleteByID(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Delete(u, id).Error)
}

// GetAdminTeamIdByUserId 通过用户id获取用户所管理的部门id
func (u *User) GetAdminTeamIdByUserId(ctx *gin.Context, userId int64) ([]int64, error) {
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
	if role.DataScope == global.CURTEAM {
		return []int64{user.TeamID}, nil
	}

	ids := make([]int64, 0)
	// 当用户权限是自定义部门时，直接返回自定义部门id
	if role.DataScope == global.CUSTOM {
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
	case global.ALLTEAM:
		// 全部数据权限时返回所有部门id
		ids = tree.GetTreeID(teamTree)
	case global.DOWNTEAM:
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
