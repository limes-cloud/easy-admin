package types

import (
	"gorm.io/gorm"
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
	ID        int64 `json:"id" gorm:"primary_key;autoIncrement;size:32;comment:主键ID"`
	CreatedAt int64 `json:"created_at,omitempty" gorm:"index;comment:创建时间"`
}

type BaseModel struct {
	ID        int64 `json:"id" gorm:"primary_key;autoIncrement;size:32;comment:主键ID"`
	CreatedAt int64 `json:"created_at,omitempty" gorm:"index;comment:创建时间"`
	UpdatedAt int64 `json:"updated_at,omitempty" gorm:"index;comment:修改时间"`
}

type DeleteModel struct {
	ID        int64          `json:"id" gorm:"primary_key;autoIncrement;size:32;comment:主键ID"`
	CreatedAt int64          `json:"created_at,omitempty" gorm:"index;comment:创建时间"`
	UpdatedAt int64          `json:"updated_at,omitempty" gorm:"index;comment:修改时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at;index;comment:删除时间"`
}
