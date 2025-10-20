package homework04

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

// 实现类型安全映射
// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。

// 定义一个 Book 结构体，包含与 books 表对应的字段。
// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，
// 并将结果映射到 Book 结构体切片中，确保类型安全。

// books 的表
type Books struct {
	gorm.Model
	ID        int    `db:"id"` // 确保这里的 db tag 与数据库列名完全一致
	Title     string `db:"title"`
	Author    string `db:"author"`
	CreatedAt string `db:"created_at"` // 如果有时间字段
	UpdatedAt string `db:"updated_at"`
	Price     int    `db:"price"`
}

// 创建测试数据
// func createTestData(db *gorm.DB) error {

// 	// 创建测试账户
// 	books := []Books{
// 		{Title: "Go语言编程实战", Author: "张三", Price: 50},
// 		{Title: "深入理解MySQL", Author: "李四", Price: 60},
// 		{Title: "Web开发进阶", Author: "王五", Price: 45},
// 		{Title: "分布式系统设计", Author: "赵六", Price: 70},
// 		{Title: "云计算架构", Author: "钱七", Price: 35},
// 	}

// 	if err := db.Create(&books).Error; err != nil {
// 		return err
// 	}

// 	fmt.Println(" 测试数据创建完成")
// 	return nil
// }

// 例如查询价格大于 50 元的书籍
func GetTechBooks(db *sqlx.DB) ([]Books, error) {

	var books []Books
	err := db.Select(&books, "SELECT id,title,author,price FROM books WHERE price>50")
	if err != nil {
		return nil, fmt.Errorf("查询价格大于 50 元的书籍没有", err)
	}
	return books, nil
}

func Run(db *sqlx.DB) error {
	//db.AutoMigrate(&Books{})
	//createTestData(db)

	books, err := GetTechBooks(db)
	if err != nil {
		return err
	}
	for _, emp := range books {
		fmt.Printf("ID: %d, 书名: %s, 作者: %s, 售价: %d\n",
			emp.ID, emp.Title, emp.Author, emp.Price)
	}
	return nil
}
