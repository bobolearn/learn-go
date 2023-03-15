package concurrency

import "fmt"

// 带缓冲的管道
func OutPut3() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch) //1
	fmt.Println(<-ch) //2
}
