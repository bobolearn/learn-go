package methods

import (
	"fmt"
	"math"
)

type Vertex5 struct {
	X, Y float64
}

func Abs4(v Vertex5) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func Scale2(v *Vertex5, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func OutPut5() {
	v := Vertex5{X: 3, Y: 4}
	Scale2(&v, 10)
	fmt.Println(Abs4(v))
}
