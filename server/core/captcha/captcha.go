package captcha

import (
	"crypto/md5"
	"fmt"
	"github.com/limeschool/easy-admin/server/config"
	e "github.com/limeschool/easy-admin/server/core/email"
	"github.com/limeschool/easy-admin/server/core/redis"
	"sync"

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
	mu    sync.RWMutex
	cache redis.Redis
	email e.Email
	m     map[string]config.Captcha
}

type Captcha interface {
	Image(ip, name string) Image
	Email(ip, name string) Email
}

func New(cs []config.Captcha, rs redis.Redis, email e.Email) Captcha {
	cpIns := captcha{
		cache: rs,
		email: email,
		m:     make(map[string]config.Captcha),
		mu:    sync.RWMutex{},
	}

	cpIns.mu.Lock()
	defer cpIns.mu.Unlock()

	for _, item := range cs {
		cpIns.m[item.Name+":"+item.Type] = item
	}
	return &cpIns
}

// getTemplate 获取指定模板
func (c *captcha) getTemplate(name, tp string) (config.Captcha, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	temp, is := c.m[name+":"+tp]
	return temp, is
}

// cid 获取用户对应验证场景的唯一id
func (c *captcha) cid(ip, name, tp string) string {
	return fmt.Sprintf("captcha:%s:%s:%x", name, tp, md5.Sum([]byte(ip)))
}

// randomCode 生成随机数验证码
func (c *captcha) randomCode(len int) string {
	rand.Seed(time.Now().Unix())
	var code = rand.Intn(int(math.Pow10(len)) - int(math.Pow10(len-1)))
	return strconv.Itoa(code + int(math.Pow10(len-1)))
}

// Image  实例化图形验证码
func (c *captcha) Image(ip, name string) Image {
	return &image{
		name:    name,
		tp:      "image",
		captcha: c,
		ip:      ip,
	}
}

// Email  实例化邮箱验证码
func (c *captcha) Email(ip, name string) Email {
	return &email{
		name:    name,
		tp:      "email",
		captcha: c,
		ip:      ip,
	}
}
