package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/internal/system/service"
	"github.com/limeschool/easy-admin/server/internal/system/types"
)

func PageUser(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.PageUserRequest{}
	if ctx.ShouldBind(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}

	if list, total, err := service.PageUser(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespList(total, list)
	}
}

func CurUser(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	if user, err := service.CurrentUser(ctx); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(user)
	}
}

func AddUser(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.AddUserRequest{}
	if err := ctx.ShouldBindJSON(&in); err != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	// 调用实现
	if err := service.AddUser(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func UpdateUser(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.UpdateUserRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	// 调用实现
	if err := service.UpdateUser(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func UpdateUserinfo(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.UpdateUserinfoRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	// 调用实现
	if err := service.UpdateCurrentUser(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func UpdateUserinfoByVerify(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.UpdateUserinfoByVerifyRequest{
		CaptchaName: "user",
	}
	if ctx.ShouldBindJSON(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	// 调用实现
	if err := service.UpdateUserinfoByVerify(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func DeleteUser(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.DeleteUserRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	// 调用实现
	if err := service.DeleteUser(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

func UserLogin(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	in := types.UserLoginRequest{
		CaptchaName: "login",
	}
	if ctx.ShouldBindJSON(&in) != nil {
		ctx.RespError(errors.ParamsError)
		return
	}
	// 调用实现
	if resp, err := service.UserLogin(ctx, &in); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

func RefreshToken(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	if resp, err := service.RefreshToken(ctx); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(resp)
	}
}

func UserLogout(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	if err := service.UserLogout(ctx); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespSuccess()
	}
}

// UserMenus 获取用户的菜单列表
func UserMenus(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	if tree, err := service.CurrentUserMenuTree(ctx); err != nil {
		ctx.RespError(err)
	} else {
		ctx.RespData(tree)
	}
}
