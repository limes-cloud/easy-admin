package core

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/config"
	"github.com/limeschool/easy-admin/server/consts"
	"github.com/limeschool/easy-admin/server/core/captcha"
	"github.com/limeschool/easy-admin/server/core/cert"
	"github.com/limeschool/easy-admin/server/core/email"
	"github.com/limeschool/easy-admin/server/core/enforcer"
	"github.com/limeschool/easy-admin/server/core/file"
	"github.com/limeschool/easy-admin/server/core/http"
	"github.com/limeschool/easy-admin/server/core/jwt"
	"github.com/limeschool/easy-admin/server/core/orm"
	"github.com/limeschool/easy-admin/server/core/redis"
	"github.com/limeschool/easy-admin/server/types"
	"go.uber.org/zap"
	"sync"
)

var ctxPool = sync.Pool{
	New: func() any {
		return &Context{}
	},
}

type Context struct {
	*gin.Context
}

func New(ctx *gin.Context) *Context {
	c := ctxPool.Get().(*Context)
	c.Context = ctx
	return c
}

// Gin 返回gin上下文
func (ctx *Context) Gin() *gin.Context {
	return ctx.Context
}

// Release 释放ctx到pool中
func (ctx *Context) Release() {
	ctx.Context = nil
	ctxPool.Put(ctx)
}

// Config 获取配置文件信息
func (ctx *Context) Config() *config.Config {
	return g.config
}

// TraceID 获取链路日志ID
func (ctx *Context) TraceID() string {
	return ctx.Gin().GetString(ctx.Config().Log.Field)
}

// SetTraceID 设置链路日志ID
func (ctx *Context) SetTraceID(id string) {
	ctx.Gin().Set(ctx.Config().Log.Field, id)
}

// Logger 获取文件日志器
func (ctx *Context) Logger() *zap.Logger {
	return g.logger.WithID(ctx.TraceID())
}

// Orm 获取数据库
func (ctx *Context) Orm() orm.Orm {
	return g.orm
}

// Redis 获取redis
func (ctx *Context) Redis() redis.Redis {
	return g.redis
}

// Enforcer 获取enforcer权限验证
func (ctx *Context) Enforcer() enforcer.Enforcer {
	return g.enforcer
}

// Cert 获取证书实例
func (ctx *Context) Cert() cert.Cert {
	return g.cert
}

// Email 邮件发送器
func (ctx *Context) Email() email.Email {
	return g.email
}

// Http 获取请求器
func (ctx *Context) Http() http.Request {
	return http.New(ctx.Config().Http, ctx.Logger())
}

func (ctx *Context) Metadata() *types.Metadata {
	val, is := ctx.Get(consts.Metadata)
	if !is {
		return nil
	}
	meta, is := val.(*types.Metadata)
	if !is {
		return nil
	}
	return meta
}

func (ctx *Context) SetMetadata(val *types.Metadata) {
	ctx.Set(consts.Metadata, val)
}

func (ctx *Context) File(subDir string) (file.File, error) {
	return file.NewFile(ctx.Config().File, subDir)
}

func (ctx *Context) SourceCtx() context.Context {
	c := context.Background()
	for key, val := range ctx.Context.Keys {
		c = context.WithValue(c, key, val)
	}
	return c
}

func (ctx *Context) ImageCaptcha(name string) captcha.Image {
	return g.captcha.Image(ctx.Gin(), name)
}

func (ctx *Context) EmailCaptcha(name string) captcha.Email {
	return g.captcha.Email(ctx.Gin(), name)
}

func (ctx *Context) ClientIP() string {
	ip := ctx.Context.ClientIP()
	if ip == "::1" {
		ip = ctx.GetHeader("X-Real-IP")
	}
	return ip
}

func (ctx *Context) Jwt() jwt.JWT {
	return jwt.New(ctx.Config().JWT, ctx.Redis(), ctx.Gin())
}
