package main

import (
	"fmt"
	"sort"
)

func main() {
	// make a new map, `make` function initializes map with memory
	// If `make` is not used, trying to add value to map would end up in panic
	m := make(map[string]string)
	// var m map[string]string // This would result in nil map panic
	// One can use `new` function as well, but new will just return memory addr, UNinitialized

	m["NC"] = "North Carolina"
	m["NY"] = "New York"
	m["AK"] = "Arkansas"
	m["OK"] = "Oklahoma"
	m["MH"] = "Maharashtra" // Dil se Desi :P ;)

	fmt.Println(m)

	// To print map in order, we need to jump through few hoops
	// Extract Keys in a slice
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	// now that Keys are populated, sort those
	sort.Strings(keys)
	fmt.Println(keys)
	// once Keys are sorted we can iterate over keys to print map in ascending order
	fmt.Println("\nSorted:")
	for k := range keys {
		fmt.Println(keys[k], m[keys[k]])
	}

}
