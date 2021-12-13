package main

import (
	"fmt"
	"strings"

	s "github.com/inancgumus/prettyslice"
)

// ---------------------------------------------------------
// EXERCISE: Add a newline after each sentence
//
//  You have a slice that contains Beatles' awesome song:
//  Yesterday. You want to add newlines after each sentence.
//
//  So, create a new slice and copy every words into it. Lastly,
//  after each sentence, add a newline character ('\n').
//
//
// ORIGINAL SLICE:
//
//   [yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday]
//
// EXPECTED SLICE (NEW):
//
//   [yesterday all my troubles seemed so far \n away now it looks as though they are here to stay \n oh I believe in yesterday \n]
//
//
// CURRENT OUTPUT
//
//  yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday
//
// EXPECTED OUTPUT
//
//  yesterday all my troubles seemed so far away
//  now it looks as though they are here to stay
//  oh I believe in yesterday
//
//
// RESTRICTIONS
//
//  + Don't use `append()`, use `copy()` instead.
//
//  + Don't cheat like this:
//
//    fmt.Println(lyric[:8])
//    fmt.Println(lyric[8:18])
//    fmt.Println(lyric[18:23])
//
//  + Create a new slice that contains the sentences
//    with line endings.
//
//
// NOTE
//
// If the program does not work, please update your
// local copy of the prettyslice package:
//
//   go get -u github.com/inancgumus/prettyslice
//
//
// ---------------------------------------------------------

func main() {
	// You need to add a newline after each sentence in another slice.
	// Don't touch the following code.
	lyric := strings.Fields(`yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday`)

	// ===================================
	//
	// ~~~ CHANGE THIS CODE ~~~
	//
	fix := make([]string, len(lyric)+3)
	cutpoints := []int{8, 10, 5}
	var n int
	for i, cp := range cutpoints {
		n += copy(fix[i+n:], lyric[n:n+cp])
		fix[i+n] = "\n"
	}
	// for i, n := 0, 0; n < len(lyric); i++ {
	// 	//
	// 	// copy to `fix` from the `lyric`
	// 	//
	// 	//   destination:
	// 	//     fix[n+i] because we don't want to delete the previous copy.
	// 	//     it moves sentence by sentence, using the cutpoints.
	// 	//
	// 	//   source:
	// 	//     lyric[n:n+cutpoints[i]] because we want copy the next sentence
	// 	//     beginning from the number of the last copied elements to the
	// 	//     n+next cutpoint (the next sentence).
	// 	//
	// 	fmt.Println("copy(fix[", n, "+", i, ":], lyric[", n, ":", n, "+", cutpoints[i], "])")
	// 	n += copy(fix[n+i:], lyric[n:n+cutpoints[i]])

	// 	//
	// 	// add a newline after each sentence.
	// 	//
	// 	// notice that the '\n' position slides as we move over.
	// 	// that's why it's: `n+i`.
	// 	//
	// 	fix[n+i] = "\n"

	// 	// uncomment to see how the fix slice changes.
	// 	// s.Show("fix slice", fix)
	// }

	s.Show("fix slice", fix)

	// print the sentences
	for _, w := range fix {
		fmt.Print(w)
		if w != "\n" {
			fmt.Print(" ")
		}
	}
	//
	// ===================================

	// Currently, it prints every sentence on the same line.
	// Don't touch the following code.
	s.Show("fix slice", fix)

	for _, w := range fix {
		fmt.Print(w)
		if w != "\n" {
			fmt.Print(" ")
		}
	}
}

func init() {
	//
	// YOU DON'T NEED TO TOUCH THIS
	//
	// This initializes some options for the prettyslice package.
	// You can change the options if you want.
	//
	// This code runs before the main function above.
	//
	// s.Colors(false)     // if your editor is light background color then enable this
	//
	s.PrintBacking = true  // prints the backing arrays
	s.MaxPerLine = 5       // prints max 15 elements per line
	s.SpaceCharacter = '⏎' // print this instead of printing a newline (for debugging)
}
