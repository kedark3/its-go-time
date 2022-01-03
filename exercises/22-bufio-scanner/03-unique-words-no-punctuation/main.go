package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Unique Words 2
//
//  Use your solution from the previous "Unique Words"
//  exercise.
//
//  Before adding the words to your map, remove the
//  punctuation characters and numbers from them.
//
//
// BE CAREFUL
//
//  Now the shakespeare.txt contains upper and lower
//  case letters too.
//
//
// EXPECTED OUTPUT
//
//  go run main.go < shakespeare.txt
//
//   There are 100 words, 69 of them are unique.
//
// ---------------------------------------------------------

func main() {
	// This is the regular expression pattern you need to use:
	// [^A-Za-z]+
	//
	// Matches to any character but upper case and lower case letters
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	re := regexp.MustCompile("[^A-Za-z]+") // just get upper/lower case chars, no numbers, punctuations, special chars

	words := make(map[string]bool)
	var wcount int
	for in.Scan() {
		wcount++
		words[strings.ToLower(re.ReplaceAllString(in.Text(), ""))] = true

	}

	fmt.Printf("There are %d words, %d of them are unique.\n", wcount, len(words))
	if err := in.Err(); err != nil {
		fmt.Printf("> Err: %v", err)
	}
}
