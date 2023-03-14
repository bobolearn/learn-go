package moretypes

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func OutPut9() {
	// 函数值
	h := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(h(3, 4))

	fmt.Println(compute(h))
	fmt.Println(compute(math.Pow))
}
