package homework

import (
	"gorm.io/gorm"
)

// students 的表
type Students struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey;autoIncrement"` // 主键，自增
	Name  string // 学生姓名，字符串类型
	Age   int    // 学生年龄，整数类型
	Grade string // 学生年级，字符串类型
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&Students{})

	//求 ：编写SQL语句向 students 表中插入一条新记录，
	//学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	//db.Create(&Students{Name: "张三", Age: 20, Grade: "三年级"})

	// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	// var students []Students
	// db.Where("age > ?", 18).Find(&students)
	// for _, student := range students {
	// 	fmt.Printf("姓名: %s, 年龄: %d\n", student.Name, student.Age)
	// }

	// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	//var students []Students
	//db.Where("name = ?", "张三").Find(&students)
	//db.Model(&students).Where("name = ?", "张三").Update("grade", "四年级")
	//for _, student := range students {
	//	fmt.Printf("姓名: %s, 年龄: %d, 年级: %s\n", student.Name, student.Age, student.Grade)
	//}

	// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	db.Where("age > 15 and id = 1").Delete(&Students{})
	db.Delete(&Students{}, "age > 15")
}
