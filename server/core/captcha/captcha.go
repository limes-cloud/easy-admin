package captcha

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/limeschool/easy-admin/server/core/email"
	"github.com/limeschool/easy-admin/server/core/metadata"
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
	res, _ := s.redis.Get(context.Background(), id).Result()
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

func getUserUid(userId int64) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("email_user_%v", userId))))
}

// NewEmailCaptcha 根据用户邮箱创建验证码
func NewEmailCaptcha(ctx *gin.Context, em string, length int, duration time.Duration) (string, error) {
	// 邮箱验证码需要限制次数
	md, err := metadata.GetFormContext(ctx)
	if err != nil {
		return "", err
	}

	store := &custom{
		redis:    global.Redis,
		duration: duration,
	}

	// 验证是否存在没有过期的验证码
	uid := getUserUid(md.UserID)
	if id := store.Get(uid, false); id != "" {
		return "", errors.DulSendCaptchaError
	}

	id := uuid.New().String()
	answer := tools.RandomCode(length)

	// 发送验证码
	str := "您的邮箱验证码为：%s，该验证码%d分钟内有效，为了保证您的账户安全，请勿向他人泄露验证码吗信息"
	content := fmt.Sprintf(str, answer, int(duration/time.Minute))
	if email.New().Send(em, "验证码发送通知", content) != nil {
		return "", errors.CaptchaSendError
	}

	// 进行加锁以及存储
	_ = store.Set(uid, "1")
	return id, store.Set(getStoreEmailId(ctx, id), answer)
}
