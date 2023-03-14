package basics

import (
	"fmt"
	"math/cmplx"
)

// 基本类型
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func OutPut11() {
	fmt.Printf("%T,%v", ToBe, ToBe)
	fmt.Printf("\n%T,%v", MaxInt, MaxInt)
	fmt.Printf("\n%T,%v", z, z)
}
