package variables

import "fmt"

func Variables() {
	fmt.Println("This is how variables are defined in Go!")

	//	using the var keyword
	var name string = "Foo"
	fmt.Println(name)

	//	using the shorthand
	age := 67
	fmt.Println(age)

	//	define multiple variables in a single line
	//var size, color, brand string = "large", "white", "marine"
	//fmt.Println(size, color, brand)

	size, color, brand := "large", "white", "marine"
	fmt.Println(size, color, brand)
}
