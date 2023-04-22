/**
 * @Author: 1280291001@qq.com
 * @Description: jwt 用于客户身份验证
 * @File: jwtErr
 * @Version: 1.0.0
 * @Date: 2023/4/18 22:30
 */

package jwt

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	jv4 "github.com/golang-jwt/jwt/v4"
	"time"
)

type jwt struct {
	redis *redis.Client
}

type Jwt interface {
	Compare(userId int64, token string) bool
	IsExist(userId int64) bool
	Store(userId int64, token string, duration time.Duration) error
	Clear(userId int64) error
	Parse(secret, token string) (any, *jwtErr)
	ParseWithUnmarshal(secret, token string, data any) *jwtErr
	Create(userId int64, expire time.Duration, secret string, data any) (string, error)
}

func NewJwt(cli *redis.Client) Jwt {
	return &jwt{
		redis: cli,
	}
}

func (j jwt) newJwtErr(err string, opts ...jwtErrOption) *jwtErr {
	je := &jwtErr{
		err: errors.New(err),
	}
	for _, opt := range opts {
		opt(je)
	}
	return je
}

// uuid 获取存储的唯一ID
func (j jwt) uuid(userId int64) string {
	return fmt.Sprintf("token_%x", j.encode(userId))
}

// uuid 对存储数据进行加密编码
func (j jwt) encode(data any) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprint(data))))
}

// Compare 对比token是否和缓存中的一致
func (j jwt) Compare(userId int64, token string) bool {
	st, err := j.redis.Get(context.Background(), j.uuid(userId)).Result()
	if err != nil {
		return false
	}
	return st == j.encode(token)
}

// IsExist 判断缓存中是否存在用户的token
func (j jwt) IsExist(userId int64) bool {
	st, _ := j.redis.Exists(context.Background(), j.uuid(userId)).Result()
	return st != 0
}

// Store 存储用户token信息到缓存数据
func (j jwt) Store(userId int64, token string, duration time.Duration) error {
	return j.redis.Set(context.Background(), j.uuid(userId), j.encode(token), duration).Err()
}

// Clear 清除用户缓存数据
func (j jwt) Clear(userId int64) error {
	return j.redis.Del(context.Background(), j.uuid(userId)).Err()
}

// Parse 解析用户token信息
func (j jwt) Parse(secret, token string) (any, *jwtErr) {
	var m jv4.MapClaims = make(map[string]any)
	parser, err := jv4.ParseWithClaims(token, &m, func(token *jv4.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	// 判断是否是验证失败
	if !parser.Valid {
		return nil, j.newJwtErr("token verify error")
	}

	exp := int64(0)
	if m["exp"] != nil {
		exp = int64(m["exp"].(float64))
	}

	if err != nil {
		return m["data"], j.newJwtErr(err.Error(),
			withVerify(parser.Valid),
			withExpired(errors.Is(err, jv4.ErrTokenExpired)),
			withExpireUnix(exp),
		)
	}

	// 成功返回
	return m["data"], nil
}

// ParseWithUnmarshal 解析并序列化
func (j jwt) ParseWithUnmarshal(secret, token string, data any) *jwtErr {
	parseData, err := j.Parse(secret, token)
	if parseData != nil {
		b, _ := json.Marshal(parseData)
		_ = json.Unmarshal(b, data)
	}
	return err
}

// Create 生成token并保存到缓存
func (j jwt) Create(userId int64, expire time.Duration, secret string, data any) (string, error) {
	claims := make(jv4.MapClaims)
	claims["exp"] = time.Now().Unix() + int64(expire.Seconds())
	claims["iat"] = time.Now().Unix()
	claims["data"] = data
	tokenJwt := jv4.New(jv4.SigningMethodHS256)
	tokenJwt.Claims = claims
	token, err := tokenJwt.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, j.Store(userId, token, expire)
}
