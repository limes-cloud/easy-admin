package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/internal/system/service"
)

func Config(c *gin.Context) {
	ctx := core.New(c)
	defer ctx.Release()

	ctx.RespData(service.Config(ctx))

}
