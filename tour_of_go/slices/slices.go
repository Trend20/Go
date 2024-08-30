package slices

import "fmt"

func GoSlices() {
	fmt.Println("Go Slices!")
	names := []string{"John", "Paul", "George"}
	fmt.Println(names)
	fmt.Println(names[1:3])

	//	you can create a slice from an array
	marks := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	var s = marks[2:6]
	fmt.Println(s)

	//A slice does not store any data, it just describes a section of an underlying array.
	//	Changing the elements of a slice modifies the corresponding elements of its underlying array.
	//	Other slices that share the same underlying array will see those changes.
	s[2] = 45
	fmt.Println(s)

	//	slice literal
	//A slice literal is like an array literal without the length.
	ages := []int{23, 32, 12, 31, 53}
	fmt.Println(ages)

	//	example of type structs
	car := []struct {
		name    string
		price   float64
		color   string
		millage string
	}{
		{"toyota", 238.90, "red", "7800km"},
		{"mazda", 298.90, "white", "9000km"},
		{"honda", 238.90, "grey", "8000km"},
	}
	fmt.Println(car)

	//slice defaults
	//	when slicing, you can decide to omit the low and high bounds to  use their defaults instead
	//	the default for the low bound is zero and the length of the slice for the high bound

	totals := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(totals)
	slice1 := totals[:6]
	fmt.Println(slice1)
	slice2 := totals[6:]
	fmt.Println(slice2)

	//	slice length and capacity

	//A slice has both a length and a capacity.
	//	The length of a slice is the number of elements it contains.
	//	The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	totalItems := len(nums)
	fmt.Println(totalItems)
	sliceCaps := cap(nums)
	fmt.Println(sliceCaps)

	//	Nil slices
	//The zero value of a slice is nil.
	//	A nil slice has a length and capacity of 0 and has no underlying array.
	var items []int
	fmt.Println(items)
	//(can we define a nil slice using the short variable declaration method?)
	//ANS: No, you cannot define a nil slice using the short variable declaration method (:=). When you use the short variable declaration,
	//Go automatically allocates memory for the slice, resulting in an empty slice rather than a nil slice.

	//	EMPTY SLICE: s:=[]string{}
	//	NIL SLICE: var s []string

	//	creating a slice with make
	//Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
	//	The make function allocates a zeroed array and returns a slice that refers to that array:
	s1 := make([]int, 6)
	fmt.Println(s1)

	//	slices of slices
	//Slices can contain any type, including other slices.
	animals := [][]string{
		[]string{"mbuzi", "ng'ombe"},
		[]string{"cow", "choma"},
		[]string{"goat", "fry"},
	}
	fmt.Println(animals)

	//	Appending to a slice
	cooks := []string{"boo", "foo", "buzz", "bond", "doe"}
	fmt.Println(cooks)
	new1 := append(cooks, "mool")
	fmt.Println(new1)

	//	range on a slice
	//The range form of the for loop iterates over a slice or map.
	//	When ranging over a slice, two values are returned for each iteration.
	//	The first is the index, and the second is a copy of the element at that index.
	moonKey := []string{"moon", "moon", "moon", "moon", "moon", "moon"}
	fmt.Println(moonKey)
	//	loop through using range
	for _, v := range moonKey {
		fmt.Println(v)
	}
}
