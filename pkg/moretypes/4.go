package moretypes

import "fmt"

func OutPut4() {
	// 切片
	arr := [6]int{2, 4, 5, 8, 0, 1}
	var slice []int = arr[1:4]
	fmt.Println(slice)

	// 切片就像数组的引用
	names := []string{"John", "Paul", "George", "Ringo", "Michael", "James"}
	fmt.Println(names)
	a := names[:2]
	b := names[2:4]

	fmt.Println(a, b)
	b[0] = "修改"
	fmt.Println(a, b)
	fmt.Println(names)
	// 切片文法
	q := []int{22, 14, 25, 18, 0, 21}
	fmt.Println(q)
	r := []bool{true, false, false, true}
	fmt.Println(r)
	s := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{3, true},
		{4, false},
		{5, true},
	}
	fmt.Println(s)
	// nil 切片
	var nilSlice []int
	fmt.Println(nilSlice, len(nilSlice), cap(nilSlice))
	if nilSlice == nil {
		fmt.Println("nil")
	}
	// 切片的长度与容量
	ss := []int{2, 3, 5, 7, 11, 13}
	s1 := ss[:0]
	fmt.Println(s1)
	s2 := ss[0:4]
	fmt.Println(s2)
	s3 := ss[2:]
	fmt.Println(s3)
	fmt.Printf("len=%d cap=%d %v\n", len(ss), cap(ss), ss)
	// 用 make 创建切片
	m := make([]int, 5)
	fmt.Println(m)
	m1 := make([]int, 0, 10)
	fmt.Println(m1)
	fmt.Printf("len=%d cap=%d %v\n", len(m), cap(m), m)
}
