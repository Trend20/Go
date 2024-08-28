package functions

import "fmt"

func Functions() {
	fmt.Println("This is how Functions work in Go!")
	studentsName := [6]string{"Mike", "Brian", "Bon", "Collins", "Brenda"}
	printNames(studentsName)
	fmt.Println(addInt(3, 4))
	fmt.Println(showNames("mary", "kevin"))
}

// function to print names
func printNames(names [6]string) string {
	for i, name := range names {
		fmt.Println(i, name)
	}
	return ""
}

// a function that adds two integer numbers
// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
func addInt(x, y int) int {
	return x + y
}

// function with multiple returns
func showNames(x, y string) (string, string) {
	return x, y
}

// named returns
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
