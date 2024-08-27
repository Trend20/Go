package main

import (
	"fmt"
	"time"
)

func main() {
	course := "education"
	switch course {
	case "computer science":
		fmt.Println("This is the most interesting course!")
	case "medicine":
		fmt.Println("This is the most complicated course!")
	case "engineering":
		fmt.Println("This is the most trending course!")
	case "Human Resource":
		fmt.Println("This course is most loved by ladies!")
	default:
		fmt.Println("This is the simplest course!")
	}

	// more than one check
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is a weekend!")
	default:
		fmt.Println("It is weekday!")
	}

	// switching pre defined variable
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// switching types
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
