package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Windows Path
//
//  1. Change the following program
//  2. It should use a raw string literal instead
//
// HINT
//  Run this program first to see its output.
//  Then you can easily understand what it does.
//
// EXPECTED OUTPUT
//  Your solution should output the same as this program.
//  Only that it should use a raw string literal instead.
// ---------------------------------------------------------

func main() {
	// HINTS:
	// \\ equals to backslash character
	// \n equals to newline character

	path := "c:\\program files\\duper super\\fun.txt\n" +
		"c:\\program files\\really\\funny.png"
	fmt.Println(path)

	path = `
c:\program files\duper super\fun.txt
c:\program files\really\funny.png
	`
	fmt.Println(path)

	// ---------------------------------------------------------
	// EXERCISE: Print JSON
	//
	//  1. Change the following program
	//  2. It should use a raw string literal instead
	//
	// HINT
	//  Run this program first to see its output.
	//  Then you can easily understand what it does.
	//
	// EXPECTED OUTPUT
	//  Your solution should output the same as this program.
	//  Only that it should use a raw string literal instead.
	// ---------------------------------------------------------

	// HINTS:
	// \t equals to TAB character
	// \n equals to newline character
	// \" equals to double-quotes character

	json := "\n" +
		"{\n" +
		"\t\"Items\": [{\n" +
		"\t\t\"Item\": {\n" +
		"\t\t\t\"name\": \"Teddy Bear\"\n" +
		"\t\t}\n" +
		"\t}]\n" +
		"}\n"

	fmt.Println(json)

	json = `
{
	"Items": [{
		"Item": {
			"name": "Teddy Bear"
		}
	}]
}	`
	fmt.Println(json)

	// ---------------------------------------------------------
	// EXERCISE: Raw Concat
	//
	//  1. Initialize the name variable
	//     by getting input from the command line
	//
	//  2. Use concatenation operator to concatenate it
	//     with the raw string literal below
	//
	// NOTE
	//  You should concatenate the name variable in the correct
	//  place.
	//
	// EXPECTED OUTPUT
	//  Let's say that you run the program like this:
	//   go run main.go inanç
	//
	//  Then it should output this:
	//   hi inanç!
	//   how are you?
	// ---------------------------------------------------------
	// uncomment the code below
	// name := "and get the name from the command-line"

	// replace and concatenate the `name` variable
	// after `hi ` below
	myname := os.Args[1]
	msg := `hi ` + myname + `!
how are you?`

	fmt.Println(msg)

	// ---------------------------------------------------------
	// EXERCISE: Count the Chars
	//
	//  1. Change the following program to work with unicode
	//     characters.
	//
	// INPUT
	//  "İNANÇ"
	//
	// EXPECTED OUTPUT
	//  5
	// ---------------------------------------------------------

	// Currently it returns 7
	// Because, it counts the bytes...
	// It should count the runes (codepoints) instead.
	//
	// When you run it with "İNANÇ", it should return 5 not 7.

	length := len(os.Args[1])
	fmt.Println(length)
	fmt.Println(utf8.RuneCountInString(os.Args[1]))

	// ---------------------------------------------------------
	// EXERCISE: Improved Banger
	//
	//  Change the Banger program the work with Unicode
	//  characters.
	//
	// INPUT
	//  "İNANÇ"
	//
	// EXPECTED OUTPUT
	//  İNANÇ!!!!!
	// ---------------------------------------------------------
	msg = os.Args[1]

	s := msg + strings.Repeat("!", len(msg))

	fmt.Println(s)
	s = msg + strings.Repeat("!", utf8.RuneCountInString(os.Args[1]))
	fmt.Println(s)

	fmt.Println(strings.ToLower(os.Args[1]))

	msg = `
	
	The weather looks good.
I should go and play.
	`
	fmt.Println(strings.TrimSpace(msg))

	name_ := "inanç           "
	fmt.Println(strings.TrimRight(name_, " "), "KK")
}
