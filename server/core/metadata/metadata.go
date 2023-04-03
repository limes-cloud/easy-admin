package metadata

import (
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"github.com/limeschool/easy-admin/server/errors"
)

const (
	JwtMapClaimsKey = "jwt-claims"
)

type Value struct {
	UserID    int64  `json:"UserID"`
	RoleID    int64  `json:"RoleID"`
	RoleKey   string `json:"RoleKey"`
	Username  string `json:"Username"`
	DataScope string `json:"DataScope"`
	TeamID    int64  `json:"TeamID"`
}

func GetFormContext(ctx *gin.Context) (*Value, error) {
	val, is := ctx.Get(JwtMapClaimsKey)
	if !is {
		return nil, errors.TokenEmptyError
	}

	meta, is := val.(*Value)
	if !is {
		return nil, errors.TokenDataError
	}
	return meta, nil
}

func (m *Value) SetToContext(ctx *gin.Context) {
	ctx.Set(JwtMapClaimsKey, m)
}

func Parse(m any) (*Value, error) {
	meta := Value{}

	// 序列化
	b, err := json.Marshal(m)
	if err != nil {
		return nil, errors.TokenDataError
	}

	// 反序列化
	if json.Unmarshal(b, &meta) != nil {
		return nil, errors.TokenDataError
	}
	return &meta, nil
}
