package service

import (
	"encoding/base64"
	"encoding/json"
	"github.com/forgoer/openssl"
	"github.com/jinzhu/copier"
	"github.com/limeschool/easy-admin/server/consts"
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/internal/system/model"
	"github.com/limeschool/easy-admin/server/internal/system/types"
	"github.com/limeschool/easy-admin/server/tools"
	"github.com/limeschool/easy-admin/server/tools/tree"
	tps "github.com/limeschool/easy-admin/server/types"
	"gorm.io/gorm"
	"time"
)

const (
	encodePasswordCert = "encodePassword"
	decodePasswordCert = "decodePassword"
)

// CurrentAdminTeamIds 获取当前用户的管理的部门id
func CurrentAdminTeamIds(ctx *core.Context) ([]int64, error) {
	md := ctx.Metadata()
	if md == nil {
		return nil, errors.MetadataError
	}

	user := model.User{}
	ids, err := user.GetAdminTeamIdByUserId(ctx, md.UserID)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

// CurrentUser 获取当前用户信息
func CurrentUser(ctx *core.Context) (*model.User, error) {
	md := ctx.Metadata()
	if md == nil {
		return nil, errors.MetadataError
	}

	user := model.User{}
	return &user, user.OneByID(ctx, md.UserID)
}

// PageUser 获取用户分页信息
func PageUser(ctx *core.Context, in *types.PageUserRequest) ([]*model.User, int64, error) {
	user := model.User{}
	// 获取用户所管理的部门
	ids, err := CurrentAdminTeamIds(ctx)
	if err != nil {
		return nil, 0, err
	}

	return user.Page(ctx, tps.PageOptions{
		Page:     in.Page,
		PageSize: in.PageSize,
		Model:    in,
		Scopes: func(db *gorm.DB) *gorm.DB {
			return db.Where("team_id in ?", ids)
		},
	})
}

// AddUser 新增用户信息
func AddUser(ctx *core.Context, in *types.AddUserRequest) error {
	user := model.User{}
	if in.Nickname == "" {
		in.Nickname = in.Name
	}

	if copier.Copy(&user, in) != nil {
		return errors.AssignError
	}

	// 获取用户所管理的部门
	ids, err := CurrentAdminTeamIds(ctx)
	if err != nil {
		return err
	}

	// 添加用户时，只允许添加当前所属部门的用户
	if !tools.InList(ids, in.TeamID) {
		return errors.NotAddTeamUserError
	}

	return user.Create(ctx)
}

// UpdateUser 更新用户信息
func UpdateUser(ctx *core.Context, in *types.UpdateUserRequest) error {
	user := model.User{}
	if user.OneByID(ctx, in.ID) != nil {
		return errors.DBNotFoundError
	}

	//超级管理员不允许修改所在部门和角色
	if in.ID == 1 {
		in.RoleID = 0
		in.TeamID = 0
		if *user.Status != *in.Status {
			return errors.SuperAdminEditError
		}
	}

	// 获取用户所管理的部门
	ids, err := CurrentAdminTeamIds(ctx)
	if err != nil {
		return err
	}

	// 只允许更新当前部门的用户信息
	if !tools.InList(ids, user.TeamID) {
		return errors.NotEditTeamUserError
	}

	// 修改部门时，也只允许修改到自己所管辖的部门
	if in.TeamID != 0 && in.TeamID != user.TeamID && !tools.InList(ids, in.TeamID) {
		return errors.NotAddTeamUserError
	}

	if copier.Copy(&user, in) != nil {
		return errors.AssignError
	}

	return user.Update(ctx)
}

// UpdateCurrentUser 更新当前用户信息
func UpdateCurrentUser(ctx *core.Context, in *types.UpdateUserinfoRequest) error {
	md := ctx.Metadata()
	if md == nil {
		return errors.MetadataError
	}

	user := model.User{}
	if err := copier.Copy(&user, in); err != nil {
		return errors.AssignError
	}
	user.ID = md.UserID

	return user.Update(ctx)
}

// DeleteUser 删除用户信息
func DeleteUser(ctx *core.Context, in *types.DeleteUserRequest) error {
	// 超级管理员不允许删除
	if in.ID == 1 {
		return errors.SuperAdminDelError
	}

	user := model.User{}
	if user.OneByID(ctx, in.ID) != nil {
		return errors.DBNotFoundError
	}

	ids, err := CurrentAdminTeamIds(ctx)
	if err != nil {
		return err
	}

	// 只允许删除当前所管理部门的人员
	if !tools.InList(ids, user.TeamID) {
		return errors.NotDelTeamUserError
	}

	return user.DeleteByID(ctx, in.ID)
}

// UpdateUserinfoByVerify 更新用户重要数据
func UpdateUserinfoByVerify(ctx *core.Context, in *types.UpdateUserinfoByVerifyRequest) error {
	md := ctx.Metadata()
	if md == nil {
		return errors.MetadataError
	}

	// 判断验证码是否正确
	if err := ctx.EmailCaptcha("user").Verify(in.CaptchaID, in.Captcha); err != nil {
		return errors.CaptchaError
	}

	user := model.User{}
	user.ID = md.UserID
	user.Password = in.Password
	user.Phone = in.Phone
	user.Email = in.Email
	return user.Update(ctx)
}

// UserLogout 用户退出登陆
func UserLogout(ctx *core.Context) error {
	md := ctx.Metadata()
	if md == nil {
		return errors.MetadataError
	}
	return ctx.Jwt().Clear(md.UserID)
}

// UserLogin 用户登陆
func UserLogin(ctx *core.Context, in *types.UserLoginRequest) (resp *types.UserLoginResponse, err error) {
	resp = new(types.UserLoginResponse)
	defer func() {
		if !(errors.Is(err, errors.UserDisableError) ||
			errors.Is(err, errors.CaptchaError)) {
			_ = AddLoginLog(ctx, in.Phone, err)
		}
	}()

	// 判断验证码是否正确
	if err = ctx.ImageCaptcha(in.CaptchaName).Verify(in.CaptchaID, in.Captcha); err != nil {
		err = errors.CaptchaError
		return
	}

	// 密码解密
	passByte, _ := base64.StdEncoding.DecodeString(in.Password)
	decryptData, err := openssl.RSADecrypt(passByte, ctx.Cert().GetCert(decodePasswordCert))
	if err != nil {
		err = errors.RsaPasswordError
		return
	}

	// 序列化密码
	var pw types.Password
	if json.Unmarshal(decryptData, &pw) != nil {
		err = errors.RsaPasswordError
		return
	}

	// 判断当前时间戳是否过期,超过10s则拒绝
	if time.Now().UnixMilli()-pw.Time > 10*1000 {
		err = errors.PasswordExpireError
		return
	}

	in.Password = pw.Password

	// 通过手机号获取用户信息
	user := model.User{}
	if err = user.OneByPhone(ctx, in.Phone); err != nil {
		err = errors.UserNotFoundError
		return
	}

	// 由于屏蔽了password，需要调用指定方法查询
	password, err := user.PasswordByPhone(ctx, in.Phone)
	if err != nil {
		err = errors.UserNotFoundError
		return
	}

	// 用户被禁用则拒绝登陆
	if !*user.Status {
		err = errors.UserDisableError
		return
	}

	// 所属角色被禁用则拒绝登陆
	role := model.Role{}
	if !role.RoleStatus(ctx, user.RoleID) {
		return nil, errors.RoleDisableError
	}

	// 对比用户密码，错误则拒绝登陆
	if !tools.CompareHashPwd(password, in.Password) {
		err = errors.PasswordError
		return
	}

	// 生成登陆token
	if resp.Token, err = ctx.Jwt().Create(user.ID, &tps.Metadata{
		UserID:    user.ID,
		RoleID:    user.RoleID,
		RoleKey:   user.Role.Keyword,
		DataScope: user.Role.DataScope,
		Username:  user.Name,
		TeamID:    user.TeamID,
	}); err != nil {
		return nil, err
	}

	// 修改登陆时间
	return resp, user.UpdateLastLogin(ctx, time.Now().Unix())
}

// RefreshToken 用户刷新token
func RefreshToken(ctx *core.Context) (*types.UserLoginResponse, error) {
	md, err := ctx.Jwt().Parse()
	if md == nil {
		return nil, errors.MetadataError
	}

	if err == nil {
		return nil, errors.RefreshActiveTokenError
	}

	if !err.CanRenewal() {
		return nil, errors.RefTokenExpiredError
	}

	token, er := ctx.Jwt().Create(md.UserID, md)
	if er != nil {
		return nil, er
	}

	return &types.UserLoginResponse{
		Token: token,
	}, er
}

// CurrentUserMenuTree 获取当前用户的菜单树
func CurrentUserMenuTree(ctx *core.Context) (tree.Tree, error) {
	md := ctx.Metadata()
	if md == nil {
		return nil, errors.MetadataError
	}

	// 如果是超级管理员就直接返回全部菜单
	if md.RoleKey == consts.JwtSuperAdmin {
		return AllMenu(ctx)
	}

	// 查询角色所属菜单
	rm := model.RoleMenu{}
	rmList, err := rm.RoleMenus(ctx, md.RoleID)
	if err != nil {
		return nil, err
	}

	// 获取菜单的所有id
	var ids []int64
	for _, item := range rmList {
		ids = append(ids, item.MenuID)
	}

	// 获取指定id的所有菜单
	var menu model.Menu
	menuList, _ := menu.All(ctx, "id in ?", ids)
	var listTree []tree.Tree
	for _, item := range menuList {
		listTree = append(listTree, item)
	}

	return tree.BuildTree(listTree), nil
}
