package main

import (
	"lesson/homework04"

	_ "github.com/go-sql-driver/mysql" // 添加这一行！
	"github.com/jmoiron/sqlx"
)

func main() {
	// db1, err := gorm.Open(mysql.Open("root:Fanzhf123!@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	// if err != nil {
	// 	panic(err)
	// }

	dsn := "root:Fanzhf123!@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}

	//homework01.Run(db)
	//homework02.Run(db)
	//homework03.Run(db)
	homework04.Run(db)

}
