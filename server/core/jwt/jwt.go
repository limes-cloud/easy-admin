package jwt

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	jv4 "github.com/golang-jwt/jwt/v4"
	"github.com/limeschool/easy-admin/server/config"
	rd "github.com/limeschool/easy-admin/server/core/redis"
	"github.com/limeschool/easy-admin/server/types"
	"strings"
	"time"
)

type jwt struct {
	redis *redis.Client
	conf  config.JWT
	ctx   *gin.Context
}

type JWT interface {
	Compare(userId int64) bool
	IsExist(userId int64) bool
	Store(userId int64, token string, duration time.Duration) error
	Clear(userId int64) error
	Parse() (*types.Metadata, *jwtErr)
	Create(userId int64, data *types.Metadata) (string, error)
	IsWhitelist(method, path string) bool
	CheckUnique(userID int64) bool
}

func New(conf config.JWT, rd rd.Redis, ctx *gin.Context) JWT {
	return &jwt{
		redis: rd.GetRedis(conf.Cache),
		ctx:   ctx,
		conf:  conf,
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
func (j jwt) Compare(userId int64) bool {
	token := j.ctx.GetHeader(j.conf.Header)
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
func (j jwt) Parse() (*types.Metadata, *jwtErr) {
	token := j.ctx.GetHeader(j.conf.Header)
	var m jv4.MapClaims = make(map[string]any)
	parser, err := jv4.ParseWithClaims(token, &m, func(token *jv4.Token) (interface{}, error) {
		return []byte(j.conf.Secret), nil
	})

	// 判断是否是验证失败
	if err != nil && parser == nil {
		return nil, j.newJwtErr("token verify error")
	}

	exp := int64(0)
	if m["exp"] != nil {
		exp = int64(m["exp"].(float64))
	}

	data := types.Metadata{}
	b, _ := json.Marshal(m["data"])
	_ = json.Unmarshal(b, &data)

	if err != nil {
		return &data, j.newJwtErr(err.Error(),
			withVerify(parser.Valid),
			withExpired(errors.Is(err, jv4.ErrTokenExpired)),
			withExpireUnix(exp),
			withRenewalUnix(int64(j.conf.Renewal.Seconds())),
		)
	}

	// 成功返回
	return &data, nil
}

// Create 生成token并保存到缓存
func (j jwt) Create(userId int64, data *types.Metadata) (string, error) {
	claims := make(jv4.MapClaims)
	claims["exp"] = time.Now().Unix() + int64(j.conf.Expire.Seconds())
	claims["iat"] = time.Now().Unix()
	claims["data"] = data
	tokenJwt := jv4.New(jv4.SigningMethodHS256)
	tokenJwt.Claims = claims
	token, err := tokenJwt.SignedString([]byte(j.conf.Secret))
	if err != nil {
		return "", err
	}
	return token, j.Store(userId, token, j.conf.Expire)
}

func (j jwt) IsWhitelist(method, path string) bool {
	return j.conf.Whitelist[strings.ToLower(method+":"+path)]
}

func (j jwt) CheckUnique(userID int64) bool {
	if !j.conf.Unique {
		return true
	}
	return j.Compare(userID)
}
