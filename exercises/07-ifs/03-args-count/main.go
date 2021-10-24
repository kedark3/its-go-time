package main

// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go I wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	l := len(args)
	if l == 0 {
		fmt.Println("Give me args")
	} else if l == 1 {
		fmt.Printf("There is one: %q\n", args[0])
	} else if l == 2 {
		fmt.Printf("There are two: \"%s %s\"\n", args[0], args[1])
	} else {
		fmt.Printf("There are %d arguments\n", l)
	}
}
