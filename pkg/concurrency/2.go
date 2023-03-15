package concurrency

import "fmt"

// 管道

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func OutPut2() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := make(chan int)
	go sum(s, c)
	go sum(s[:5], c)
	a, b := <-c, <-c
	fmt.Println(a, b, a+b)
}
