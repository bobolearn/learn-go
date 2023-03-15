package concurrency

import (
	"fmt"
	"time"
)

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case <-quit:
			fmt.Println("quit")
			return
		case c <- x:
			x, y = y, x+y
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}

	}
}

func OutPut5() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}
