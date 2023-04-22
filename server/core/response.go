package core

import (
	"github.com/limeschool/easy-admin/server/types"
)

// RespSuccess
//
//	@Description: 返回成功
func (ctx *Context) RespSuccess() {
	ctx.JSON(200, &types.Response{
		Code:    ctx.Config().Service.SuccessCode,
		Msg:     "success",
		TraceID: ctx.TraceID(),
	})
}

// RespData
//
//	@Description: 返回成功并且携带数据
//	@param data 成功返回的数据
func (ctx *Context) RespData(data any) {
	ctx.JSON(200, &types.Response{
		Code:    ctx.Config().Service.SuccessCode,
		Msg:     "success",
		Data:    data,
		TraceID: ctx.TraceID(),
	})
}

// RespList
//
//	@Description: 返回成功并且携带数据列表数量，用于分页查询
//	@param total 总的数量条数
//	@param data  分页查询的数据
func (ctx *Context) RespList(total int64, data any) {
	ctx.JSON(200, &types.ResponseList{
		Code:    ctx.Config().Service.SuccessCode,
		Msg:     "success",
		Data:    data,
		Total:   total,
		TraceID: ctx.TraceID(),
	})
}

// RespError
//
//	@Description: 返回数据错误信息
//	@param err
func (ctx *Context) RespError(err error) {
	if response, is := err.(*types.Response); is {
		response.TraceID = ctx.TraceID()
		ctx.JSON(200, response)
	} else {
		ctx.JSON(200, &types.Response{
			Code:    ctx.Config().Service.ErrorCode,
			Msg:     err.Error(),
			TraceID: ctx.TraceID(),
		})
	}
}
