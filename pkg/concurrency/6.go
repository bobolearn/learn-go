package concurrency

import (
	"fmt"
	"sync"
	"time"
)

//  sync.Mutex 互斥锁

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Add(key string) {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Get(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func OutPut6() {
	c := SafeCounter{v: make(map[string]int, 0)}
	for i := 0; i < 1000; i++ {
		go c.Add("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Get("somekey"))
}
