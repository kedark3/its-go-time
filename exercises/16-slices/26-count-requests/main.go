package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Print daily requests
//
//  You've got request logs of a web server. The log data
//  contains 8-hourly totals per each day. It is stored
//  in the `reqs` slice.
//
//  Find and print the total requests per day, as well as
//  the grand total.
//
//  See the `reqs` slice and the steps in the code below.
//
//
// RESTRICTIONS
//
//   1. You need to produce the daily slice, don't just loop
//      and print the element totals directly. The goal is
//      gaining more experience in slice operations.
//
//   2. Your code should work even if you add to or remove the
//      existing elements from the `reqs` slice.
//
//      For example, after solving the exercise, try it with
//      this new data:
//
//  reqs := []int{
//       500, 600, 250,
//       200, 400, 50,
//       900, 800, 600,
//       750, 250, 100,
//       150, 654, 235,
//       320, 534, 765,
//       121, 876, 285,
//       543, 642,
//       // the last element is missing (your code should be able to handle this)
//       // that is why you shouldn't calculate the `size` below manually.
//  }
//
//      The grand total of the new data should be 10525.
//
//
// EXPECTED OUTPUT
//
//   Please run `solution/main.go` to see the expected
//   output.
//
//   go run solution/main.go
//
// ---------------------------------------------------------

func main() {
	// There are 3 request totals per day (8-hourly)
	const N = 3

	// DAILY REQUESTS DATA (8-HOURLY TOTALS PER DAY)
	reqs := []int{
		500, 600, 250,
		200, 400, 50,
		900, 800, 600,
		750, 250, 100,
		150, 654, 235,
		320, 534, 765,
		121, 876, 285,
		543, 642,
		// the last element is missing (your code should be able to handle this)
		// that is why you shouldn't calculate the `size` below manually.
	}
	// reqs := []int{
	// 	500, 600, 250, // 1st day: 1350 requests
	// 	200, 400, 50, // 2nd day: 650 requests
	// 	900, 800, 600, // 3rd day: 2300 requests
	// 	750, 250, 100, // 4th day: 1100 requests
	// 	// grand total: 5400 requests
	// }

	// ================================================
	// #1: Make a new slice with the exact size needed.

	size := len(reqs) / N // you need to find the size.
	daily := make([][]int, 0, size)

	// ================================================

	// ================================================
	// #2: Group the `reqs` per day into the slice: `daily`.
	//
	// So the daily will be:
	// [
	//  [500, 600, 250]
	//  [200, 400, 50]
	//  [900, 800, 600]
	//  [750, 250, 100]
	// ]

	for i := 0; i < len(reqs); i += N {
		if i+N < len(reqs) {
			daily = append(daily, reqs[i:i+N])
		} else {

			daily = append(daily, reqs[i:i+len(reqs)%N])
		}
	}
	// ================================================
	// #3: Print the results

	// Print a header
	fmt.Printf("%-10s%-10s\n", "Day", "Requests")
	fmt.Println(strings.Repeat("=", 20))

	// Loop over the daily slice and its inner slices to find
	// the daily totals and the grand total.
	var grandSum int
	for i, day := range daily {
		var sum int
		for _, reqCount := range day {
			fmt.Printf("%d%15d\n", i+1, reqCount)
			sum += reqCount
		}
		grandSum += sum
		fmt.Printf("%15s: %d\n", "TOTAL", sum)
	}
	fmt.Printf("\n%15s: %d\n", "GRAND TOTAL", grandSum)
	// ================================================
}
