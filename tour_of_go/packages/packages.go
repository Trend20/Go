//Every Go program is made up of packages and the program start running in package main.

package packages

import (
	"fmt"
	"math/rand"
)

func Packages() {
	fmt.Println("These are Go Packages!")
	fmt.Println("My favourite number is", rand.Intn(10))
}
