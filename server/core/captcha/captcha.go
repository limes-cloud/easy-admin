package captcha

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/config"
	e "github.com/limeschool/easy-admin/server/core/email"
	"github.com/limeschool/easy-admin/server/core/redis"

	"math"
	"math/rand"
	"strconv"
	"time"
)

type res struct {
	ID     string `json:"id"`
	Base64 string `json:"base64,omitempty"`
	Expire int    `json:"expire"`
}

type captcha struct {
	rs    redis.Redis
	email e.Email
	m     map[string]config.Captcha
}

type Captcha interface {
	Image(ctx *gin.Context, name string) Image
	Email(ctx *gin.Context, name string) Email
}

func New(cs []config.Captcha, rs redis.Redis, email e.Email) Captcha {
	cpIns := captcha{
		rs:    rs,
		email: email,
		m:     make(map[string]config.Captcha),
	}
	for _, item := range cs {
		cpIns.m[item.Name+":"+item.Type] = item
	}
	return &cpIns
}

// isTemplate 判断是否存在指定的模板
func (c *captcha) isTemplate(name, tp string) bool {
	_, is := c.m[name+":"+tp]
	return is
}

// getTemplate 获取指定模板
func (c *captcha) getTemplate(name, tp string) config.Captcha {
	return c.m[name+":"+tp]
}

// clientUUID 获取用户对应验证场景的唯一id
func (c *captcha) clientUUID(ctx *gin.Context, name, tp string) string {
	ip := ctx.ClientIP()
	if ip == "::1" {
		ip = ctx.GetHeader("X-Real-IP")
	}
	return fmt.Sprintf("captcha:%s:%s:%x", name, tp, md5.Sum([]byte(ip)))
}

// randomCode 生成随机数验证码
func (c *captcha) randomCode(len int) string {
	rand.Seed(time.Now().Unix())
	var code = rand.Intn(int(math.Pow10(len)) - int(math.Pow10(len-1)))
	return strconv.Itoa(code + int(math.Pow10(len-1)))
}

// Image  实例化图形验证码
func (c *captcha) Image(ctx *gin.Context, name string) Image {
	return &image{
		name:    name,
		tp:      "image",
		captcha: c,
		ctx:     ctx,
	}
}

// Email  实例化邮箱验证码
func (c *captcha) Email(ctx *gin.Context, name string) Email {
	return &email{
		name:    name,
		tp:      "email",
		captcha: c,
		ctx:     ctx,
	}
}
