package methods

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs3() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func OutPut3() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs3())
}
