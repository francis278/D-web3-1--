package main

import (
	"fmt"
	"sync"
	"time"
)

type Task func()

func scheduler(tasks []Task) {
	var wg sync.WaitGroup
	start := time.Now()
	for i, task := range tasks {
		wg.Add(1)
		go func(id int, t Task) {
			defer wg.Done()
			taskStart := time.Now()
			t()
			cost := time.Since(taskStart)
			fmt.Printf("任务%d 执行时间: %v\n", id, cost)
		}(i, task)
	}

	wg.Wait()
	fmt.Printf("总执行时间: %v\n", time.Since(start))
}

// 设计一个任务调度器，接收一组任务（可以用函数表示），
// 并使用协程并发执行这些任务，同时统计每个任务的执行时间。
func main() {

	tasks := []Task{
		func() { time.Sleep(1 * time.Second) },
		func() { time.Sleep(2 * time.Second) },
		func() { time.Sleep(500 * time.Millisecond) },
	}

	scheduler(tasks)
}
