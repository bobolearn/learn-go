package flowcontrol

import (
	"fmt"
	"runtime"
	"time"
)

func OutPut3() {
	fmt.Println("OutPut3")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OSX")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("windows")
	default:
		fmt.Println("unknown")
	}
}

func OutPut4() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("two far away")
	}
}
func OutPut5() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Printf("Good afternoon\n")
	default:
		fmt.Println("Good evening")
	}
}
