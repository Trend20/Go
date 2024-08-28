package control_structures

import "fmt"

func GoSwitch() {
	name := []string{"John", "Paul", "George", "Ringo"}
	switch len(name) {
	case 4:
		fmt.Println("Hey 4")
	case 5:
		fmt.Println("Hey 5")
	case 6:
		fmt.Println("Hey 6")
	default:
		fmt.Println("Hey all")
	}
}
