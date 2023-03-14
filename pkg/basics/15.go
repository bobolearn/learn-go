package basics

import (
	"fmt"
)

const PI = 3.1415
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(n int) int {
	return n*10 + 1
}
func needFloat(n float64) float64 {
	return n * 0.1
}
func OutPut15() {

	fmt.Println(PI)
	fmt.Println(needFloat(PI))
	fmt.Println(needFloat(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needInt(Small))
}
