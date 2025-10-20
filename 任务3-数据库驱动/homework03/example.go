package homework03

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

// Sqlx入门
// 使用SQL扩展库进行查询

// 假设你已经使用Sqlx连接到一个数据库，
// 并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。

// 要求 ：
// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，
// 并将结果映射到一个自定义的 Employee 结构体切片中。
// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，
// 并将结果映射到一个 Employee 结构体中。

// employees 的表
type Employee struct {
	gorm.Model
	ID         int    `db:"id"` // 确保这里的 db tag 与数据库列名完全一致
	Name       string `db:"name"`
	Department string `db:"department"`
	// 添加所有数据库中的字段
	CreatedAt string `db:"created_at"` // 如果有时间字段
	UpdatedAt string `db:"updated_at"`
	Salary    int    `db:"salary"`
}

// 创建测试数据
// func createTestData(db *gorm.DB) error {

// 	// 创建测试账户
// 	employees := []Employee{
// 		{Name: "fan", Department: "IT部", Salary: 50000},
// 		{Name: "zhang", Department: "后勤部", Salary: 10000},
// 		{Name: "li", Department: "运输部", Salary: 20000},
// 	}

// 	if err := db.Create(&employees).Error; err != nil {
// 		return err
// 	}

// 	fmt.Println(" 测试数据创建完成")
// 	return nil
// }

// 使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，
func GetTechEmployees(db *sqlx.DB) ([]Employee, error) {

	var employees []Employee
	err := db.Select(&employees, "SELECT id,name,department,created_at,updated_at,salary FROM employees WHERE department=?", "技术部")
	if err != nil {
		return nil, fmt.Errorf("查询技术部员工失败: %w", err)
	}
	return employees, nil
}

// 使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，
func GetHighestPaidEmployees(db *sqlx.DB) ([]Employee, error) {

	var employees []Employee
	err := db.Select(&employees, "SELECT id,name,department,created_at,updated_at,salary FROM employees order by salary desc")
	if err != nil {
		return nil, fmt.Errorf("表中工资最高的员工信息没有: %w", err)
	}
	return employees, nil
}

func Run(db *sqlx.DB) error {
	//db.AutoMigrate(&Employee{})
	//createTestData(db)

	employees, err := GetTechEmployees(db)
	if err != nil {
		return err
	}
	for _, emp := range employees {
		fmt.Printf("技术部员工是ID: %d, 姓名: %s, 部门: %s, 薪资: %d\n",
			emp.ID, emp.Name, emp.Department, emp.Salary)
	}

	employees1, err := GetHighestPaidEmployees(db)
	if err != nil {
		return err
	}

	fmt.Printf("工资最高的员工信息是：ID: %d, 姓名: %s, 部门: %s, 薪资: %d\n",
		employees1[0].ID, employees1[0].Name, employees1[0].Department, employees1[0].Salary)

	return nil
}
