// constants declares variables that do not change.
// Go supports constants of strings, booleans and numerics.y
// Constant variables are declared using the "const" keyword.
// A const statement can appear anywhere a var statement can.
// Constant expressions perform arithmetic with arbitrary precision.
// A numeric constant has no type until it’s given one, such as by an explicit conversion.
// A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here math.Sin expects a float64.

package main

import (
	"fmt"
	"math"
)

const greeting string = "Hello Golang!"

func main() {
	fmt.Println(greeting)
	const title string = "This is how constant variables are declared"
	fmt.Println(title)

	const n = 50000
	fmt.Println(math.Sin(n))
}
