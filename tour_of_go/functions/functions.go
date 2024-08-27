package functions

import "fmt"

func Functions() {
	fmt.Println("This is how Functions work in Go!")
	studentsName := [6]string{"Mike", "Brian", "Bon", "Collins", "Brenda"}
	printNames(studentsName)
}

// function to print names
func printNames(names [6]string) string {
	for i, name := range names {
		fmt.Println(i, name)
	}
	return ""
}
