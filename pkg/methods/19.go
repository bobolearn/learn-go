package methods

import (
	"fmt"
	"time"
)

type myError struct {
	When time.Time
	What string
}

func (e *myError) Error() string {
	return fmt.Sprintf("%v: %s", e.When, e.What)
}
func run() error {
	return &myError{
		time.Now(),
		"something went wrong",
	}
}

func OutPut19() {
	fmt.Println(run())
}
