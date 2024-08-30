package structs

import "fmt"

type Vertex struct {
	x, y int
}

type Car struct {
	name    string
	model   string
	price   float64
	millage string
}

func GoStructs() {
	//a struct is a collection of fields
	fmt.Println("Go Structs!")
	fmt.Println(Vertex{3, 4})
	fmt.Println(Car{name: "wang", model: "wang"})

	//Struct fields are accessed using a dot.
	myCar := Car{name: "wing", model: "wing"}
	name := myCar.name
	fmt.Println(name)

	//	struct pointers
	m := Vertex{9, 7}
	//create a pointer to m
	p := &m
	fmt.Println(p)
	//access the struct value using the pointer
	fmt.Println(p.y)

	//struct literal
	type Person struct {
		name     string
		age      int
		location string
	}

	//	person instances
	person1 := Person{"Byron", 13, "Nairobi"}
	person2 := Person{"Cool", 37, "Nakuru"}
	person3 := Person{"More", 56, "Kisumu"}
	fmt.Println(person1, person2, person3)

	//	person pointer
	pointPerson1 := &person1
	pointPerson1.name = "Mbuvi"
	fmt.Println(pointPerson1)
}
