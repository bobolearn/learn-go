package moretypes

import "fmt"

func OutPut7() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
}
