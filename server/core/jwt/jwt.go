package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/limeschool/easy-admin/server/core/metadata"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/global"
	"github.com/limeschool/easy-admin/server/tools"
	"time"
)

const (
	prefix  = "token_"
	metaKey = "data"
)

func Compare(ctx *gin.Context, userId int64, token string) bool {
	key := prefix + tools.Md5(userId)
	st, err := global.Redis.Get(ctx, key).Result()
	if err != nil {
		return false
	}
	return st == tools.Md5(token)
}

func Store(ctx *gin.Context, userId int64, token string, duration time.Duration) error {
	key := prefix + tools.Md5(userId)
	return global.Redis.Set(ctx, key, tools.Md5(token), duration).Err()
}

func Clear(ctx *gin.Context, userId int64) error {
	key := prefix + tools.Md5(userId)
	return global.Redis.Del(ctx, key).Err()
}

func Parse(secret, token string) (any, error) {
	var m jwt.MapClaims
	parser, err := jwt.ParseWithClaims(token, &m, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !parser.Valid {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.TokenExpiredError
		}
		return nil, errors.TokenValidateError
	}

	return m[metaKey], nil
}

func Create(ctx *gin.Context, md *metadata.Value) (string, error) {
	// 生成token携带信息
	conf := global.Config.Middleware.Jwt

	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Unix() + int64(conf.Expire.Seconds())
	claims["iat"] = time.Now().Unix()
	claims[metaKey] = md
	tokenJwt := jwt.New(jwt.SigningMethodHS256)
	tokenJwt.Claims = claims
	token, err := tokenJwt.SignedString([]byte(conf.Secret))
	if err != nil {
		return "", err
	}

	return token, Store(ctx, md.UserID, token, conf.Expire)
}

func ParseMapClaimsAndExpired(ctx *gin.Context) (any, bool, bool) {
	var claims jwt.MapClaims
	conf := global.Config.Middleware.Jwt

	// 解密
	token := ctx.Request.Header.Get(conf.Header)
	_, _ = jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.Secret), nil
	})

	// 解密失败
	if claims == nil {
		return nil, true, true
	}

	// 获取时间差
	unix := time.Now().Unix() - int64(claims["iat"].(float64))
	return claims[metaKey], unix > int64(conf.Expire.Seconds()), unix > int64(conf.Expire.Seconds()+conf.Renewal.Seconds())
}
