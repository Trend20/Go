package main

import (
	"fmt"
	"time"
)

func main() {
	// current date and time
	n := time.Now()
	fmt.Println("Course started at:", n)

	// create a custom date
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go launched at:", t)
	fmt.Println(t.Format(time.ANSIC))
}
