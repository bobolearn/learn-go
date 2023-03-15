package methods

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s is %d years old\n", p.Name, p.Age)
}

func OutPut17() {
	a := Person{"bobo", 28}
	z := Person{"lili", 28}
	fmt.Println(a, z)
}
