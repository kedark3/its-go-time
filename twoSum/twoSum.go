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

func twoSumTwoPointersSortedArray(numbers []int, target int) []int {
	// Since the input is sorted in increasing order
	// we use two pointers traversing over the slice
	lo, hi := 0, len(numbers)-1
	for i := 0; i < len(numbers); i++ {
		// if lo and hi are same
		// we need to break out with no solution
		if lo == hi {
			break
		} else if numbers[lo]+numbers[hi] == target {
			// if we find solution we return it
			// but add +1 as we need to have 1 based indices
			return []int{lo + 1, hi + 1}
		} else if numbers[lo]+numbers[hi] < target {
			// if addition is less than target
			// we need to increase `lo` by one
			lo++
		} else if numbers[lo]+numbers[hi] > target {
			// if addition is less than target
			// we need to decrease `hi` by one
			hi--
		}

	}
	return []int{}
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
