package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func counter(num *int64) {

	defer wg.Done()
	for i := 0; i < 1000; i++ {
		newValue := atomic.AddInt64(num, 1)
		fmt.Printf("计算: %d\n", newValue)
		time.Sleep(100 * time.Millisecond) // 模拟处理时间
	}
}

var wg sync.WaitGroup

// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func main() {

	var i int64 = 0

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

	finalValue := atomic.LoadInt64(&i)
	fmt.Printf("计数器的值 %d\n", finalValue)
}
