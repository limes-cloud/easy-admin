package config

import "time"

type Orm struct {
	Enable        bool          `mapstructure:"enable"`        //是否启用数据库
	Drive         string        `mapstructure:"drive"`         //数据库类型
	Name          string        `mapstructure:"name"`          //数据库名字【代称】
	Dsn           string        `mapstructure:"dsn"`           //数据库dsn
	MaxLifetime   time.Duration `mapstructure:"maxLifetime"`   //连接最大存活时长
	MaxOpenConn   int           `mapstructure:"maxOpenConn"`   //最大连接数
	MaxIdleConn   int           `mapstructure:"maxIdleConn"`   //最大空闲连接数
	Level         int           `mapstructure:"level"`         //日志打印等级
	SlowThreshold time.Duration `mapstructure:"slowThreshold"` //慢查询阈值
	Replicas      []string      `mapstructure:"replicas"`      //读服务器配置
}
