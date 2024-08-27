package main

import "fmt"

func main() {

	m := 20
	// if block
	if m > 10 {
		fmt.Println("You should work hard to reach the pass mark.")
	}

	// check the value(if else)
	if m <= 20 {
		fmt.Println("You failed and you should work harder!")
	} else {
		fmt.Println("Congratulations! You passed.")
	}

	// logical operators
	if m > 20 || m <= 20 {
		fmt.Println("This is the pass mark for this course.")
	}
}
