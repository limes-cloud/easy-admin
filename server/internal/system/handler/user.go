package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/response"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/internal/system/service"
	"github.com/limeschool/easy-admin/server/internal/system/types"
)

func PageUser(ctx *gin.Context) {
	// 检验参数
	in := types.PageUserRequest{}
	if ctx.ShouldBind(&in) != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}

	if list, total, err := service.PageUser(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.List(ctx, in.Page, len(list), int(total), list)
	}
}

func CurUser(ctx *gin.Context) {
	if user, err := service.CurrentUser(ctx); err != nil {
		response.Error(ctx, err)
	} else {
		response.Data(ctx, user)
	}
}

func AddUser(ctx *gin.Context) {
	// 检验参数
	in := types.AddUserRequest{}
	if err := ctx.ShouldBindJSON(&in); err != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}
	// 调用实现
	if err := service.AddUser(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.Success(ctx)
	}
}

func UpdateUser(ctx *gin.Context) {
	// 检验参数
	in := types.UpdateUserRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}
	// 调用实现
	if err := service.UpdateUser(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.Success(ctx)
	}
}

func UpdateUserinfo(ctx *gin.Context) {
	// 检验参数
	in := types.UpdateUserinfoRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}
	// 调用实现
	if err := service.UpdateCurrentUser(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.Success(ctx)
	}
}

func UpdateUserinfoByVerify(ctx *gin.Context) {
	// 检验参数
	in := types.UpdateUserinfoByVerifyRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}
	// 调用实现
	if err := service.UpdateUserinfoByVerify(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.Success(ctx)
	}
}

func DeleteUser(ctx *gin.Context) {
	// 检验参数
	in := types.DeleteUserRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}
	// 调用实现
	if err := service.DeleteUser(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.Success(ctx)
	}
}

func UserLogin(ctx *gin.Context) {
	// 检验参数
	in := types.UserLoginRequest{}
	if ctx.ShouldBindJSON(&in) != nil {
		response.Error(ctx, errors.ParamsError)
		return
	}
	// 调用实现
	if resp, err := service.UserLogin(ctx, &in); err != nil {
		response.Error(ctx, err)
	} else {
		response.Data(ctx, resp)
	}
}

func RefreshToken(ctx *gin.Context) {
	if resp, err := service.RefreshToken(ctx); err != nil {
		response.Error(ctx, err)
	} else {
		response.Data(ctx, resp)
	}
}

func UserLogout(ctx *gin.Context) {
	if err := service.UserLogout(ctx); err != nil {
		response.Error(ctx, err)
	} else {
		response.Success(ctx)
	}
}

// UserMenus 获取用户的菜单列表
func UserMenus(ctx *gin.Context) {
	if tree, err := service.CurrentUserMenuTree(ctx); err != nil {
		response.Error(ctx, err)
	} else {
		response.Data(ctx, tree)
	}
}
