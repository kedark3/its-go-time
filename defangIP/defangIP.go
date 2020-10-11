package main

import "fmt"

// Leetcode https://leetcode.com/problems/defanging-an-ip-address/
func defangIPaddr(address string) string {
	// Approach 1: Just use internal function to GoLang
	//return strings.Replace(address, ".", "[.]", -1)
	result := ""
	for _, ch := range address {
		if string(ch) == "." {
			result += "[.]"
		} else {
			result += string(ch)
		}
	}
	return result
}

func main() {
	ip := "1.13.41.111"
	fmt.Println(defangIPaddr(ip))
}
