package basic_types

import "fmt"

func BasicTypes() {
	fmt.Println("These are the Basic Types Go!")
	//bool
	//string
	//int int8(byte) int16 int32 int64
	//float64 float32
	//uint uint8 uint16 uint32(rune) uint64
	//	complex32 complex64

	//	ZERO VALUES
	// variables declared without their explicit initial value are given their zero values
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	//	type conversions
	//the expression T(v) converts the value to the type T
	var myInt int = 6
	float1 := float64(myInt)
	fmt.Println(myInt, float1)

	//	type inference
	price := 780
	fmt.Println(price)
}
