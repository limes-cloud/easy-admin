package lock

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

// TestLock 基准测试
func BenchmarkLock(b *testing.B) {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	if err := client.Ping(context.TODO()).Err(); err != nil {
		panic(fmt.Sprintf("redis 初始化失败:%v", err.Error()))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func() {
			lock := NewLock(client, "test1")
			err := lock.AcquireFunc(func() error {
				_, err := client.Get(context.TODO(), "data").Result()
				if err != nil {
					return err
				}
				return nil
			}, func() error {
				// todo get data
				b.Log("do set data")
				return client.Set(context.TODO(), "data", "1", 100*time.Second).Err()
			})
			if err != nil {
				b.Log(err)
			}
			lock.Release()
		}()
	}
	b.StopTimer()
}
