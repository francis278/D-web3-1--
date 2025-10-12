package main

import (
	"fmt"
)

// 查找字符串数组中的最长公共前缀
func main() {
	testCases := [][]string{
		{"flower", "flow", "flight"},
		{"dog", "racecar", "car"},
		{"", "abc", "def"},
		{"abc", "abc", "abc"},
		{"a"},
		{},
	}

	for i, strs := range testCases {
		result := longestCommonPrefix(strs)
		fmt.Printf("测试%d: %v -> \"%s\"\n", i+1, strs, result)
	}
}

// 方法1：纵向扫描
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]

		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}
