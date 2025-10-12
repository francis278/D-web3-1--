package main

import (
	"fmt"
)

// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func main() {
	testCases := [][]int{
		{1, 2, 3},    // 123 → 124 → [1,2,4]
		{9, 9, 9},    // 999 → 1000 → [1,0,0,0]
		{0},          // 0 → 1 → [1]
		{4, 3, 2, 1}, // 4321 → 4322 → [4,3,2,2]
	}

	for _, v := range testCases {
		result := plusOne(v)
		fmt.Println(result)
	}
}

func plusOne(num []int) []int {

	// 数组转数字
	number := arrayToNumber(num)

	// 数字加1
	number++

	// 数字转数组
	arr := numberToDigits(number)

	return arr
}

func arrayToNumber(num []int) int {
	if len(num) == 0 {
		return 0
	}

	number := 0

	for _, v := range num {
		number = number*10 + v
	}
	return number
}

func numberToDigits(num int) []int {
	if num == 0 {
		return []int{0}
	}

	var numbers []int

	for num > 0 {
		number := num % 10
		numbers = append([]int{number}, numbers...)
		num = num / 10
	}

	return numbers
}
