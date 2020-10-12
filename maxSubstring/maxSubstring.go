package main

import (
	"fmt"
)

// Approach 0: run multiple nested loops to find for all i and j indices the given string has all unique characters.
//			   this could run into O(n^3)
// Approach 1: In this approach, we use HashSet where each character we see, we add to hashset, and increment j, if for
//             some character j, it already exists in hashset, we remove character at i from set and increment i
//             variable ans is always max of current ans and difference between (j - i) - caveat take the j - i after incrementing j else add +1 to ans
// Approach 3: Another approach is to skip by n number of characters if char at j is seen, increment i to max(map.get(s.charAt(j)), i) assming values for each char
//             in map are indices of the characters
func main() {
	v2 := true
	ans := 0
	set := make(map[string]bool)

	longString := "aabcdac"
	n := len(longString)

	for i, j := 0, 0; i < n && j < n; {
		if _, ok := set[string(longString[j])]; !ok {
			set[string(longString[j])] = v2
			j++
			ans = max(ans, j-i)
			fmt.Println(longString[i:j])
		} else {
			delete(set, string(longString[i]))
			i++
		}
	}
	fmt.Println(ans)

}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
