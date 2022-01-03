package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Uppercaser
//
//  Use a scanner to convert the lines to uppercase, and
//  print them.
//
//  1. Feed the shakespeare.txt to your program.
//
//  2. Scan the input using a new Scanner.
//
//  3. Print each line.
//
// EXPECTED OUTPUT
//  Please run the solution to see the expected output.
// ---------------------------------------------------------

func main() {
	//os.Stdin.Close()  // You can close the Stdin as it is a file as is everything else is in Linux
	in := bufio.NewScanner(os.Stdin)

	in.Split(bufio.ScanLines)

	for in.Scan() {
		fmt.Println(strings.ToUpper(in.Text()))
	}
	// Closing stdin would cause the error, which will get printed here. But it may print any other error that may occur with bufio scanner
	if err := in.Err(); err != nil {
		fmt.Printf("> Err: %v", err)
	}
}
