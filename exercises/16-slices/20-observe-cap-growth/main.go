package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Observe the capacity growth
//
//  Write a program that appends elements to a slice
//  10 million times in a loop. Observe how the capacity of
//  the slice changes.
//
//
// STEPS
//
//  1. Create a nil slice
//
//  2. Loop 1e7 times
//
//  3. On each iteration: Append an element to the slice
//
//  4. Print the length and capacity of the slice "only"
//     when its capacity changes.
//
//  BONUS: Print also the growth rate of the capacity.
//
//
// EXPECTED OUTPUT
//
//  len:0               cap:0               growth:NaN
//  len:1               cap:1               growth:+Inf
//  len:2               cap:2               growth:2.00
//  ... and so on.
//
// ---------------------------------------------------------

func main() {
	var (
		nums   []int
		oldCap float64
	)

	// loop 10 million times
	for len(nums) < 1e7 {
		// get the capacity
		c := float64(cap(nums))

		// only print when the capacity changes
		if c == 0 || c != oldCap {
			// print also the growth ratio: c/oldCap
			fmt.Printf("len:%-15d cap:%-15g growth:%-15.2f\n",
				len(nums), c, c/oldCap)
		}

		// keep track of the previous capacity
		oldCap = c

		// append an arbitrary element to the slice
		nums = append(nums, 1)
	}

}
