package enforcer

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/config"
	"gorm.io/gorm"
	"strings"
)

type enforcer struct {
	whitelist map[string]bool
	e         *casbin.Enforcer
}

type instance struct {
	ctx *gin.Context
}

type Enforcer interface {
	IsWhitelist(method, path string) bool
	Instance() *casbin.Enforcer
}

func New(conf config.Enforcer, db *gorm.DB) Enforcer {
	m := model.NewModel()
	m.AddDef("r", "r", "sub, obj, act")
	m.AddDef("p", "p", "sub, obj, act")
	m.AddDef("g", "g", "_, _")
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	m.AddDef("m", "m", "g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act")

	a, err := adapter.NewAdapterByDB(db)
	if err != nil {
		panic(err)
	}
	object, _ := casbin.NewEnforcer(m, a)
	if err = object.LoadPolicy(); err != nil {
		panic("enforcer 初始化失败" + err.Error())
	}
	return &enforcer{
		whitelist: conf.Whitelist,
		e:         object,
	}
}

// IsWhitelist
//
//	@Description: 判断接口资源是否为白名单
//	@receiver e
//	@param method 请求方法
//	@param path 请求路径
//	@return bool
func (e *enforcer) IsWhitelist(method, path string) bool {
	return e.whitelist[strings.ToLower(method+":"+path)]
}

// Instance
//
//	@Description: 获取casbin实例
//	@receiver e
//	@return *casbin.Enforcer
func (e *enforcer) Instance() *casbin.Enforcer {
	return e.e
}
