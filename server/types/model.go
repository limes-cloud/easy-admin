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
