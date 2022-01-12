// ---------------------------------------------------------
// EXERCISE: Swap
//
//  Using funcs, swap values through pointers, and swap
//  the addresses of the pointers.
//
//
//  1- Swap the values through a func
//
//     a- Declare two float variables
//
//     b- Declare a func that can swap the variables' values
//        through their memory addresses
//
//     c- Pass the variables' addresses to the func
//
//     d- Print the variables
//
//
//  2- Swap the addresses of the float pointers through a func
//
//     a- Declare two float pointer variables and,
//        assign them the addresses of float variables
//
//     b- Declare a func that can swap the addresses
//        of two float pointers
//
//     c- Pass the pointer variables to the func
//
//     d- Print the addresses, and values that are
//        pointed by the pointer variables
//
// ---------------------------------------------------------

package main

import "fmt"

func main() {
	f1, f2 := 1.0, 2.0
	swapVal(&f1, &f2)
	fmt.Println(f1, f2)

	f1ptr, f2ptr := &f1, &f2
	fmt.Println(f1ptr, f2ptr)

	f1ptr, f2ptr = swapAddr(f1ptr, f2ptr)
	fmt.Println(f1ptr, f2ptr)
}

func swapVal(f1, f2 *float64) {
	// function receives addresses(pointers) to variable f1,f2
	// it then uses * to dereference values at those addresses and
	// swap the values. As the values are swapped at the memory locations
	// the change reflects back in the main function, without having to return the value
	*f1, *f2 = *f2, *f1
}

func swapAddr(f1, f2 *float64) (*float64, *float64) {
	return f2, f1
}
