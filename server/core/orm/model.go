package orm

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/types"
	"gorm.io/gorm"
	"strings"
	"time"
)

type PageOptions struct {
	Page     int
	PageSize int
	Model    any
	Scopes   func(db *gorm.DB) *gorm.DB
}

type AllOptions struct {
	Model  any
	Scopes func(db *gorm.DB) *gorm.DB
}

type CreateModel struct {
	ID        int64 `json:"id" gorm:"primary_key"`
	CreatedAt int64 `json:"created_at,omitempty" gorm:"index"`
}

type BaseModel struct {
	ID        int64 `json:"id" gorm:"primary_key"`
	CreatedAt int64 `json:"created_at,omitempty" gorm:"index"`
	UpdatedAt int64 `json:"updated_at,omitempty" gorm:"index"`
}

type DeleteModel struct {
	ID        int64          `json:"id" gorm:"primary_key"`
	CreatedAt int64          `json:"created_at,omitempty" gorm:"index"`
	UpdatedAt int64          `json:"updated_at,omitempty" gorm:"index"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// TransferErr 将数据库的错误转换成中文
func TransferErr(m map[string]string, err error) error {
	if err == nil {
		return nil
	}

	if customErr, ok := err.(*types.Response); ok {
		return customErr
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.DBNotFoundError
	}

	if strings.Contains(err.Error(), "Duplicate") {
		str := err.Error()
		str = strings.ReplaceAll(str, "'", "")
		str = strings.TrimPrefix(str, "Error 1062: Duplicate entry ")
		arr := strings.Split(str, " for key ")
		return errors.NewF(`%v "%v" 已存在`, m[arr[1]], arr[0])
	}

	if strings.Contains(err.Error(), "FOREIGN KEY") {
		return errors.NewF(`数据正在被使用中，无法删除`)
	}

	return errors.DBError
}

// DelayDelCache 数据延迟双删
func DelayDelCache(redis *redis.Client, key string) {
	redis.Del(context.Background(), key)
	go func() {
		time.Sleep(1 * time.Second)
		redis.Del(context.Background(), key)
	}()
}
