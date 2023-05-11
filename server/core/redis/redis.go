package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/limeschool/easy-admin/server/config"
	"sync"
)

type rd struct {
	mu sync.RWMutex
	m  map[string]*redis.Client
}

type Redis interface {
	// Get
	//
	//	@Description: 获取指定名称的redis实例，如果实例不存在则会返回报错
	//	@param name 实例名称
	//	@return *redis.Client
	//	@return error
	Get(name string) (*redis.Client, error)
	// GetRedis
	//
	//	@Description: 获取指定名称的redis实例，如果实例不存在则会nil
	//	@param name 实例名称
	//	@return *redis.Client
	GetRedis(name string) *redis.Client
}

func New(rc []config.Redis) Redis {
	redisIns := rd{
		mu: sync.RWMutex{},
		m:  make(map[string]*redis.Client),
	}

	redisIns.mu.Lock()
	defer redisIns.mu.Unlock()

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
	o.mu.RLock()
	defer o.mu.RUnlock()

	if o.m[name] == nil {
		return nil, errors.New("not exist redis")
	}
	return o.m[name], nil
}

func (o *rd) GetRedis(name string) *redis.Client {
	o.mu.RLock()
	defer o.mu.RUnlock()

	return o.m[name]
}
