package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Break Up
//
//  1. Extend the "Only Evens" exercise
//  2. This time, use an infinite loop.
//  3. Break the loop when it reaches to the `max`.
//
// RESTRICTIONS
//  You should use the `break` statement once.
//
// HINT
//  Do not forget incrementing the `i` before the `continue`
//  statement and at the end of the loop.
//
// EXPECTED OUTPUT
//    go run main.go 1 10
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
	i, sum := min, 0
	for {
		if i%2 == 0 {
			sum += i
			fmt.Printf("%2d", i)
		}
		if i != max && i%2 == 0 {
			fmt.Print(" + ")
		} else if i == max {
			fmt.Printf(" = %2d\n", sum)
		}
		if i > max {
			break
		} else {
			i++
		}

	}
}
