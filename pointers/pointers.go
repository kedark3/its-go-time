package main

import "fmt"

// NOTE: Pointers in GoLang are pretty much exactly what you'd see in C.
func main() {
	var p *int

	if p != nil {
		fmt.Println("Value of Pointer p is: ", *p)
	} else {
		fmt.Println("P is nil")
	}

	var v int = 42
	p = &v
	if p != nil {
		fmt.Println("Value of Pointer p is: ", *p)
	} else {
		fmt.Println("P is nil")
	}
	*p = *p * 5
	fmt.Println("Value of Pointer p is: ", *p)
	fmt.Println("Value of Pointer v is: ", v)

	var p1, p2 *int
	v1, v2 := 5, 10
	p1, p2 = &v1, &v2
	fmt.Println(*p1 + *p2)
}
