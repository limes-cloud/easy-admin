package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/limeschool/easy-admin/server/global"
)

func Init() {
	conf := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Host,
		Username: conf.Username,
		Password: conf.Password,
	})
	if err := client.Ping(context.TODO()).Err(); err != nil {
		panic(fmt.Sprintf("redis 初始化失败:%v", err.Error()))
	}
	global.Redis = client
}
