package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Slice the numbers
//
//   We've a string that contains even and odd numbers.
//
//   1. Convert the string to an []int
//
//   2. Print the slice
//
//   3. Slice it for the even numbers and print it (assign it to a new slice variable)
//
//   4. Slice it for the odd numbers and print it (assign it to a new slice variable)
//
//   5. Slice it for the two numbers at the middle
//
//   6. Slice it for the first two numbers
//
//   7. Slice it for the last two numbers (use the len function)
//
//   8. Slice the evens slice for the last one number
//
//   9. Slice the odds slice for the last two numbers
//
//
// EXPECTED OUTPUT
//   go run main.go
//
//     nums        : [2 4 6 1 3 5]
//     evens       : [2 4 6]
//     odds        : [1 3 5]
//     middle      : [6 1]
//     first 2     : [2 4]
//     last 2      : [3 5]
//     evens last 1: [6]
//     odds last 2 : [3 5]
//
//
// NOTE
//
//  You can also use my prettyslice package for printing the slices.
//
//
// HINT
//
//  Find a function in the strings package for splitting the string into []string
//
// ---------------------------------------------------------

func main() {
	// uncomment the declaration below
	data := "2 4 6 1 3 5"
	var nums []int
	for _, n := range strings.Split(data, " ") {
		num, _ := strconv.Atoi(n)
		nums = append(nums, num)
	}
	fmt.Printf("nums         : %v\n", nums)
	fmt.Printf("evens        : %v\n", nums[:3])
	fmt.Printf("odds         : %v\n", nums[3:])
	fmt.Printf("middle       : %v\n", nums[2:4])
	fmt.Printf("first 2      : %v\n", nums[:2])
	fmt.Printf("last 2       : %v\n", nums[len(nums)-2:])
	fmt.Printf("evens last 1 : %v\n", nums[2:3])
	fmt.Printf("odds last 2  : %v\n", nums[4:])
}
