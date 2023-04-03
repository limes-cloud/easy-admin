package orm

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"github.com/limeschool/easy-admin/server/global"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
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

func Init() {
	configs := global.Config.Orm
	global.Orm = make(map[string]*gorm.DB)

	for _, conf := range configs {
		if !conf.Enable {
			continue
		}

		// 连接主数据库
		db, err := gorm.Open(open(conf.Drive, conf.Dsn), &gorm.Config{
			Logger: newOrmLog(conf),
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

		global.Orm[conf.Name] = db
		sdb, _ := db.DB()
		sdb.SetConnMaxLifetime(conf.MaxLifetime)
		sdb.SetMaxOpenConns(conf.MaxOpenConn)
		sdb.SetMaxIdleConns(conf.MaxIdleConn)
	}

}
