package main

/* This import format is called "Factored" import, you could write them on separate lines too as
import "math"
import "fmt"

*/
import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welome to the playground!")
	// Fun Fact: This returns 2009-11-10 in the Go Playground!!
	fmt.Println("The time is", time.Now())
}
