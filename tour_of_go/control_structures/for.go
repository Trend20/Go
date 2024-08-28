package control_structures

import "fmt"

func ControlStructures() {
	//for loops
	names := []string{"John", "Paul", "George", "Ringo"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	//	using range
	for i, v := range names {
		fmt.Println("This is how to loop using the range keyword", i, v)
	}

	//	get the sum of all the numbers in a slice
	sum := 0
	myNums := []int{1, 2, 3, 4, 5, 80, 34, 78, 43}
	for _, n := range myNums {
		sum += n
	}
	fmt.Println(sum)

	//	if you omit the loop body it loops forever
	//for {
	//
	//}
}
