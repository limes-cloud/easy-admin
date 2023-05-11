package captcha

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type email struct {
	name    string
	tp      string
	captcha *captcha
	ip      string
}

type Email interface {

	// New  发送邮箱验证码
	//
	//	@param ip 用户ip，防止同一个用户多次发送验证码
	//	@param email  发送邮箱
	//	@return res   验证码id,过期时间
	//	@return error 验证通过则返回nil 否则返回定义错误原因
	New(email string) (*res, error)
	// Verify 验证邮箱验证码
	//
	//	@param ctx 用户获取用户唯一场景id
	//	@param id  验证码ID
	//	@param answer 验证码
	//	@return error 验证通过则返回nil 否则返回定义错误原因
	Verify(id, answer string) error
}

func (e *email) New(email string) (*res, error) {
	// 获取指定模板的配置
	cp, is := e.captcha.getTemplate(e.name, e.tp)
	if !is {
		return nil, errors.New(fmt.Sprintf("%s captcha not exist", e.tp))
	}

	// 生成随机验证码
	answer := e.captcha.randomCode(cp.Length)

	// 获取验证码存储器
	cache := e.captcha.cache.GetRedis(cp.Cache)

	// 获取当前用户的场景唯一id
	cid := e.captcha.cid(e.ip, e.name, e.tp)

	// 清除上一次生成的结果,防止同时造成大量生成请求占用内存
	if id, _ := cache.Get(context.Background(), cid).Result(); id != "" {
		cache.Del(context.Background(), id)
	}

	// 获取当前验证码验证码唯一id
	uid := uuid.New().String()
	if err := cache.Set(context.Background(), uid, answer, cp.Expire).Err(); err != nil {
		return nil, err
	}

	// 将本次验证码挂载到当前的场景id上
	if err := cache.Set(context.Background(), cid, uid, cp.Expire).Err(); err != nil {
		return nil, err
	}

	// 发送邮件
	if err := e.captcha.email.NewSender(cp.Template).Send(email, map[string]any{
		"answer": answer,
		"minute": int(cp.Expire.Minutes()),
	}); err != nil {
		return nil, err
	}

	// 返回生成结果
	return &res{
		ID:     uid,
		Expire: int(cp.Expire.Seconds()),
	}, nil
}

func (e *email) Verify(id, answer string) error {
	// 获取指定模板的配置
	cp, is := e.captcha.getTemplate(e.name, e.tp)
	if !is {
		return errors.New(fmt.Sprintf("%s captcha not exist", e.tp))
	}

	// 获取验证码存储器
	cache := e.captcha.cache.GetRedis(cp.Cache)

	// 获取当前用户的场景唯一id
	cid := e.captcha.cid(e.ip, e.name, e.tp)

	// 验证用户是否生成过验证码id
	sid, err := cache.Get(context.Background(), cid).Result()
	if err != nil {
		return err
	}

	// 对比用户当前的验证码场景是否一致
	if sid != id {
		return errors.New(fmt.Sprintf("captcha id %s not exist", id))
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
