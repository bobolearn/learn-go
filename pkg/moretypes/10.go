package moretypes

import "fmt"

func address() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func OutPut10() {
	// 函数的闭包
	pos, neg := address(), address()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
