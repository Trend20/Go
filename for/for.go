// for is Go's only looping construct

package main

import "fmt"

func main() {
	// The most basic type, with a single condition.
	m := 5

	for m <= 5 {
		fmt.Println(m)
		m = m + 1
	}

	// A classic initial/condition/after for loop.

	for k := 10; k <= 15; k++ {
		fmt.Println(k)
	}

	// for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
	for {
		fmt.Println("This is how for loops are written in Golang!")
		break
	}

	// You can also continue to the next iteration of the loop.
	for n := 0; n <= 20; n++ {
		if n%2 == 0 {
			fmt.Println("If condition inside a for loop.")
			continue
		}
		fmt.Println()
	}

	// print the first ten numbers in go
	i := 1
	for i <= 10 {
		fmt.Println("This is a display of the first ten numbers")
		i = i + 1
		fmt.Println(i)
	}

	for i := 0; i < 20; i++ {
		fmt.Println(i)
		g := i + 1
		fmt.Println(g)
	}
}
