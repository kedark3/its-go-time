package main

// ---------------------------------------------------------
// EXERCISE: Vowel or Consonant
//
//  Detect whether a letter is vowel or consonant.
//
// NOTE
//  y or w is called a semi-vowel.
//  Check out: https://en.oxforddictionaries.com/explore/is-the-letter-y-a-vowel-or-a-consonant/
//
// HINT
//  + You can find the length of an sument using the len function.
//
//  + len(os.ss[1]) will give you the length of the 1st sument.
//
// BONUS
//  Use strings.IndexAny function to detect the vowels.
//  Search on Google for: golang pkg strings IndexAny
//
// Furthermore, you can also use strings.ContainsAny. Check out: https://golang.org/pkg/strings/#ContainsAny
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a letter
//
//  go run main.go hey
//    Give me a letter
//
//  go run main.go a
//    "a" is a vowel.
//
//  go run main.go y
//    "y" is sometimes a vowel, sometimes not.
//
//  go run main.go w
//    "w" is sometimes a vowel, sometimes not.
//
//  go run main.go x
//    "x" is a consonant.
// ---------------------------------------------------------
import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 1 || len(args[1]) > 1 { // use unicode/utf-8 package for len if using utf-8 strings, but this code doesn't need to support that
		fmt.Println("Give me a letter")
		return
	}
	s := args[1]
	// Approach 1
	// if s == "a" || s == "e" || s == "i" || s == "o" || s == "u" {
	// 	fmt.Printf("%q is a vowel.\n", s)
	// } else if s == "w" || s == "y" {
	// 	fmt.Printf("%q is sometimes a vowel, sometimes not.\n", s)
	// } else {
	// 	fmt.Printf("%q is a consonant.\n", s)
	// }

	// Approach 2
	if strings.ContainsAny(s, "aeiou") {
		fmt.Printf("%q is a vowel.\n", s)
	} else if s == "w" || s == "y" {
		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", s)
	} else {
		fmt.Printf("%q is a consonant.\n", s)
	}
}
