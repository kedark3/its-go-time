package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Sum the Numbers
//
//  1. By using a loop, sum the numbers between 1 and 10.
//  2. Print the sum.
//
// EXPECTED OUTPUT
//  Sum: 55
// ---------------------------------------------------------

func main() {

	for i, sum := 1, 0; i <= 10; i++ {
		sum += i
		fmt.Printf("%d --> %d\n", i, sum)
	}
}
