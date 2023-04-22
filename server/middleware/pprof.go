package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core"
	"net/http"
	"net/http/pprof"
)

// PprofApi pprof 分析工具
func PprofApi(group *gin.Engine) {
	prefix := "/debug/pprof"

	secret := core.GlobalConfig().Middleware.Pprof.Secret
	query := core.GlobalConfig().Middleware.Pprof.Query

	api := group.Group(prefix)
	api.GET("/", pprofServer(pprof.Index, secret, query))
	api.GET("/cmdline", pprofServer(pprof.Cmdline, secret, query))
	api.GET("/profile", pprofServer(pprof.Profile, secret, query))
	api.GET("/symbol", pprofServer(pprof.Symbol, secret, query))
	api.POST("/symbol", pprofServer(pprof.Symbol, secret, query))
	api.GET("/trace", pprofServer(pprof.Trace, secret, query))
	api.GET("/allocs", pprofServer(pprof.Handler("allocs").ServeHTTP, secret, query))
	api.GET("/block", pprofServer(pprof.Handler("block").ServeHTTP, secret, query))
	api.GET("/goroutine", pprofServer(pprof.Handler("goroutine").ServeHTTP, secret, query))
	api.GET("/heap", pprofServer(pprof.Handler("heap").ServeHTTP, secret, query))
	api.GET("/mutex", pprofServer(pprof.Handler("mutex").ServeHTTP, secret, query))
	api.GET("/threadcreate", pprofServer(pprof.Handler("threadcreate").ServeHTTP, secret, query))
}

func pprofServer(handler http.HandlerFunc, secret, query string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Query(query) == secret {
			handler.ServeHTTP(ctx.Writer, ctx.Request)
		}
	}
}
