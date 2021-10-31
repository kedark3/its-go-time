package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Sum up to N
//
//  1. Get two numbers from the command-line: min and max
//  2. Convert them to integers (using Atoi)
//  3. By using a loop, sum the numbers between min and max
//
// RESTRICTIONS
//  1. Be sure to handle the errors. So, if a user doesn't
//     pass enough arguments or she passes non-numerics,
//     then warn the user and exit from the program.
//
//  2. Also, check that, min < max.
//
// HINT
//  For converting the numbers, you can use `strconv.Atoi`.
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
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
		sum += i
		fmt.Printf("%2d", i)
		if i != max {
			fmt.Print(" + ")
		} else if i == max {
			fmt.Printf(" = %2d\n", sum)
		}
	}
}
