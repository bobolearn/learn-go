package main

import (
	"fmt"
	lesson "learn-go/pkg/lesson"

	quote "rsc.io/quote"
)

func main() {
	fmt.Println("Starting Learn Go!")
	lesson.TestInfo()
	fmt.Println(quote.Opt())
}
