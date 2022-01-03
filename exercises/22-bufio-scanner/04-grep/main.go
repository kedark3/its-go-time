package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Grep Clone
//
//  Create a grep clone. grep is a command-line utility for
//  searching plain-text data for lines that match a specific
//  pattern.
//
//  1. Feed the shakespeare.txt to your program.
//
//  2. Accept a command-line argument for the pattern
//
//  3. Only print the lines that contains that pattern
//
//  4. If no pattern is provided, print all the lines
//
//
// EXPECTED OUTPUT
//
//  go run main.go come < shakespeare.txt
//
//    come night come romeo come thou day in night
//    come gentle night come loving black-browed night
//
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Printf("\nPlease pass 1 pattern for grepping. \n")
		return
	}
	pattern := args[0]
	re := regexp.MustCompile(pattern)
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		// With this call it only matches words
		if strings.Contains(line, pattern) {
			fmt.Println(line)
		}
		// With this it matches even with Regex patterns much line real grep
		if match := re.FindStringIndex(line); match != nil {
			fmt.Println(line)
		}
	}

	if err := in.Err(); err != nil {
		fmt.Printf("> Err: %v", err)
	}

}
