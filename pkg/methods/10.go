package methods

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}
func OutPut10() {
	var i I = &T{S: "hello"}
	i.M()
	var f = F(math.Pi)
	f.M()

	// 报错
	// var j I
	// j.M()

	var j interface{}
	fmt.Println(j)

	j = 43
	fmt.Println(j)
}
