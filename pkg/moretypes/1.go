package moretypes

import "fmt"

// 指针
func OutPut() {
	i := 11
	j := 22
	pi := &i         // 指向 i
	fmt.Println(*pi) // i 的值
	*pi = 111        // 设置 i 的值
	fmt.Println(i)

	pj := &j
	*pj = *pj + 11
	fmt.Println(j)

}
