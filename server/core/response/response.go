package response

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/trace"
	"github.com/limeschool/easy-admin/server/types"
)

const (
	code    = 200
	errCode = 400
)

func Success(ctx *gin.Context) {
	ctx.JSON(code, &types.Response{
		Code:    code,
		Msg:     "success",
		TraceID: trace.GetID(ctx),
	})
}

func Data(ctx *gin.Context, data interface{}) {
	ctx.JSON(code, &types.Response{
		Code:    code,
		Msg:     "success",
		Data:    data,
		TraceID: trace.GetID(ctx),
	})
}

func List(ctx *gin.Context, page, pageSize, total int, data interface{}) {
	ctx.JSON(code, &types.ResponseList{
		Code:     code,
		Msg:      "success",
		Data:     data,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
		TraceID:  trace.GetID(ctx),
	})
}

func Error(ctx *gin.Context, err error) {
	if response, is := err.(*types.Response); is {
		response.TraceID = trace.GetID(ctx)
		ctx.JSON(code, response)
	} else {
		ctx.JSON(code, &types.Response{
			Code:    errCode,
			Msg:     err.Error(),
			TraceID: trace.GetID(ctx),
		})
	}
}
