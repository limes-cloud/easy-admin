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
	AcquireFunc(f func() error, do func() error) error
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

// AcquireFunc 分布式锁 f()从redis获取 do()从mysql获取
func (l *lock) AcquireFunc(f func() error, do func() error) error {
	for {
		// 获取数据
		if err := f(); err == nil {
			return nil
		}

		// 数据不存在则去拿锁
		if res := l.redis.SetNX(context.TODO(), l.key, true, l.duration); res.Err() == nil && res.Val() {
			break
		}

		// 防止频繁自旋
		time.Sleep(1 * time.Millisecond)
	}

	// 在获取一下redis数据，防止重复读取数据库
	if err := f(); err == nil {
		return nil
	}
	return do()
}

// Acquire 获取分布式锁，不建议直接使用此方法，可以使用AcquireFunc
func (l *lock) Acquire() {
	for {
		// 获得锁
		if res := l.redis.SetNX(context.TODO(), l.key, true, l.duration); res.Err() == nil && res.Val() {
			break
		}
	}
}

// TryAcquire 尝试获取锁，不会阻塞
func (l *lock) TryAcquire() bool {
	if res := l.redis.SetNX(context.TODO(), l.key, true, l.duration); res.Err() == nil && res.Val() {
		return true
	}
	return false
}

// Release 释放锁
func (l *lock) Release() {
	l.redis.Del(context.TODO(), l.key)
}
