package main

import (
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Populate and Lookup
//
//  Add elements to the maps that you've declared in the
//  first exercise, and try them by looking up for the keys.
//
//  Either use the `make()` or `map literals`.
//
//  After completing the exercise, remove the data and check
//  that your program still works.
//
//
//  1. Phone numbers by last name
//     --------------------------
//     bowen  202-555-0179
//     dulin  03.37.77.63.06
//     greco  03489940240
//
//     Print the dulin's phone number.
//
//
//  2. Product availability by Product ID
//     ----------------
//     617841573 true
//     879401371 false
//     576872813 true
//
//     Is Product ID 879401371 available?
//
//
//  3. Multiple phone numbers by last name
//     ------------------------------------------------------
//     bowen  [202-555-0179]
//     dulin  [03.37.77.63.06 03.37.70.50.05 02.20.40.10.04]
//     greco  [03489940240 03489900120]
//
//     What is Greco's second phone number?
//
//
//  4. Shopping basket by Customer ID
//     -------------------------------
//     100 [617841573:4 576872813:2]
//     101 [576872813:5 657473833:20]
//     102 []
//
//     How many of 576872813 the customer 101 is going to buy?
//                (Product ID)  (Customer ID)
//
//
// EXPECTED OUTPUT
//
//   1. Run the solution to see the output
//   2. Here is the output with empty maps:
//
//      dulin's phone number: N/A
//      Product ID #879401371 is not available
//      greco's 2nd phone number: N/A
//      Customer #101 is going to buy 5 from Product ID #576872813.
//
// ---------------------------------------------------------

// type StringHeader struct {
// 	pointer uintptr
// 	length  int
// }

func main() {

	var (
		phoneNums map[string]string

		products map[int]bool

		phoneNumsMulti map[string][]string

		basket map[int]map[int]int
	)
	// var (
	// 	phoneNums = map[string]string{
	// 		"bowen": "202-555-0179",
	// 		"dulin": "03.37.77.63.06",
	// 		"greco": "03489940240",
	// 	}

	// 	products = map[int]bool{
	// 		617841573: true,
	// 		879401371: false,
	// 		576872813: true,
	// 	}

	// 	phoneNumsMulti = map[string][]string{
	// 		"bowen": {"202-555-0179"},
	// 		"dulin": {"03.37.77.63.06", "03.37.70.50.05", "02.20.40.10.04"},
	// 		"greco": {"03489940240", "03489900120"},
	// 	}
	// 	basket = map[int]map[int]int{
	// 		100: {617841573: 4, 576872813: 2},
	// 		101: {576872813: 5, 657473833: 20},
	// 		102: {},
	// 	}
	// )

	dulin, ok := phoneNums["dulin"]

	if !ok {
		dulin = "N/A"
	}
	fmt.Printf("dulin's phone number: %s\n", dulin)

	_, ok = products[87901371]

	s := "available"
	// ptr := *(*StringHeader)(unsafe.Pointer(&s))
	// fmt.Printf("%q %20d: %+v \n", s, &s, ptr)

	if !ok {
		s = "not available"
	}

	// ptr = *(*StringHeader)(unsafe.Pointer(&s))
	// we can see that the pointer of var `s` doesn't change but the ptr in string header changes if we update value of `s`
	// This means that overriding string literal causes new backend array allocation and Go GC will reclaim old value
	// fmt.Printf("%q %16d: %+v \n", s, &s, ptr)
	/*output for this portion would be:

	"available"         824633786944: {pointer:4814129 length:9}
	"not available"     824633786944: {pointer:4817040 length:13}

	*/
	fmt.Printf("Product ID #879401371 is %s\n", s)

	who, phone := "greco", "N/A"
	if phones := phoneNumsMulti[who]; len(phones) >= 2 {
		phone = phones[1]
	}
	fmt.Printf("%s's 2nd phone number: %s\n", who, phone)
	cid, pid := 101, 576872813
	fmt.Printf("Customer #%d is going to buy %d from Product ID #%d.\n", cid, basket[cid][pid], pid)
}

// func dump(s string) {
// 	ptr := *(*StringHeader)(unsafe.Pointer(&s))
// 	fmt.Printf("%q %d: %+v \n", s, &s, ptr)
// }
