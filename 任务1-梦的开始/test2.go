package main

import (
	"fmt"
)

func isPalindromeNumber(num int) bool {
	if num < 0 {
		return false
	}
	if num < 10 {
		return true
	}

	return compareArrays(numberToDigits(num), reverseArray(num))

}

// 数字转数组
func numberToDigits(x int) []int {
	if x == 0 {
		return []int{0}
	}
	var digits []int
	for x > 0 {
		digit := x % 10
		digits = append([]int{digit}, digits...) // 添加到前面，保持顺序
		x /= 10
	}
	return digits
}

// 反转数组
func reverseArray(num int) []int {
	if num == 0 {
		return []int{0}
	}
	var digits []int

	for num > 0 {
		digit := num % 10

		digits = append(digits, digit) // 直接追加到后面（自然就是反转的）

		num = num / 10
	}
	return digits
}

// 比较两个数组是否相等
func compareArrays(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// 判断一个整数是否是回文数
func main() {

	tests := []int{121, -121, 10, 12321, 0, 1221, 123, 1}
	fmt.Println("回文数判断结果:")

	for _, num := range tests {
		result := isPalindromeNumber(num)
		fmt.Printf("%d -> %t\n", num, result)
	}

}
