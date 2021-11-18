package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Slicing by arguments
//
//   We've a []string, you will get arguments from the command-line,
//   then you will print only the elements that are requested.
//
//   1. Print the []string (it's in the code below)
//
//   2. Get the starting and stopping positions from the command-line
//
//   3. Print the []string slice by slicing it using the starting and stopping
//      positions
//
//   4. Handle the error cases (see the expected output below)
//
//   5. Add new elements to the []string slice literal.
//      Your program should work without changing the rest of the code.
//
//   6. Now, play with your program, get a deeper sense of how the slicing
//      works.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    Provide only the [starting] and [stopping] positions
//
//
//  (error because: we expect only two arguments)
//
//  go run main.go 1 2 4
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    Provide only the [starting] and [stopping] positions
//
//
//  (error because: starting index >= 0 && stopping pos <= len(slice) )
//
//  go run main.go -1 5
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    Wrong positions
//
//
//  (error because: starting <= stopping)
//
//  go run main.go 3 2
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    Wrong positions
//
//
//  go run main.go 0
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Normandy Verrikan Nexus Warsaw]
//
//
//  go run main.go 1
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Verrikan Nexus Warsaw]
//
//
//  go run main.go 2
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Nexus Warsaw]
//
//
//  go run main.go 3
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Warsaw]
//
//
//  go run main.go 4
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    []
//
//
//  go run main.go 0 3
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Normandy Verrikan Nexus]
//
//
//  go run main.go 1 3
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Verrikan Nexus]
//
//
//  go run main.go 1 2
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Verrikan]
//
// ---------------------------------------------------------

func main() {
	// uncomment the slice below
	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}
	fmt.Printf("%q\n\n", ships)

	from, to := 0, len(ships)

	switch poss := os.Args[1:]; len(poss) {
	default:
		fallthrough
	case 0:
		fmt.Println("Provide only the [starting] and [stopping] positions")
		return
	case 2:
		to, _ = strconv.Atoi(poss[1])
		fallthrough
	case 1:
		from, _ = strconv.Atoi(poss[0])
	}

	if l := len(ships); from < 0 || from > l || to > l || from > to {
		fmt.Println("Wrong positions")
		return
	}

	fmt.Println(ships[from:to])
}
