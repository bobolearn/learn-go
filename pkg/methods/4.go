package methods

import (
	"fmt"
	"math"
)

// 指针接收者
type Vertex4 struct {
	X, Y float64
}

func (v Vertex4) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (v *Vertex4) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func OutPut4() {
	v := Vertex4{X: 3, Y: 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
