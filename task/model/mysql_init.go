package model

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Database(connString string) {
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Logger.LogMode(logger.Info)
	if gin.Mode() == "release" {
		db.Logger.LogMode(logger.Warn)
	}
	DB = db
}
