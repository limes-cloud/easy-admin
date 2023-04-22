package service

import (
	"github.com/jinzhu/copier"
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/internal/system/model"
	"github.com/limeschool/easy-admin/server/internal/system/types"
	"github.com/limeschool/easy-admin/server/tools"
	"github.com/limeschool/easy-admin/server/tools/tree"
)

func AllTeam(ctx *core.Context) (tree.Tree, error) {
	team := model.Team{}
	return team.Tree(ctx)
}

func AddTeam(ctx *core.Context, in *types.AddTeamRequest) error {
	team := model.Team{}

	// 获取用户所管理的部门
	ids, err := CurrentAdminTeamIds(ctx)
	if err != nil {
		return err
	}

	if !tools.InList(ids, in.ParentID) {
		return errors.NotAddTeamError
	}

	if copier.Copy(&team, in) != nil {
		return errors.AssignError
	}

	return team.Create(ctx)
}

func UpdateTeam(ctx *core.Context, in *types.UpdateTeamRequest) error {
	team := model.Team{}
	if in.ParentID != 0 && in.ID == in.ParentID {
		return errors.TeamParentIdError
	}

	// 获取用户所管理的部门
	ids, err := CurrentAdminTeamIds(ctx)
	if err != nil {
		return err
	}

	if !tools.InList(ids, in.ID) {
		return errors.NotEditTeamError
	}
	if copier.Copy(&team, in) != nil {
		return errors.AssignError
	}
	return team.Update(ctx)
}

func DeleteTeam(ctx *core.Context, in *types.DeleteTeamRequest) error {
	team := model.Team{}

	// 获取用户所管理的部门
	ids, err := CurrentAdminTeamIds(ctx)
	if err != nil {
		return err
	}

	if !tools.InList(ids, in.ID) {
		return errors.NotDelTeamError
	}
	return team.DeleteByID(ctx, in.ID)
}
