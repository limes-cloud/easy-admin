package service

import (
	"github.com/jinzhu/copier"
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/internal/system/model"
	"github.com/limeschool/easy-admin/server/internal/system/types"
	tps "github.com/limeschool/easy-admin/server/types"
)

// AllFile 获取通知分页信息
func AllFile(ctx *core.Context, in *types.AllFileRequest) ([]*model.File, error) {
	md := ctx.Metadata()
	if md == nil {
		return nil, errors.MetadataError
	}

	file := model.File{}
	return file.All(ctx, tps.AllOptions{
		Model: in,
	})

}

// AddFile 新增通知信息
func AddFile(ctx *core.Context, in *types.AddFileRequest) error {
	file := model.File{}
	if copier.Copy(&file, in) != nil {
		return errors.AssignError
	}
	return file.Create(ctx)
}

// UpdateFile 更新通知信息
func UpdateFile(ctx *core.Context, in *types.UpdateFileRequest) error {
	notice := model.File{}
	if copier.Copy(&notice, in) != nil {
		return errors.AssignError
	}
	return notice.Update(ctx)
}

// DeleteFile 删除通知信息
func DeleteFile(ctx *core.Context, in *types.DeleteFileRequest) error {
	notice := model.File{}
	return notice.DeleteByID(ctx, in.ID)
}
