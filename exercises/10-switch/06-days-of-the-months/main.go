package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Days in a Month
//
//  Refactor the previous exercise from the if statement
//  section.
//
//  "Print the number of days in a given month."
//
//  Use a switch statement instead of if statements.
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}

	year := time.Now().Year()
	leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	days, month := 28, os.Args[1]

	switch m := strings.ToLower(month); {
	case m == "april" ||
		m == "june" ||
		m == "september" ||
		m == "november":
		days = 30
	case m == "january" ||
		m == "march" ||
		m == "may" ||
		m == "july" ||
		m == "august" ||
		m == "october" ||
		m == "december":
		days = 31
	case m == "february":
		if leap {
			days = 29
		}
	default:
		fmt.Printf("%q is not a month.\n", month)
		return
	}

	fmt.Printf("%q has %d days.\n", month, days)
}
