package moretypes

import "fmt"

type Vertex struct {
	X int
	Y int
}

func OutPut1() {
	// 结构体
	fmt.Println(Vertex{1, 2})
	v := Vertex{3, 4}
	pv := &v
	pv.X = 5
	fmt.Println(v.X, v.Y)
	// 结构体文法
	two := Vertex{X: 11, Y: 22}
	one := Vertex{X: 10}
	empty := Vertex{}
	p := &Vertex{44, 55}
	fmt.Println(two, one, empty, p)

}
