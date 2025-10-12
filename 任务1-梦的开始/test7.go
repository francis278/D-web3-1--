package main

import (
	"fmt"
	"sort"
)

// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
// 可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，
// 遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；
// 如果没有重叠，则将当前区间添加到切片中。

func merge(intervals [][]int) [][]int {
	// 步骤1：处理特殊情况
	if len(intervals) <= 1 {
		return intervals
	}

	// 步骤2：按区间起点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 步骤3：初始化结果数组
	result := [][]int{intervals[0]}

	// 步骤4：遍历合并区间
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		last := result[len(result)-1]

		// 检查是否重叠
		if current[0] <= last[1] {
			// 重叠，合并区间（取最大的end值）
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			// 不重叠，添加到结果
			result = append(result, current)
		}
	}

	return result
}

func main() {
	// 测试用例
	testCases := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
		{{1, 4}, {0, 4}},
		{{1, 4}, {2, 3}},
		{{1, 4}, {0, 2}, {3, 5}},
		{},
		{{1, 3}},
	}

	for i, intervals := range testCases {
		result := merge(intervals)
		fmt.Printf("测试%d: %v → %v\n", i+1, intervals, result)
	}
}
