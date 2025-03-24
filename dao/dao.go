package dao

import (
	"abc/config"
	_ "github.com/go-sql-driver/mysql" //驱动
	"github.com/jinzhu/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	Db, err = gorm.Open("mysql", config.MySQLDB)
	if err != nil {
		//日志
	}
	if Db.Error != nil {
		//日志
	}

	//Db.DB().SetConnMaxLifetime(time.Hour)
	//Db.DB().SetMaxIdleConns(10)
	//Db.DB().SetMaxOpenConns(100)
}
