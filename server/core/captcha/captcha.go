package captcha

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/limeschool/easy-admin/server/core/email"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/limeschool/easy-admin/server/global"
	"github.com/limeschool/easy-admin/server/tools"
	captcha "github.com/mojocn/base64Captcha"
	"time"
)

type custom struct {
	redis    *redis.Client
	duration time.Duration
}

func (s *custom) Set(id string, value string) error {
	return s.redis.Set(context.Background(), id, value, s.duration).Err()
}
func (s *custom) Get(id string, clear bool) string {
	res := s.redis.Get(context.Background(), id).String()
	if clear {
		s.Clear(id)
	}
	return res
}

func (s *custom) Clear(id string) {
	s.redis.Del(context.Background(), id)
}

func (s *custom) Verify(id string, answer string, clear bool) bool {
	res, err := s.redis.Get(context.Background(), id).Result()
	if err != nil {
		return false
	}
	if clear {
		s.Clear(id)
	}
	return res == answer
}

func getStoreImageId(ctx *gin.Context, id string) string {
	id += tools.ClientIP(ctx)
	return "image_" + tools.Md5(id)
}

func getStoreEmailId(ctx *gin.Context, id string) string {
	id += tools.ClientIP(ctx)
	return "email_" + tools.Md5(id)
}

// VerifyImage 验证图像
func VerifyImage(ctx *gin.Context, id, answer string) bool {
	store := &custom{
		redis: global.Redis,
	}

	sid := getStoreImageId(ctx, id)

	if store.Verify(sid, answer, false) {
		store.Clear(sid)
		return true
	}
	return false
}

// VerifyEmail 验证邮箱
func VerifyEmail(ctx *gin.Context, id, answer string) bool {
	store := &custom{
		redis: global.Redis,
	}

	sid := getStoreEmailId(ctx, id)

	if store.Verify(sid, answer, false) {
		store.Clear(sid)
		return true
	}
	return false
}

// NewImageCaptcha 根据用户ip创建图形验证码
func NewImageCaptcha(ctx *gin.Context, length int, duration time.Duration) (string, string, error) {
	id := uuid.New().String()
	answer := tools.RandomCode(length)
	store := &custom{
		redis:    global.Redis,
		duration: duration,
	}

	// 生成二维码
	dt := captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	cpt := captcha.NewCaptcha(dt, store)
	item, err := cpt.Driver.DrawCaptcha(answer)
	if err != nil {
		return "", "", errors.CreateCaptchaError
	}

	return id, item.EncodeB64string(), store.Set(getStoreImageId(ctx, id), answer)
}

// NewEmailCaptcha 根据用户邮箱创建验证码
func NewEmailCaptcha(ctx *gin.Context, em string, length int, duration time.Duration) (string, error) {
	id := uuid.New().String()
	answer := tools.RandomCode(length)
	store := &custom{
		redis:    global.Redis,
		duration: duration,
	}

	// 发送二维码
	str := "您的邮箱验证码为：%s，该验证码%v分钟内有效，为了保证您的账户安全，请勿向他人泄露验证码吗信息"
	content := fmt.Sprintf(str, answer, int(duration/time.Minute))

	if email.New().Send(em, "验证码发送通知", content) != nil {
		return "", errors.CreateCaptchaError
	}

	return id, store.Set(getStoreEmailId(ctx, id), answer)
}
