package methods

import (
	"fmt"
	"math"
)

type Vertex2 struct {
	X, Y float64
}

func Abs2(v Vertex2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func OutPut2() {
	v := Vertex2{X: 3, Y: 4}
	fmt.Println(Abs2(v))
}
