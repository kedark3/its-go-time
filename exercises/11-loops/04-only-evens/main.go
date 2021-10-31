package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Not enough arguments. Usage: go run main.go <min> <max> e.g. go run main.go 1 10")
		return
	}
	min, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("found following error while parsing numeric value for min: %v\n", err)
		return
	}
	max, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("found following error while parsing numeric value for max: %v\n", err)
		return
	}

	for i, sum := min, 0; i <= max; i++ {
		if i%2 == 0 {
			sum += i
			fmt.Printf("%2d", i)
		}
		if i != max && i%2 == 0 {
			fmt.Print(" + ")
		} else if i == max {
			fmt.Printf(" = %2d\n", sum)
		}
	}
}
