package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"github.com/limeschool/easy-admin/server/global"
)

func Init() {
	conf := global.Config.Middleware.Casbin
	m := model.NewModel()
	m.AddDef("r", "r", "sub, obj, act")
	m.AddDef("p", "p", "sub, obj, act")
	m.AddDef("g", "g", "_, _")
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	m.AddDef("m", "m", "g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act")

	a, err := adapter.NewAdapterByDB(global.Orm[conf.DB])
	if err != nil {
		panic(err)
	}
	object, _ := casbin.NewEnforcer(m, a)
	if err = object.LoadPolicy(); err != nil {
		panic("casbin 初始化失败" + err.Error())
	}
	global.Casbin = object
}

func IsWhiteList(in string) bool {
	conf := global.Config.Middleware.Casbin
	return conf.Whitelist[in]
}

func IsBaseApi(ctx *gin.Context, method, path string) bool {
	//menu := model.Menu{}
	//return menu.GetBaseApiPath(ctx)[method+":"+path]
	return true
}
