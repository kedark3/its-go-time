package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors [2]string
	colors[0] = "Red"
	colors[1] = "Blue"

	fmt.Println("Colors array: ", colors)
	// declaring array with values
	var numbers = [5]int{5, 2, 3, 1, 4}
	fmt.Println(numbers[3], numbers, len(numbers))
	// Declaring a Slice
	var colorsSlice = []string{"Red", "Green"}
	fmt.Println(colorsSlice)
	colorsSlice = append(colorsSlice, "Blue")
	fmt.Println(colorsSlice)
	fmt.Println("Slicing:", colorsSlice[2:])
	colorsSlice = append(colorsSlice, "BigBlue")
	colorsSlice = append(colorsSlice, "BigB")
	colorsSlice = append(colorsSlice, "BigN")
	fmt.Println("Slicing:", colorsSlice[2:len(colorsSlice)-2])
	fmt.Println(colorsSlice)
	sort.Strings(colorsSlice)
	fmt.Println(colorsSlice)

}
