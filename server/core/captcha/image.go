package captcha

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/limeschool/easy-admin/server/errors"
	"github.com/mojocn/base64Captcha"
)

type image struct {
	ip      string
	name    string
	tp      string
	captcha *captcha
}

type Image interface {
	// New  发送邮箱验证码
	//
	//	@param ip 用户ip，防止同一个用户多次发送验证码
	//	@return res   验证码id，验证码图片base64编码，过期时间
	//	@return error 验证通过则返回nil 否则返回定义错误原因
	New() (*res, error)
	// Verify 验证邮箱验证码
	//
	//	@param ctx 用户获取用户唯一场景id
	//	@param id  验证码ID
	//	@param answer 验证码
	//	@return error 验证通过则返回nil 否则返回定义错误原因
	Verify(id, answer string) error
}

func (i *image) New() (*res, error) {
	// 获取指定模板的配置
	cp, is := i.captcha.getTemplate(i.name, i.tp)
	if !is {
		return nil, errors.New(fmt.Sprintf("%s captcha not exist", i.tp))
	}

	// 生成随机验证码
	answer := i.captcha.randomCode(cp.Length)

	// 生成验证码对应图片的base64
	dt := base64Captcha.NewDriverDigit(cp.Height, cp.Width, cp.Length, cp.Skew, cp.DotCount)
	item, err := dt.DrawCaptcha(answer)
	if err != nil {
		return nil, err
	}

	// 获取验证码存储器
	cache := i.captcha.cache.GetRedis(cp.Cache)
	// 获取当前用户的场景唯一id
	cid := i.captcha.cid(i.ip, i.name, i.tp)

	// 清除上一次生成的结果,防止同时造成大量生成请求占用内存
	if id, _ := cache.Get(context.Background(), cid).Result(); id != "" {
		cache.Del(context.Background(), id)
	}

	// 获取当前验证码验证码唯一id
	uid := uuid.New().String()
	if err = cache.Set(context.Background(), uid, answer, cp.Expire).Err(); err != nil {
		return nil, err
	}

	// 将本次验证码挂载到当前的场景id上
	if err = cache.Set(context.Background(), cid, uid, cp.Expire).Err(); err != nil {
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
	// 获取指定模板的配置
	cp, is := i.captcha.getTemplate(i.name, i.tp)
	if !is {
		return errors.New(fmt.Sprintf("%s captcha not exist", i.tp))
	}

	// 获取验证码存储器
	cache := i.captcha.cache.GetRedis(cp.Cache)

	// 获取当前用户的场景唯一id
	cid := i.captcha.cid(i.ip, i.name, i.tp)

	// 获取用户当前的验证码场景id
	sid, err := cache.Get(context.Background(), cid).Result()
	if err != nil {
		return err
	}
	// 对比用户当前的验证码场景是否一致
	if sid != id {
		return errors.New(fmt.Sprintf("captcha id %s  not exist", id))
	}

	// 获取指定验证码id的答案
	ans, err := cache.Get(context.Background(), id).Result()
	if err != nil {
		return err
	}
	// 对比答案是否一致
	if ans != answer {
		return errors.New("verify fail")
	}

	// 验证通过清除缓存
	return cache.Del(context.Background(), cid, id).Err()
}
