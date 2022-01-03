package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Log Parser from Stratch
//
//  You've watched the lecture. Now, try to create the same
//  log parser program on your own. Do not look at the lecture,
//  and the existing source code.
//
//
// EXPECTED OUTPUT
//
//  go run main.go < log.txt
//
//   DOMAIN                             VISITS
//   ---------------------------------------------
//   blog.golang.org                        30
//   golang.org                             10
//   learngoprogramming.com                 20
//
//   TOTAL                                  60
//
//
//  go run main.go < log_err_missing.txt
//
//   wrong input: [golang.org] (line #3)
//
//
//  go run main.go < log_err_negative.txt
//
//   wrong input: "-100" (line #3)
//
//
//  go run main.go < log_err_str.txt
//
//   wrong input: "FOUR" (line #3)
//
// ---------------------------------------------------------

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanLines)

	var lines, total int
	visitLog := map[string]int{}
	domains := []string{}

	fmt.Printf("%-30s%10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 50))

	for in.Scan() {
		lines++
		line := strings.Split(in.Text(), " ")
		if len(line) < 2 {
			fmt.Printf("wrong input : %v (line #%d)\n", line, lines)
			continue
		}
		domain := line[0]
		visits, err := strconv.Atoi(line[1])
		if visits < 0 || err != nil {
			fmt.Printf("wrong input : %q (line #%d)\n", line[1], lines)
			continue
		}
		if _, ok := visitLog[domain]; !ok {
			domains = append(domains, domain)
		}
		visitLog[domain] += visits
		total += visits
	}
	sort.Strings(domains)
	// Just note even though we sort domains slice in order to traverse and print map in sorted order
	// `fmt` prints maps in key sorted order starting go 1.12 so following print statement will do the same
	// fmt.Printf("VisitLog %#v\n", visitLog)
	for _, domain := range domains {
		fmt.Printf("%-30s%10d\n", domain, visitLog[domain])
	}

	fmt.Printf("\n%-30s%10d\n", "TOTAL", total)

	if err := in.Err(); err != nil {
		fmt.Printf("> Err: %v\n", err)
	}
}
