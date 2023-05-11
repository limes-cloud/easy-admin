package install

import (
	"github.com/limeschool/easy-admin/server/core"
	"github.com/limeschool/easy-admin/server/internal/system/model"
	"gorm.io/gorm"
)

type Model interface {
	TableName() string
	InitData(*gorm.DB) error
}

type install struct {
	db *gorm.DB
	ms []Model
}

func Init() {
	// 判断是否安装
	ins := install{
		ms: []Model{
			&model.Menu{},
			&model.Role{},
			&model.RoleMenu{},
			&model.Team{},
			&model.User{},
			&model.Notice{},
			&model.NoticeUser{},
			&model.LoginLog{},
		},
		db: core.GlobalOrm().GetDB(model.DBName()),
	}

	// 判断是否已经安装
	if ins.isInstall() {
		return
	}

	// 进行安装
	if err := ins.Install(); err != nil {
		panic("系统初始化失败:" + err.Error())
	}
}

func (ins *install) Install() error {
	for _, tb := range ins.ms {
		if err := ins.db.Migrator().AutoMigrate(tb); err != nil {
			return err
		}
		if err := tb.InitData(ins.db); err != nil {
			return err
		}
	}
	return nil
}

func (ins *install) isInstall() bool {
	is := true
	for _, tb := range ins.ms {
		is = is && ins.db.Migrator().HasTable(tb)
	}
	return is
}
