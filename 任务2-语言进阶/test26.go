package main

import "fmt"

type Person struct {
	Name string
	Age  string
}

type Employee struct {
	person     Person
	EmployeeID string
}

func PrintInfo(employee Employee) {

	fmt.Printf("员工EmployeeID: %s\n", employee.EmployeeID)

	fmt.Printf("员工Name: %s\n", employee.person.Name)

	fmt.Printf("员工Age: %s\n", employee.person.Age)

}

// 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
// 再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
// 为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
func main() {
	person := Person{Name: "fan", Age: "41"}

	employee := Employee{person: person, EmployeeID: "95956531"}

	PrintInfo(employee)
}
