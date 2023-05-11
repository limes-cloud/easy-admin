package orm

import (
	"errors"
	"fmt"
	"github.com/glebarez/sqlite"
	"github.com/limeschool/easy-admin/server/config"
	"github.com/limeschool/easy-admin/server/core/logger"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"sync"
)

const (
	_mysql      = "mysql"
	_postgresql = "postgresql"
	_sqlite     = "sqlite"
	_sqlServer  = "sqlServer"
	_tidb       = "tidb"
)

func open(drive, dsn string) gorm.Dialector {
	switch drive {
	case _mysql, _tidb:
		return mysql.Open(dsn)
	case _postgresql:
		return postgres.Open(dsn)
	case _sqlite:
		return sqlite.Open(dsn)
	case _sqlServer:
		return sqlserver.Open(dsn)
	default:
		return nil
	}
}

type orm struct {
	mu sync.RWMutex
	db map[string]*gorm.DB
}

type Orm interface {
	// Get
	//
	//	@Description: 获取指定名称的orm实例，如果实例不存在则会返回报错
	//	@param name 实例名称
	//	@return *gorm.DB
	//	@return error
	Get(name string) (*gorm.DB, error)
	// GetDB
	//
	//	@Description: 获取指定名称的orm实例，如果实例不存在则返回nil
	//	@param name 实例名称
	//	@return *gorm.DB
	GetDB(name string) *gorm.DB
	// GormWhere
	//
	//	@Description: 根据入参结构自动拼接where字段
	//	@param db gorm.DB gorm.DB实例
	//	@param tb 表名
	//	@param val 查询的结构可以为 struct/map，struct会根据tag规则，map会采用等式
	//	@return *gorm.DB
	GormWhere(db *gorm.DB, tb string, val interface{}) *gorm.DB
}

// New
//
//	@Description: 初始化数据库
//	@param conf 数据库配置
//	@param logger 日志器
//	@return Orm 数据库实例
func New(cm []config.Orm, logger logger.Logger) Orm {
	ormIns := orm{
		db: make(map[string]*gorm.DB),
		mu: sync.RWMutex{},
	}

	ormIns.mu.Lock()
	defer ormIns.mu.Unlock()

	for _, conf := range cm {
		if !conf.Enable {
			continue
		}

		// 连接主数据库
		db, err := gorm.Open(open(conf.Drive, conf.Dsn), &gorm.Config{
			Logger: newOrmLog(conf, logger),
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "",
				SingularTable: true,
			},
		})
		if err != nil {
			panic(fmt.Errorf("主数据库%v连接失败：%v", conf.Name, err.Error()))
		}

		// 连接从数据库
		var replicas []gorm.Dialector
		for _, dsn := range conf.Replicas {
			replicas = append(replicas, open(conf.Drive, dsn))
		}
		if err = db.Use(dbresolver.Register(dbresolver.Config{
			Replicas: replicas,
			Policy:   dbresolver.RandomPolicy{},
		})); err != nil {
			panic(fmt.Errorf("从数据库连接失败：%v", err.Error()))
		}

		ormIns.db[conf.Name] = db
		sdb, _ := db.DB()
		sdb.SetConnMaxLifetime(conf.MaxLifetime)
		sdb.SetMaxOpenConns(conf.MaxOpenConn)
		sdb.SetMaxIdleConns(conf.MaxIdleConn)
	}
	return &ormIns
}

func (o *orm) Get(name string) (*gorm.DB, error) {
	o.mu.RLock()
	defer o.mu.RUnlock()

	if o.db[name] == nil {
		return nil, errors.New("not exist db")
	}
	return o.db[name], nil
}

func (o *orm) GetDB(name string) *gorm.DB {
	o.mu.RLock()
	defer o.mu.RUnlock()

	return o.db[name]
}
