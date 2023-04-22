package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/limeschool/easy-admin/server/config"
)

type rd struct {
	m map[string]*redis.Client
}

type Redis interface {
	Get(name string) (*redis.Client, error)
	GetRedis(name string) *redis.Client
}

func New(rc []config.Redis) Redis {
	redisIns := rd{
		m: make(map[string]*redis.Client),
	}
	for _, conf := range rc {
		if !conf.Enable {
			continue
		}

		client := redis.NewClient(&redis.Options{
			Addr:     conf.Host,
			Username: conf.Username,
			Password: conf.Password,
		})
		if err := client.Ping(context.TODO()).Err(); err != nil {
			panic(fmt.Sprintf("redis 初始化失败:%v", err.Error()))
		}

		redisIns.m[conf.Name] = client
	}
	return &redisIns
}

func (o *rd) Get(name string) (*redis.Client, error) {
	if o.m[name] == nil {
		return nil, errors.New("not exist db")
	}
	return o.m[name], nil
}

func (o *rd) GetRedis(name string) *redis.Client {
	return o.m[name]
}
