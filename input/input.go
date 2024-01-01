package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("This is how to get user input from the console:")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text:")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered: ", input)
}
