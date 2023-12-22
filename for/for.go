package main

import "fmt"

func main() {
	m := 5

	for m <= 5 {
		fmt.Println(m)
		m = m + 1
	}

	for k := 10; k <= 15; k++ {
		fmt.Println(k)
	}

	for {
		fmt.Println("This is how for loops are written in Golang!")
		break
	}

	for n := 0; n <= 20; n++ {
		if n%2 == 0 {
			fmt.Println("If condition inside a for loop.")
			continue
		}
		fmt.Println()
	}
}
