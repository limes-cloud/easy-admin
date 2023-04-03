package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v8"
	"github.com/limeschool/easy-admin/server/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	Logger *zap.Logger
	Orm    map[string]*gorm.DB
	Redis  *redis.Client
	Casbin *casbin.Enforcer
	Cert   map[string][]byte
)
