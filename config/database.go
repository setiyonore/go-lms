package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database *gorm.DB
var DATABASE_URI string = "root:@tcp(localhost:3306)/lms?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() error {
	var err error

	Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
		// SkipDefaultTransaction: true,
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}
	return nil
}
