package main

import (
	"fmt"
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Sort the backing array
//
//  1. Sort only the middle 3 items.
//
//  2. All the slices should see your changes.
//
//
// RESTRICTION
//
//  Do not sort manually. Sort by using the sort package.
//
//
// EXPECTED OUTPUT
//
//  Original: [pacman mario tetris doom galaga frogger asteroids simcity metroid defender rayman tempest ultima]
//
//  Sorted  : [pacman mario tetris doom galaga asteroids frogger simcity metroid defender rayman tempest ultima]
//
//
// HINT:
//
//   Middle items are         : [frogger asteroids simcity]
//
//   After sorting they become: [asteroids frogger simcity]
//
// ---------------------------------------------------------

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}

	fmt.Println("Original:", items)
	// ADD YOUR CODE HERE
	mid := len(items) / 2
	smid := items[mid-1 : mid+2]
	sort.Strings(smid)
	fmt.Println()
	fmt.Println("Sorted  :", items)
}
