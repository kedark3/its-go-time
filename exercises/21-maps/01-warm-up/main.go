package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Warm-up
//
//  Create and print the following maps.
//
//  1. Phone numbers by last name
//  2. Product availability by Product ID
//  3. Multiple phone numbers by last name
//  4. Shopping basket by Customer ID
//
//     Each item in the shopping basket has a Product ID and
//     quantity. Through the map, you can tell:
//     "Mr. X has bought Y bananas"
//
// ---------------------------------------------------------

func main() {
	// Hint: Store phone numbers as text

	// #1
	// Key        : Last name
	// Element    : Phone number
	phoneNums := map[string]string{"Doe": "1234567", "Smith": "999-888-7765", "Singh": "678-901-0123"}

	// #2
	// Key        : Product ID
	// Element    : Available / Unavailable
	products := map[int]bool{1: true, 2: false}

	// #3
	// Key        : Last name
	// Element    : Phone numbers
	var phoneNumsMulti map[string][]string

	// #4
	// Key        : Customer ID
	// Element Key:
	//   Key: Product ID Element: Quantity
	var cuBasket map[int]map[int]int
	fmt.Printf("phones      : %#v\n", phoneNums)
	fmt.Printf("products    : %#v\n", products)
	fmt.Printf("phonesMulti : %#v\n", phoneNumsMulti)
	fmt.Printf("Basket      : %#v\n", cuBasket)

}
