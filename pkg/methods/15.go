package methods

import "fmt"

// 类型断言

func OutPut15() {
	var i interface{} = "hello world"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// 报错  f := i.(float64)

}
