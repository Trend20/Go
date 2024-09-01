package maps

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func GoMaps() {
	fmt.Println("Go Maps!")
	//A map maps keys to values.
	//	The zero value of a map is nil. A nil map has no keys, nor can keys be added.
	//	The make function returns a map of the given type, initialized and ready for use.
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])

	//	MAP LITERALS
	//Map literals are like struct literals, but the keys are required.
	var m2 = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m2)

	//If the top-level type is just a type name, you can omit it from the elements of the literal.
	var m3 = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m3)

	//	MUTATING MAPS
	//Insert or update an element in map m:
	var m4 = make(map[string]int)
	//INSERT
	m4["Total"] = 45
	fmt.Println(m4)
	//	UPDATE
	m4["Total"] = 80
	fmt.Println(m4)
	//	DELETE
	delete(m4, "Total")
	fmt.Println(m4)
}
