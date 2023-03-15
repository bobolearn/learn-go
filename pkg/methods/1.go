package methods

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 方法就是一类带特殊的 接收者 参数的函数。
func OutPut() {
	v := Vertex{X: 3, Y: 4}
	fmt.Println(v.Abs())
}
