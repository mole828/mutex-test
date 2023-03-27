package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	count int
	mu    sync.Mutex
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	c1 := &Counter{}
	c2 := &Counter{}
	var wg sync.WaitGroup
	start := time.Now()

	// 使用互斥锁
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			c1.Inc()
			wg.Done()
		}()
	}

	wg.Wait()
	time_with := time.Since(start)
	fmt.Printf("With mutex: %d, took %s\n", c1.Get(), time_with)

	// 不使用互斥锁
	start = time.Now()
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			c2.count++
			wg.Done()
		}()
	}

	wg.Wait()
	time_without := time.Since(start)
	fmt.Printf("Without mutex: %d, took %s\n", c2.count, time_without)
	fmt.Printf("|x|x|%d,%d|%s,%s|", c1.Get(), c2.Get(), time_with, time_without)
}
