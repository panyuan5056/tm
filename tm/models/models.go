package models

import (
	"fmt"

	"tm/pkg/logging"
	"tm/pkg/setting"

	"gorm.io/driver/mysql"
	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var (
		err error
	)
	if setting.DBTYPE == "sqlite3" {
		//DB, err = gorm.Open(sqlite.Open(setting.DBNAME), &gorm.Config{})
	} else {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			setting.USER,
			setting.PASSWORD,
			setting.HOST,
			setting.PORT,
			setting.DBNAME)
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		logging.Info(err)
	}

	DB.AutoMigrate(&Auth{})
	//db.Create(&Auth{Username: "root", Password: "111111"})

}
