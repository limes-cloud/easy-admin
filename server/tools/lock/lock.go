package lock

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type lock struct {
	redis    *redis.Client
	key      string
	val      string
	duration time.Duration
}

type Lock interface {
	Acquire()
	TryAcquire() bool
	Release()
	AcquireFunc(f func() (any, error), do func() (any, error)) (any, error)
}

func NewLockWithDuration(redis *redis.Client, key string, duration time.Duration) Lock {
	return &lock{
		redis:    redis,
		key:      key,
		duration: duration,
	}
}

func NewLock(redis *redis.Client, key string) Lock {
	return &lock{
		redis:    redis,
		key:      key,
		duration: 30 * time.Second,
	}
}

func (l *lock) AcquireFunc(f func() (any, error), do func() (any, error)) (any, error) {
	for {
		// 获取数据
		data, err := f()
		if err == nil {
			return data, nil
		}

		// 数据不存在则去拿锁
		if res := l.redis.SetNX(context.TODO(), l.key, true, l.duration); res.Err() == nil && res.Val() {
			break
		}

		// 防止频繁自旋
		time.Sleep(100 * time.Millisecond)
	}

	return do()
}

// Acquire 获取分布式锁
func (l *lock) Acquire() {
	for {
		// 获得锁
		if res := l.redis.SetNX(context.TODO(), l.key, true, l.duration); res.Err() == nil && res.Val() {
			break
		}
		// 防止频繁自旋
		time.Sleep(100 * time.Millisecond)
	}
}

// TryAcquire 尝试获取锁，不会阻塞
func (l *lock) TryAcquire() bool {
	if res := l.redis.SetNX(context.TODO(), l.key, true, l.duration); res.Err() == nil && res.Val() {
		return true
	}
	return false
}

func (l *lock) Release() {
	l.redis.Del(context.TODO(), l.key)
}
