package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB(url, databaseName, userName, password string, maxConn, maxOpen int) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, password, url, databaseName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(maxConn)
	sqlDB.SetMaxOpenConns(maxOpen)
	return nil
}
