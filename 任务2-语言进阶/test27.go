package main

import (
	"fmt"
	"time"
)

func send(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Printf("发送: %d\n", i)
		time.Sleep(100 * time.Millisecond) // 模拟处理时间
	}
	close(ch)
}

func receive(ch <-chan int) {

	for v := range ch {
		fmt.Printf("接收: %d\n", v)
		time.Sleep(150 * time.Millisecond) // 模拟处理时间
	}
}

// 编写一个程序，使用通道实现两个协程之间的通信。
// 一个协程生成从1到10的整数，并将这些整数发送到通道中，
// 另一个协程从通道中接收这些整数并打印出来。
func main() {

	ch := make(chan int)
	go send(ch)
	go receive(ch)

	// 等待足够时间让协程完成工作
	time.Sleep(2 * time.Second)
}
