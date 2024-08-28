package control_structures

import "fmt"

func IfCondition() {
	fmt.Println("If condition?")
	marks := []int{67, 9, 3, 4, 2, 6, 8}
	for _, mark := range marks {
		fmt.Println(mark)
		if mark < 10 {
			fmt.Println("This is below average")
		} else {
			fmt.Println("Congratulations, you passed your exams")
		}
	}
}
