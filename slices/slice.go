package main

import "fmt"

func main() {
	//simple empty slice
	var x []float64
	fmt.Println(x)
	//use make to create a slice
	y := make([]float64, 5)
	fmt.Println(y)
}
