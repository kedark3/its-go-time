package main

import "fmt"

// Leetcode https://leetcode.com/problems/running-sum-of-1d-array/
func runningSum(nums []int) []int {
	// :results: to hold output slice
	var results []int
	// :runningSum: to hold current sum
	var runningSum int
	for _, num := range nums {
		runningSum += num
		results = append(results, runningSum)
	}
	return results

}

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4, 5}))
}
