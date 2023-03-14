package flowcontrol

import "fmt"

func OutPut1() {
	sum := 0
	jsum := 1
	for i := 0; i < 10; i++ {
		sum += i
	}
	for jsum < 1000 {
		jsum += jsum
	}
	fmt.Println(sum)
	fmt.Println(jsum)
}
