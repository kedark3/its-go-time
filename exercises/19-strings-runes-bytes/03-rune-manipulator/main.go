package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Rune Manipulator
//
//  Please read the comments inside the following code.
//
// EXPECTED OUTPUT
//  Please run the solution.
// ---------------------------------------------------------

func main() {
	words := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	for _, w := range words {
		// Print the byte and rune length of the strings
		// Hint: Use len and utf8.RuneCountInString
		fmt.Printf("\n%q has len: %d, RuneCount: %d\n", w, len(w), utf8.RuneCountInString(w))
		// Print the bytes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Printf("%q can be represented in hex as \"% [1]x\"\n ", w)

		// Print the runes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Printf("Runes of the %q in hexadecimal\n", w)
		for _, ch := range w { // this for range loops over the "runes" and not "bytes" , it is the default behavior of a for range loop
			fmt.Printf("% x", ch)
		}
		// Print the runes of the strings as rune literals
		// Hint: Use for range

		fmt.Printf("\nRunes of the %q in rune literals\n", w)
		for _, ch := range w { // this for range loops over the "runes" and not "bytes" , it is the default behavior of a for range loop
			fmt.Printf("%q", ch)
		}

		// Print the first rune and its byte size of the strings
		// Hint: Use utf8.DecodeRuneInString
		r, s := utf8.DecodeRuneInString(w)
		fmt.Printf("\n first Rune: %q, its size is %d\n", r, s)

		// Print the last rune of the strings
		// Hint: Use utf8.DecodeLastRuneInString
		r, s = utf8.DecodeLastRuneInString(w)
		fmt.Printf("last Rune: %q, its size is %d\n", r, s)
		// Slice and print the first two runes of the strings
		_, first := utf8.DecodeRuneInString(w)
		_, second := utf8.DecodeRuneInString(w[first:])
		fmt.Printf("first 2 : %q\n", w[:first+second])
		// Slice and print the last two runes of the strings
		_, last1 := utf8.DecodeLastRuneInString(w)
		_, last2 := utf8.DecodeLastRuneInString(w[:len(w)-last1])
		fmt.Printf("last 2  : %q\n", w[len(w)-last2-last1:])

		// Convert the string to []rune
		// Print the first and last two runes
		rs := []rune(w)
		fmt.Printf("first 2 : %q\n", string(rs[:2]))
		fmt.Printf("last 2  : %q\n", string(rs[len(rs)-2:]))

		fmt.Println("\n", strings.Repeat("*", 80))
	}

	// for _, s := range words {
	// 	fmt.Printf("%q\n", s)

	// 	// Print the byte and rune length of the strings
	// 	// Hint: Use len and utf8.RuneCountInString
	// 	fmt.Printf("\thas %d bytes and %d runes\n",
	// 		len(s), utf8.RuneCountInString(s))

	// 	// Print the bytes of the strings in hexadecimal
	// 	// Hint: Use % x verb
	// 	fmt.Printf("\tbytes   : % x\n", s)

	// 	// Print the runes of the strings in hexadecimal
	// 	// Hint: Use % x verb
	// 	fmt.Print("\trunes   :")
	// 	for _, r := range s {
	// 		fmt.Printf("% x", r)
	// 	}
	// 	fmt.Println()

	// 	// Print the runes of the strings as rune literals
	// 	// Hint: Use for range
	// 	fmt.Print("\trunes   :")
	// 	for _, r := range s {
	// 		fmt.Printf("%q", r)
	// 	}
	// 	fmt.Println()

	// 	// Print the first rune and its byte size of the strings
	// 	// Hint: Use utf8.DecodeRuneInString
	// 	r, size := utf8.DecodeRuneInString(s)
	// 	fmt.Printf("\tfirst   : %q (%d bytes)\n", r, size)

	// 	// Print the last rune of the strings
	// 	// Hint: Use utf8.DecodeLastRuneInString
	// 	r, size = utf8.DecodeLastRuneInString(s)
	// 	fmt.Printf("\tlast    : %q (%d bytes)\n", r, size)

	// 	// Slice and print the first two runes of the strings
	// _, first := utf8.DecodeRuneInString(s)
	// _, second := utf8.DecodeRuneInString(s[first:])
	// fmt.Printf("\tfirst 2 : %q\n", s[:first+second])

	// 	// Slice and print the last two runes of the strings
	// _, last1 := utf8.DecodeLastRuneInString(s)
	// _, last2 := utf8.DecodeLastRuneInString(s[:len(s)-last1])
	// fmt.Printf("\tlast 2  : %q\n", s[len(s)-last2-last1:])

	// 	// Convert the string to []rune
	// 	// Print the first and last two runes
	// rs := []rune(s)
	// fmt.Printf("\tfirst 2 : %q\n", string(rs[:2]))
	// fmt.Printf("\tlast 2  : %q\n", string(rs[len(rs)-2:]))
	// }
}
