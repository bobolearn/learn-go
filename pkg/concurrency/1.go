package concurrency

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// Go 线程
func OutPut() {
	go say("OutPut")
	say("hello")
}
