package main

import (
	"fmt"
	"github.com/Trend20/Go/tree/main/tour_of_go/basic_types"
	"github.com/Trend20/Go/tree/main/tour_of_go/control_structures"
	"github.com/Trend20/Go/tree/main/tour_of_go/functions"
	"github.com/Trend20/Go/tree/main/tour_of_go/imports"
	"github.com/Trend20/Go/tree/main/tour_of_go/packages"
	"github.com/Trend20/Go/tree/main/tour_of_go/variables"
)

func main() {
	fmt.Println("Welcome to the Tour of Go!")
	packages.Packages()
	imports.Imports()
	functions.Functions()
	variables.Variables()
	basic_types.BasicTypes()
	control_structures.ControlStructures()
	control_structures.IfCondition()
	control_structures.GoSwitch()
}
