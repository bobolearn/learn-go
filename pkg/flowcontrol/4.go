package flowcontrol

import "fmt"

func OutPut6() {
	defer fmt.Println("OutPut6")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("this is OutPut6")
}
