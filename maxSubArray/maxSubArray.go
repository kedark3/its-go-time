package main

import "fmt"

// Use Greedy approach
func maxSubArraySum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	currSum, maxSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		currSum = max(nums[i], nums[i]+currSum)
		maxSum = max(currSum, maxSum)

	}
	return maxSum
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func main() {
	fmt.Println(maxSubArraySum([]int{1, 2, 3, 4, -2, -2}))
}
