package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Find the Average
//
//  Your goal is to fill an array with numbers and find the average.
//
//  1. Get the numbers from the command-line.
//
//  2. Create an array and assign the given numbers to that array.
//
//  3. Print the given numbers and their average.
//
// RESTRICTION
//   + Maximum 5 numbers can be provided
//   + If one of the arguments are not a valid number, skip it
//
// EXPECTED OUTPUT
//   go run main.go
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5 6
//     Please tell me numbers (maximum 5 numbers).
//
//   go run main.go 1 2 3 4 5
//     Your numbers: [1 2 3 4 5]
//     Average: 3
//
//   go run main.go 1 a 2 b 3
//     Your numbers: [1 0 2 0 3]
//     Average: 2
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]
	if len(args) == 0 || len(args) > 5 {
		fmt.Println("Please tell me numbers (maximum 5 numbers).")
		return
	}
	var nums [5]int
	var sum, count int
	for i, item := range args {
		n, err := strconv.Atoi(item)
		if err == nil {
			nums[i] = n
			sum += n
			count += 1
		}
	}
	fmt.Println("Your numbers:", nums)
	fmt.Println("Average:", sum/count)
}
