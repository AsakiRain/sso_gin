package db

import (
	"log"
	"sso_gin/config"
	"sso_gin/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	MYSQL *gorm.DB
)

func SetupMYSQL() {
	dblink := config.Dbuser + ":" + config.Dbpass + "@tcp(" + config.Dbhost + ":" + config.Dbport + ")/" + config.Dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dblink), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	err = db.AutoMigrate(&model.User{}, &model.RegFlow{}, &model.Minecraft{})
	if err != nil {
		log.Fatalf("迁移模型失败：%v", err)
	}
	MYSQL = db
}
