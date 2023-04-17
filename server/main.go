package main

import (
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/global"
	systemRouter "github.com/limeschool/easy-admin/server/internal/system/router"
	"github.com/limeschool/easy-admin/server/middleware"
	"log"
)

func main() {

	// 创建gin引擎
	engine := gin.New()

	// 核心组件初始化
	core.Init()

	// 注册中间件
	api := middleware.Registry(engine)

	// 初始化静态资源
	engine.Static("/static", "./static")

	// 路由初始化
	systemRouter.Init(api)

	// 启动服务
	service := global.Config.Service
	if err := engine.Run(service.Addr); err != nil {
		log.Fatalln(err)
	}
}
