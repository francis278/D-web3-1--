package main

import (
	"fmt"
)

// 定义一个 Shape 接口
type Shape interface {
	Area()
	Perimeter()
}

// 创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() {
	area := r.Width * r.Height
	fmt.Printf("Rectangle面积: %.2f\n", area)
}

func (r Rectangle) Perimeter() {
	perimeter := 2 * (r.Width + r.Height)
	fmt.Printf("Rectangle周长: %.2f\n", perimeter)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() {
	area := 3.14 * c.radius * c.radius
	fmt.Printf("Circle面积: %.2f\n", area)
}

func (c Circle) Perimeter() {
	perimeter := 2 * 3.14 * c.radius
	fmt.Printf("Circle周长: %.2f\n", perimeter)
}

// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
// 然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
func main() {
	rectangle := Rectangle{Width: 4, Height: 5}
	rectangle.Area()
	rectangle.Perimeter()

	circle := Circle{radius: 1}
	circle.Area()
	circle.Perimeter()
}
