package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core/response"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/global"
	"github.com/zeromicro/go-zero/core/load"
)

// CpuLoadShedding 自适应降载
func CpuLoadShedding() gin.HandlerFunc {
	conf := global.Config.Middleware.CupLoadShedding
	csd := load.NewAdaptiveShedder(
		load.WithCpuThreshold(int64(conf.Threshold)),
		load.WithBuckets(conf.Bucket),
		load.WithWindow(conf.Window),
	)
	return func(ctx *gin.Context) {
		promise, err := csd.Allow()
		if err != nil {
			response.Error(ctx, errors.ServerBusyError)
			ctx.Abort()
			return
		}

		defer func() {
			if ctx.Err() == context.DeadlineExceeded {
				promise.Fail()
			} else {
				promise.Pass()
			}
		}()
	}
}
