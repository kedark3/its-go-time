// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Basics
//
//  Let's warm up with the pointer basics. Please follow the
//  instructions inside the code. Run the solution to see
//  its output if you need to.
// ---------------------------------------------------------

package main

import "fmt"

type computer struct {
	brand string
}

func main() {
	// create a nil pointer with the type of pointer to a computer
	var c *computer

	// compare the pointer variable to a nil value, and say it's nil
	if c == nil {
		fmt.Printf("c is a %v pointer.\n", c)
	}
	// create an apple computer by putting its address to a pointer variable
	apple := &computer{brand: "Apple"}
	// put the apple into a new pointer variable
	clone := apple
	// compare the apples: if they are equal say so and print their addresses
	if apple == clone {
		fmt.Printf("Apple %p and Clone %p are equal\n", apple, clone)
	}
	// create a sony computer by putting its address to a new pointer variable
	sony := &computer{brand: "Sony"}
	// compare apple to sony, if they are not equal say so and print their
	// addresses
	if apple != sony {
		fmt.Printf("Apple %p and Sony %p are not equal\n", apple, sony)
	}
	// put apple's value into a new ordinary variable
	v := *apple
	// print apple pointer variable's address, and the address it points to
	// and, print the new variable's addresses as well
	fmt.Printf("apple pointer variable's address %p, the address it points to %p and the new variable's address %p\n ",
		&apple, apple, &v)
	// compare the value that is pointed by the apple and the new variable
	// if they are equal say so
	// print the values:
	// the value that is pointed by the apple and the new variable
	if *apple == v {
		fmt.Printf("Apple %v and V %v are equal\n", *apple, v)
	}

	// change sony's brand to hp using the func — print sony's brand
	change(*sony, "hp")     // This line doesn't change outcome of following print statement
	fmt.Println(sony.brand) // because sony was passed using *sony, hence its value was passed, and changing value in func have no effect in main
	// write a func that returns the value that is pointed by the given *computer
	// print the returned value
	fmt.Println(apple, value(apple))
	fmt.Println(sony, value(sony))
	// write a new constructor func that returns a pointer to a computer
	// and call the func 3 times and print the returned values' addresses
	for i := 0; i < 3; i++ {
		fmt.Printf("%p\n", newComputer())
	}
}

// create a new function: change
// the func can change the given computer's brand to another brand
func change(c computer, brand string) {
	c.brand = brand
	fmt.Println(c)
}

// write a func that returns the value that is pointed by the given *computer
func value(c *computer) computer {
	return *c
}

func newComputer() *computer {
	return &computer{}
}
