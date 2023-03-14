package moretypes

import "fmt"

type Location struct {
	Lat, Long float64
}

var mymap map[string]Location

func OutPut8() {
	// 映射
	mymap = make(map[string]Location)
	mymap["ShenZhen"] = Location{37.7717, -122.4014}
	mymap["GuangZhou"] = Location{Lat: 30.7717, Long: -12.4014}

	fmt.Println(mymap["ShenZhen"])
	fmt.Println(mymap["GuangZhou"])

	mymap2 := map[string]Location{
		"ShenZhen":  {37.7717, -122.4014},
		"GuangZhou": {Lat: 30.7717, Long: -12.4014},
	}
	fmt.Println(mymap2["ShenZhen"])
	fmt.Println(mymap2["GuangZhou"])
	// 修改映射
	mymap2["ShenZhen"] = Location{11.7717, -11.4014}
	fmt.Println(mymap2["ShenZhen"])
	// 删除元素：
	delete(mymap2, "GuangZhou")
	fmt.Println(mymap2)
	// 通过双赋值检测某个键是否存在：
	v, ok := mymap2["GuangZhou"]
	fmt.Println("The value:", v, "Present?", ok)
}
