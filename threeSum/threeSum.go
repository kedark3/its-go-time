package main

// Need to review other approaches
import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	// Create a map for 2Sum problem
	var output [][]int
	sort.Ints(nums)
	for i, num := range nums {
		if num >= 0 {
			break
		}
		if i == 0 || nums[i-1] != nums[i] {
			// fmt.Println("Calling 1st time", nums, i)
			twoSum(nums, i, &output)
		}
	}
	return output

}

func twoSum(nums []int, index int, output *[][]int) [][]int {
	lo, hi := index+1, len(nums)-1
	fmt.Println(index, lo, hi)
	for lo < hi {
		sum := nums[index] + nums[lo] + nums[hi]
		if sum == 0 {
			fmt.Println(nums[index], nums[lo], nums[hi])
			*output = append(*output, []int{nums[index], nums[lo], nums[hi]})
			lo++
			hi--
			for lo < hi && nums[lo] == nums[lo-1] {
				lo++
			}
			continue
		} else if sum < 0 {
			lo++
		} else if sum > 0 {
			hi--
		}

	}
	return *output
}

func main() {

	fmt.Print(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	// -4, -1, -1, 0, 1, 2
}
