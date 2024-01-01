package main

import (
	"fmt"
	"math"
)

func main() {
	// add integer values
	i1, i2, i3 := 46, 70, 80
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum:", intSum)

	// add floating point value
	f1, f2, f3 := 56.8, 43.9, 58.2
	floatSum := f1 + f2 + f3
	fmt.Println("This is the float sum:", floatSum)

	// rounding off the float no to the nearest decimal place
	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("The value is now:", floatSum)
}
