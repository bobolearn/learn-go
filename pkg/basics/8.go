package basics

import "fmt"

// 变量声明
var c, python, javascrip bool

// 变量的初始化
var i, j = 1, 2

func OutPut8() {
	var one, two, three = true, false, "no"
	// 短变量声明
	k := 343.5
	fmt.Println(c, python, javascrip)
	fmt.Println(i, j, one, two, three)
	fmt.Println(k)
}
