package initialize

import (
	"errors"
	"fmt"
	"time"
	"turingapp/global"
	"turingapp/model"
	"turingapp/service"

	"gorm.io/gorm/logger"

	"github.com/techidea8/restgo"
	"github.com/techidea8/restgo/database"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

//初始化表格
func NamingStrategy(cfg *gorm.Config) {
	cfg.DisableForeignKeyConstraintWhenMigrating = true
	cfg.NamingStrategy = schema.NamingStrategy{
		TablePrefix:   global.Conf.Database.TablePrefix, // 表名前缀，`User` 的表名应该是 `t_users`
		SingularTable: true,                             // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
	}
}

func Newloger(level int) database.OrmOption {
	return func(cfg *gorm.Config) {
		lovel := logger.Silent
		if 2 == level {
			lovel = logger.Error
		}
		if 3 == level {
			lovel = logger.Warn
		}
		if 4 == level {
			lovel = logger.Info
		}
		cfg.Logger = New(logger.Config{LogLevel: lovel}, global.Log)
	}
}

//flag:true,禁用全局事务
//flag:false,启用全局事务
func SkipDefaultTransaction(flag bool) database.OrmOption {
	return func(cfg *gorm.Config) {
		cfg.SkipDefaultTransaction = flag
	}
}

//注册控制器
func InitDatabase() *gorm.DB {
	global.DbEngin = database.InitDataBase(global.Conf.Database,
		NamingStrategy,
		SkipDefaultTransaction(true),
		Newloger(global.Conf.Database.LogMode))
	err := global.DbEngin.AutoMigrate(
		new(model.User),
		new(model.Weixin),
		new(model.Dict),
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	return global.DbEngin
}

//初始化admin用户和密码
func InitAdmin(mobile, passwd string) uint {
	user := &model.User{}
	user.Mobile = mobile
	user.Passwd = service.BuildPasswd(passwd)
	user.AsAdmin()
	user.ExpireAt = restgo.DateTimeFromTime(time.Now().AddDate(100, 0, 0))
	if u, err := service.FindUserByMobile(user.Mobile); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.DbEngin.Create(user)
			return user.Id
		} else {
			return u.Id
		}
	} else {
		return u.Id
	}

}
