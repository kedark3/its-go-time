package main

import "fmt"

func main() {
	var colors [2]string
	colors[0] = "Red"
	colors[1] = "Blue"

	fmt.Println("Colors array: ", colors)
	// declaring array with values
	var numbers = [5]int{5, 2, 3, 1, 4}
	fmt.Println(numbers[3], numbers, len(numbers))
}
