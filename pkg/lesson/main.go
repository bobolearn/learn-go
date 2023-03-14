package lesson

import (
	"fmt"
	g "learn-go/pkg/greetings"
	"log"
)

func TestInfo() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	//message, err := g.Hello("bobocode")
	names := []string{"lili", "bobo", "lala"}
	messages, err := g.Hellos(names)
	if err != nil {
		log.Fatal(err) // 控制台报错退出 exit status 1
	}
	fmt.Println(messages)
}
