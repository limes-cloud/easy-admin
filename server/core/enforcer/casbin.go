package enforcer

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/limeschool/easy-admin/server/config"
	"github.com/limeschool/easy-admin/server/core/orm"
)

type enforcer struct {
	enforcer *casbin.Enforcer
}

type Enforcer interface {
	// Instance
	//
	//	@Description: 获取casbin实例
	//	@receiver e
	//	@return *casbin.Enforcer
	Instance() *casbin.Enforcer
}

func New(conf *config.Enforcer, db orm.Orm) Enforcer {
	m := model.NewModel()
	m.AddDef("r", "r", "sub, obj, act")
	m.AddDef("p", "p", "sub, obj, act")
	m.AddDef("g", "g", "_, _")
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	m.AddDef("m", "m", "g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act")

	a, err := adapter.NewAdapterByDB(db.GetDB(conf.DB))
	if err != nil {
		panic(err)
	}
	object, _ := casbin.NewEnforcer(m, a)
	if err = object.LoadPolicy(); err != nil {
		panic("enforcer 初始化失败" + err.Error())
	}
	return &enforcer{
		enforcer: object,
	}
}

func (e *enforcer) Instance() *casbin.Enforcer {
	return e.enforcer
}
