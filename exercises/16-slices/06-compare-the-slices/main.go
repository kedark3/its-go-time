package main

import (
	"fmt"
	"sort"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Compare the slices
//
//  1. Split the namesA string and get a slice
//
//  2. Sort all the slices
//
//  3. Compare whether the slices are equal or not
//
//
// EXPECTED OUTPUT
//
//   They are equal.
//
//
// HINTS
//
//   1. strings.Split function splits a string and
//      returns a string slice
//
//   2. Comparing slices: First check whether their length
//      are the same or not; only then compare them.
//
// ---------------------------------------------------------

func main() {
	namesA := strings.Split("Da Vinci, Wozniak, Carmack", ", ")
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}
	ok := "equal"
	sort.Strings(namesA)
	sort.Strings(namesB)
	fmt.Println(namesA)
	fmt.Println(namesB)
	if len(namesA) == len(namesB) {
		for i := range namesA {
			if namesA[i] != namesB[i] {
				return
			}
		}
		fmt.Printf("They are %s\n", ok)
	}
}
