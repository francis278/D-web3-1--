package main

import (
	homework02 "lesson/homework02"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(dst ...interface{}) *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:Fanzhf123!@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(dst...)

	return db
}

func main() {
	db, err := gorm.Open(mysql.Open("root:Fanzhf123!@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	homework02.Run(db)

}
