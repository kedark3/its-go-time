package main

import "fmt"

// Solves https://leetcode.com/problems/two-sum/
func main() {
	nums := []int{3, 3}
	target := 6
	fmt.Println(twoSumBruteForce(nums, target))
	fmt.Println(twoSumTwoPassHashMap(nums, target))
	fmt.Println(twoSumSinglePassHashMap(nums, target))

}

func twoSumBruteForce(nums []int, target int) []int {
	r := make([]int, 2, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				r[0] = i
				r[1] = j
				break
			}
		}
	}
	return r
}

func twoSumTwoPassHashMap(nums []int, target int) []int {
	m := make(map[int]int)
	r := make([]int, 2, 2)
	for index, val := range nums {
		m[val] = index
	}
	for i, val := range nums {
		complement := target - val
		if v, ok := m[complement]; ok && m[complement] != i {
			r[0] = v
			r[1] = i
		}
	}
	return r
}

func twoSumSinglePassHashMap(nums []int, target int) []int {
	m := make(map[int]int)
	r := make([]int, 2, 2)
	for index, val := range nums {

		if index == 0 {
			m[val] = index
			continue
		}
		complement := target - val
		if v, ok := m[complement]; ok {
			r[0] = v
			r[1] = index
		}
		m[val] = index
	}
	return r
}
