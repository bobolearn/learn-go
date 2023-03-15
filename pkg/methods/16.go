package methods

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v,%v\n", v, v*2)
	case string:
		fmt.Printf("%v,%v\n", v, len(v))
	default:
		fmt.Printf("%T", v)
	}
}

func OutPut16() {
	do(12)
	do("test")
	do(true)
}
