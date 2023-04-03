package config

import (
	"flag"
	"github.com/limeschool/easy-admin/server/global"
	"github.com/spf13/viper"
	"os"
)

// Init 初始化配置
func Init() {
	vp := viper.New()

	// 解析命令行参数
	flagPath := flag.String("c", "", "the config file path")
	flag.Parse()

	// 读取配置文件
	if *flagPath != "" {
		vp.SetConfigType("yaml")
		vp.SetConfigFile(*flagPath)
		if err := vp.ReadInConfig(); err != nil {
			panic("配置初始化失败：" + err.Error())
		}
	} else {
		addr := os.Getenv("CONFIG_ADDR")
		if addr == "" {
			panic("环境变量CONFIG_ADDR未配置")
		}

		tp := os.Getenv("CONFIG_TYPE")
		if tp == "" {
			panic("环境变量CONFIG_TYPE未配置")
		}

		drive := os.Getenv("CONFIG_DRIVE")
		if addr == "" {
			panic("环境变量CONFIG_DRIVE未配置")
		}

		path := os.Getenv("CONFIG_PATH")
		if tp == "" {
			panic("环境变量CONFIG_PATH未配置")
		}

		if err := vp.AddRemoteProvider(drive, addr, path); err != nil {
			panic("配置仓库连接失败：" + err.Error())
		}

		if err := vp.ReadRemoteConfig(); err != nil {
			panic("配置初始化失败：" + err.Error())
		}
	}

	if err := vp.Unmarshal(&global.Config); err != nil {
		panic("配置初始化失败：" + err.Error())
	}
}
