package main

import "fmt"

// THIS EXERCISE IS OPTIONAL

// ---------------------------------------------------------
// EXERCISE: Print hexes
//
//  1. Print 0 to 9 in hexadecimal
//
//  2. Print 10 to 15 in hexadecimal
//
//  3. Print 17 in hexadecimal
//
//  4. Print 25 in hexadecimal
//
//  5. Print 50 in hexadecimal
//
//  6. Print 100 in hexadecimal
//
// EXPECTED OUTPUT
//  0 1 2 3 4 5 6 7 8 9
//  10 11 12 13 14 15
//  17
//  25
//  50
//  100
//
// NOTES
//  https://stackoverflow.com/questions/910309/how-to-turn-hexadecimal-into-decimal-using-brain
//
// https://simple.wikipedia.org/wiki/Hexadecimal_numeral_system
//
// ---------------------------------------------------------

func main() {
	// EXAMPLES:

	// I'm going to print 10 in hexadecimal
	// fmt.Println(0xa)

	// I'm going to print 16 in hexadecimal
	// 0x10
	//   ^^-----  1 * 0 = 0
	//   |
	//   +------ 16 * 1 = 16
	//                  = 16
	// fmt.Println(0x10)

	// I'm going to print 150 in hexadecimal
	// 0x96
	//   ^^-----  1 * 6 = 6
	//   |
	//   +------ 16 * 9 = 144
	//                  = 150
	// fmt.Println(0x96)

	// COMMENT-OUT ALL THE CODE ABOVE, THEN,
	// ADD YOUR OWN SOLUTIONS BELOW
	fmt.Println(0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9)
	fmt.Println(0xa, 0xb, 0xc, 0xd, 0xe, 0xf)
	fmt.Println(0x11)
	fmt.Println(0x19)
	fmt.Println(0x32)
	fmt.Println(0x64)

}
