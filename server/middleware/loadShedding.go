package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/zeromicro/go-zero/core/load"
)

// CpuLoadShedding 自适应降载
func CpuLoadShedding() gin.HandlerFunc {
	conf := core.GlobalConfig().Middleware.CupLoadShedding
	csd := load.NewAdaptiveShedder(
		load.WithCpuThreshold(int64(conf.Threshold)),
		load.WithBuckets(conf.Bucket),
		load.WithWindow(conf.Window),
	)
	return func(c *gin.Context) {
		ctx := core.New(c)
		defer ctx.Release()

		promise, err := csd.Allow()
		if err != nil {
			ctx.RespError(errors.ServerBusyError)
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
