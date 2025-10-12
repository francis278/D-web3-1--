package main

import "fmt"

func isValid(s string) bool {

	stack := make([]rune, 0)
	pairs := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, char := range s {

		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if top != pairs[char] {
				return false
			}
			stack = stack[:len(stack)-1]

		}
	}
	return len(stack) == 0
}

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
func main() {
	// 添加更多测试用例
	testCases := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
	}

	fmt.Println("\n 更多测试结果:")
	for _, test := range testCases {
		fmt.Printf("isValid(\"%s\") = %t\n", test, isValid(test))
	}
}
