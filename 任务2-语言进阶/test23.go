package main

import (
	"fmt"
	"time"
)

// 打印从1到10的奇数
func printOddNumber() {
	for i := 1; i < 10; i++ {
		if i%2 == 1 {
			fmt.Printf("打印从1到10的奇数: %v\n", i)
		}
	}
}

// 打印从2到10的偶数
func printEvenNumber() {
	for i := 2; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("打印从2到10的偶数: %v\n", i)
		}
	}
}

// 编写一个程序，使用 go 关键字启动两个协程，
// 一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func main() {

	go printOddNumber()
	go printEvenNumber()
	// 等待足够时间让协程执行完成
	time.Sleep(1 * time.Second)
}
