package config

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"time"
)

// Config 系统配置
type Config struct {
	v          *viper.Viper
	isRemote   bool
	Service    Service
	Log        Log
	Orm        []Orm
	Middleware Middleware
	Redis      []Redis
	Enforcer   Enforcer
	JWT        JWT
	Cert       []Cert
	Email      Email
	File       File
	Http       Http
	Captcha    []Captcha
}

// New 初始化配置
func New() *Config {
	flagPath := flag.String("c", "", "the config file path")
	flag.Parse()

	vp := viper.New()
	conf := Config{
		v: vp,
	}

	if *flagPath != "" {
		// 读取本地配置文件
		vp.SetConfigType("yaml")
		vp.SetConfigFile(*flagPath)
		if err := vp.ReadInConfig(); err != nil {
			panic("配置初始化失败：" + err.Error())
		}
	} else {
		// 读取远程配置文件，支持consul\etcd
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
		if path == "" {
			panic("环境变量CONFIG_PATH未配置")
		}

		if err := vp.AddRemoteProvider(drive, addr, path); err != nil {
			panic("配置仓库连接失败：" + err.Error())
		}

		vp.SetConfigType(tp)
		if err := vp.ReadRemoteConfig(); err != nil {
			panic("配置初始化失败：" + err.Error())
		}
	}

	// 是否为远程获取配置
	conf.isRemote = *flagPath == ""

	if err := vp.Unmarshal(&conf); err != nil {
		panic("配置初始化失败：" + err.Error())
	}

	return &conf
}

// Watch 配置变更监听
func (c *Config) Watch(f func(c *Config)) {
	c.v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("on change")
		if err := c.v.Unmarshal(&c); err != nil {
			fmt.Printf("配置变更失败：%v", err.Error())
		} else {
			fmt.Printf("配置变更成功！")
			f(c)
		}
	})
	if c.isRemote {
		go func() {
			for {
				//delay after each request
				time.Sleep(time.Second * 5)
				if err := viper.WatchRemoteConfig(); err != nil {
					_ = fmt.Errorf("unable to read remote config: %v", err)
					continue
				}
			}
		}()
	} else {
		c.v.WatchConfig()
	}
}
