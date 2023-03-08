package router

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	dsn := "root@tcp(mysql:3306)/sample-mysql-db?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
