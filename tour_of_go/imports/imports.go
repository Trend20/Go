package imports

//multiple imports are grouped in parenthesis
import (
	"fmt"
	"math"
	"math/rand/v2"
)

func Imports() {
	fmt.Println("This is how Go handle imports!")
	fmt.Println("This is a random number:", rand.Float64())
	fmt.Println("The square root of 8 is :", math.Sqrt(8))
}
