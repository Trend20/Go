package control_structures

import (
	"fmt"
	"time"
)

func GoSwitch() {
	name := []string{"John", "Paul", "George", "Ringo"}
	switch len(name) {
	case 4:
		fmt.Println("Hey 4")
	case 5:
		fmt.Println("Hey 5")
	case 6:
		fmt.Println("Hey 6")
	default:
		fmt.Println("Hey all")
	}
}

func SwitchDay() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// defer statement
func Defer() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

// Stacking defers
// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
// To learn more about defer statements read this blog post.
func StackedDefer() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
