package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "fmaulll:pass1234@tcp(todo-db.cn6dwyksqo66.ap-northeast-1.rds.amazonaws.com:3306)/todo"
	database, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Todo{})

	DB = database
}
