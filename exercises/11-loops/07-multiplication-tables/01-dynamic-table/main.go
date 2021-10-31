package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Dynamic Table
//
//  Get the size of the table from the command-line
//    Passing 5 should create a 5x5 table
//    Passing 10 for a 10x10 table
//
// RESTRICTION
//  Solve this exercise without looking at the original
//  multiplication table exercise.
//
// HINT
//  There was a max constant in the original program.
//  That determines the size of the table.
//
// EXPECTED OUTPUT
//
//  go run main.go
//    Give me the size of the table
//
//  go run main.go -5
//    Wrong size
//
//  go run main.go ABC
//    Wrong size
//
//  go run main.go 1
//    X    0    1
//    0    0    0
//    1    0    1
//
//  go run main.go 2
//    X    0    1    2
//    0    0    0    0
//    1    0    1    2
//    2    0    2    4
//
//  go run main.go 3
//    X    0    1    2    3
//    0    0    0    0    0
//    1    0    1    2    3
//    2    0    2    4    6
//    3    0    3    6    9
// ---------------------------------------------------------

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Give me the size of the table.")
		return
	}
	size, err := strconv.Atoi(os.Args[1])
	if err != nil || size < 1 {
		fmt.Println("Wrong size")
		return
	}
	fmt.Print("X")
	for i := 0; i <= size; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= size; i++ {
		fmt.Printf("%d", i)
		for j := 0; j <= size; j++ {
			if val := i * j; val != 0 {
				fmt.Printf("%5d", val)
			} else {
				fmt.Printf("%5s", "-")
			}
		}
		fmt.Println()
	}
	fmt.Println()

}
