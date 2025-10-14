package main

import (
	"fmt"
	"time"
)

func send(ch chan<- int) {
	for i := 0; i < 100; i++ {
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

// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，
// 消费者协程从通道中接收这些整数并打印。
func main() {

	ch := make(chan int, 5)
	go send(ch)
	go receive(ch)

	// 等待足够时间让协程完成工作
	time.Sleep(25 * time.Second)
}
