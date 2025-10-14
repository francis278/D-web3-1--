package main

import "fmt"

func numMultiplyTwo(num *[]int) {

	for i := range *num {
		(*num)[i] = (*num)[i] * 2
	}
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
func main() {

	num := []int{4, 6, 10, 50}
	numMultiplyTwo(&num)
	fmt.Printf("结果: %v\n", num)

}
