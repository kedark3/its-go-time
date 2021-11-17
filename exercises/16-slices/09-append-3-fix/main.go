package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Append #3 — Fix the problems
//
//  Fix the problems in the code below.
//
// BONUS
//
//  Simplify the code.
//
// EXPECTED OUTPUT
//
//  toppings: [black olives green peppers onions extra cheese]
//
// ---------------------------------------------------------

func main() {
	// toppings := []int{"black olives", "green peppers"}

	// var pizza [3]string
	// append(pizza, ...toppings)
	// pizza = append(toppings, "onions")
	// toppings = append(pizza, extra cheese)

	// fmt.Printf("pizza       : %s\n", pizza)
	toppings := []string{"black olives", "green peppers"}

	var pizza []string
	pizza = append(pizza, toppings...)
	toppings = append(pizza, "onions", "extra", "cheese")

	fmt.Printf("pizza       : %s\n", toppings)
}
