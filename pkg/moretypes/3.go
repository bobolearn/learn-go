package moretypes

import "fmt"

func OutPut3() {
	// 数组
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])

	arr := [6]int{2, 4, 5, 8, 0, 1}
	fmt.Println(arr)
}
