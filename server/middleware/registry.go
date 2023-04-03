package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/global"
)

func Registry(engine *gin.Engine) *gin.RouterGroup {
	conf := global.Config.Middleware

	// 开启全局404
	engine.NoRoute(Resp404())

	// 开启健康检查
	engine.GET("/healthy", Healthy())

	// 开启跨域
	if conf.Cors.Enable {
		engine.Use(Cors())
	}

	// 开启链路id和全局异常捕捉恢复
	engine.Use(Trace(), Recovery())

	// 开启pprof
	if conf.Pprof.Enable {
		PprofApi(engine)
	}

	api := engine.Group("/api")
	// 开启请求日志
	if conf.RequestLog.Enable {
		api.Use(RequestLog())
	}

	// 开启全局限流
	if conf.RateLimit.Enable {
		api.Use(RateLimit())
	}

	// 开启ip限流
	if conf.IpLimit.Enable {
		api.Use(IpLimit())
	}

	// 开启自适应限流
	if conf.CupLoadShedding.Enable {
		api.Use(CpuLoadShedding())
	}

	// 开启jwt验证
	if conf.Jwt.Enable {
		api.Use(JwtAuth())
	}

	// 开启casbin鉴权
	if conf.Casbin.Enable {
		api.Use(Casbin())
	}

	return api
}
