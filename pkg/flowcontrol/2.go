package flowcontrol

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func OutPut2() {
	//if 的简短语句
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 10))
}
