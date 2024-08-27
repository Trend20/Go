// Go has various value types including STRINGS, INTEGERS, FLOATS, BOOLEANS etc.
// strings can be concatenated using the "+" operator e.g: "Hello" + "World" = Hello World

package main

import "fmt"

func main() {
	fmt.Println("Values are listed below.")
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("go" + "lang")
	fmt.Println("7.0 / 3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
