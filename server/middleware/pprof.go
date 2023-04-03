package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/global"
	"net/http"
	"net/http/pprof"
)

// PprofApi pprof 分析工具
func PprofApi(group *gin.Engine) {
	prefix := "/debug/pprof"

	secret := global.Config.Middleware.Pprof.Secret
	api := group.Group(prefix)
	api.GET("/", pprofServer(pprof.Index, secret))
	api.GET("/cmdline", pprofServer(pprof.Cmdline, secret))
	api.GET("/profile", pprofServer(pprof.Profile, secret))
	api.GET("/symbol", pprofServer(pprof.Symbol, secret))
	api.POST("/symbol", pprofServer(pprof.Symbol, secret))
	api.GET("/trace", pprofServer(pprof.Trace, secret))
	api.GET("/allocs", pprofServer(pprof.Handler("allocs").ServeHTTP, secret))
	api.GET("/block", pprofServer(pprof.Handler("block").ServeHTTP, secret))
	api.GET("/goroutine", pprofServer(pprof.Handler("goroutine").ServeHTTP, secret))
	api.GET("/heap", pprofServer(pprof.Handler("heap").ServeHTTP, secret))
	api.GET("/mutex", pprofServer(pprof.Handler("mutex").ServeHTTP, secret))
	api.GET("/threadcreate", pprofServer(pprof.Handler("threadcreate").ServeHTTP, secret))
}

func pprofServer(handler http.HandlerFunc, secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Query(global.Config.Middleware.Pprof.Query) == secret {
			handler.ServeHTTP(ctx.Writer, ctx.Request)
		}
	}
}
