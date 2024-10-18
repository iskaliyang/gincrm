package initialize

import (
	"fmt"
	"gincrm/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func InitMySQL() {
	dns := global.Config.Mysql.DNS
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger:         newLogger,
		NamingStrategy: schema.NamingStrategy{SingularTable: false},
	})
	if err != nil {
		fmt.Println("数据库初始化失败", err)
		return
	}

	_db, err := db.DB()
	if err != nil {
		fmt.Println("数据库初始化失败")
		return
	}

	_db.SetMaxIdleConns(global.Config.Mysql.MinIdleConn)
	_db.SetMaxOpenConns(global.Config.Mysql.MinOpenConn)
	_db.SetConnMaxLifetime(time.Duration(global.Config.Mysql.ConnMaxLifetime))
	global.DB = db
}
