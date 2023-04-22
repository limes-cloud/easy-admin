package captcha

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/mojocn/base64Captcha"
)

type image struct {
	ctx  *gin.Context
	name string
	tp   string
	*captcha
}

type Image interface {
	New() (*res, error)
	Verify(id, answer string) error
}

// New
//
//	@Description:  生成验证码图片
//	@receiver i
//	@param ctx 用于获取用户唯一场景id，保证场景验证码数据唯一
//	@return *captchaRes 返回验证码的uuid,base64,过期时间
//	@return error
func (i *image) New() (*res, error) {
	if !i.isTemplate(i.name, i.tp) {
		return nil, errors.New(fmt.Sprintf("%s captcha not exist", i.tp))
	}

	// 获取指定模板的配置
	cp := i.getTemplate(i.name, i.tp)

	// 生成随机验证码
	answer := i.randomCode(cp.Length)

	// 生成二维码
	dt := base64Captcha.NewDriverDigit(cp.Height, cp.Width, cp.Length, cp.Skew, cp.DotCount)
	item, err := dt.DrawCaptcha(answer)
	if err != nil {
		return nil, err
	}

	cache := i.rs.GetRedis(cp.Cache)
	cid := i.clientUUID(i.ctx, i.name, i.tp)

	// 清除上一次生成的结果,防止同时造成大量生成请求占用内存
	if id, _ := cache.Get(i.ctx, cid).Result(); id != "" {
		cache.Del(i.ctx, id)
	}

	// 进行结果缓存
	uid := uuid.New().String()
	if err = cache.Set(i.ctx, uid, answer, cp.Expire).Err(); err != nil {
		return nil, err
	}

	// 将答案ID挂载到当前的请求ip上
	if err = cache.Set(i.ctx, cid, uid, cp.Expire).Err(); err != nil {
		return nil, err
	}

	// 返回生成结果
	return &res{
		ID:     uid,
		Base64: item.EncodeB64string(),
		Expire: int(cp.Expire.Seconds()),
	}, nil
}

// Verify
//
//	@Description:  验证验证码
//	@receiver i
//	@param ctx
//	@param id 验证码的uuid
//	@param answer 验证码内容
//	@return error 验证通过则返回nil 否则返回定义错误原因
func (i *image) Verify(id, answer string) error {
	if !i.isTemplate(i.name, i.tp) {
		return errors.New(fmt.Sprintf("%s captcha not exist", i.tp))
	}

	// 获取指定模板的配置
	cp := i.getTemplate(i.name, i.tp)

	cache := i.rs.GetRedis(cp.Cache)
	cid := i.clientUUID(i.ctx, i.name, i.tp)

	// 验证id是否存在
	sid, err := cache.Get(i.ctx, cid).Result()
	if err != nil {
		return err
	}
	if sid != id {
		return errors.New(fmt.Sprintf("captcha id %s  not exist", id))
	}

	// 获取id的answer
	ans, err := cache.Get(i.ctx, id).Result()
	if err != nil {
		return err
	}
	if ans != answer {
		return errors.New("verify fail")
	}

	// 验证通过清除缓存
	return cache.Del(i.ctx, cid, id).Err()
}
