package main

import (
	"fmt"
)

func main() {
	fmt.Printf("I'm %d years old.\n", 30)

	// ---------------------------------------------------------
	// EXERCISE: Print Your Name and LastName
	//
	//  Print your name and lastname using Printf
	//
	// EXPECTED OUTPUT
	//  My name is Inanc and my lastname is Gumus.
	//
	// BONUS
	//  Store the formatting specifier (first argument of Printf)
	//    in a variable.
	//  Then pass it to printf
	// ---------------------------------------------------------

	// BONUS: Use a variable for the format specifier
	formatter := "My name is %s and my lastname is %s.\n"
	fmt.Printf(formatter, "kedar", "kulkarni")
	// ---------------------------------------------------------
	// EXERCISE: False Claims
	//
	//  Use printf to print the expected output using a variable.
	//
	// EXPECTED OUTPUT
	//  These are false claims.
	// ---------------------------------------------------------

	// UNCOMMENT THE FOLLOWING CODE
	// AND DO NOT CHANGE IT AFTERWARDS
	tf := false

	// TYPE YOUR CODE HERE
	fmt.Printf("These are %t claims.", tf)

	// ---------------------------------------------------------
	// EXERCISE: Print the Temperature
	//
	//  Print the current temperature in your area using Printf
	//
	// NOTE
	//  Do not use %v verb
	//  Output "shouldn't" be like 29.500000
	//
	// EXPECTED OUTPUT
	//  Temperature is 29.5 degrees.
	// ---------------------------------------------------------

	fmt.Printf("Temperature is %.1f degrees.\n", 29.5)
	// ---------------------------------------------------------
	// EXERCISE: Double Quotes
	//
	//  Print "hello world" with double-quotes using Printf
	//  (As you see here)
	//
	// NOTE
	//  Output "shouldn't" be just: hello world
	//
	// EXPECTED OUTPUT
	//  "hello world"
	// ---------------------------------------------------------
	fmt.Printf("%q\n", "hello world")

	// ---------------------------------------------------------
	// EXERCISE: Print the Type
	//
	//  Print the type and value of 3 using fmt.Printf
	//
	// EXPECTED OUTPUT
	//  Type of 3 is int
	// ---------------------------------------------------------
	fmt.Printf("Type of %d is %[1]T\n", 3)

	// ---------------------------------------------------------
	// EXERCISE: Print the Type #2
	//
	//  Print the type and value of 3.14 using fmt.Printf
	//
	// EXPECTED OUTPUT
	//  Type of 3.14 is float64
	// ---------------------------------------------------------
	fmt.Printf("Type of %.2f is %[1]T\n", 3.14)

	// ---------------------------------------------------------
	// EXERCISE: Print the Type #3
	//
	//  Print the type and value of "hello" using fmt.Printf
	//
	// EXPECTED OUTPUT
	// 	Type of hello is string
	// ---------------------------------------------------------
	fmt.Printf("Tpe of %s is %[1]T.\n", "hello")

	// ---------------------------------------------------------
	// EXERCISE: Print the Type #4
	//  Print the type and value of true using fmt.Printf
	//
	// EXPECTED OUTPUT
	//  Type of true is bool
	// ---------------------------------------------------------
	fmt.Printf("Type of %t is %[1]T.\n", true)

}
