package arrays

import "fmt"

func GoArrays() {
	fmt.Println("Go Arrays")

	names := [5]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)

	marks := [5]int{10, 20, 30, 40, 50}
	fmt.Println(marks)
	marks[4] = 300
	fmt.Println(marks)
}
