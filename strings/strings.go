package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "My String"
	fmt.Printf("str1: %v : %T\n", str1, str1)

	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.ToLower(str1))

	lValue := "Hello"
	rValue := "HELLO"

	// case sensitive comparison
	fmt.Println("Equal?", (lValue == rValue))
	// non-case sensitive comparison
	fmt.Println("Equal Non-case-Sensitive?", strings.EqualFold(lValue, rValue))

	fmt.Println("Contains Str?", strings.Contains(str1, "Str"))
}
