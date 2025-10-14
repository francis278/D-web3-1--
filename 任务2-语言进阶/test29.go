package main

import (
	"fmt"
	"sync"
	"time"
)

func counter(num *int) {

	defer wg.Done()
	for i := 0; i < 100; i++ {
		mu.Lock()
		*num++
		current := *num
		mu.Unlock()
		fmt.Printf("计算: %d\n", current)
		time.Sleep(100 * time.Millisecond) // 模拟处理时间
	}
}

var mu sync.Mutex
var wg sync.WaitGroup

// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func main() {

	i := 0

	wg.Add(10)
	go counter(&i)
	go counter(&i)
	go counter(&i)
	go counter(&i)
	go counter(&i)
	go counter(&i)
	go counter(&i)
	go counter(&i)
	go counter(&i)
	go counter(&i)

	wg.Wait()

	mu.Lock()
	fmt.Printf("计数器的值 %d\n", i)
	mu.Unlock()
}
