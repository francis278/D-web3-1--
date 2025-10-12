package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	fmt.Printf("数组: %v, 目标: %d\n", nums, target)
	fmt.Printf("结果: %v\n", result)
	fmt.Printf("说明: nums[%d] + nums[%d] = %d + %d = %d\n",
		result[0], result[1], nums[result[0]], nums[result[1]], target)
}
