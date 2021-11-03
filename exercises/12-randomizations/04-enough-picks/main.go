package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Enough Picks
//
//  If the player's guess number is below 10;
//  then make sure that the computer generates a random
//  number between 0 and 10.
//
//  However, if the guess number is above 10; then the
//  computer should generate the numbers
//  between 0 and the guess number.
//
// WHY?
//  This way the game will be harder.
//
//  Because, in the current version of the game, if
//  the player's guess number is for example 3; then the
//  computer picks a random number between 0 and 3.
//
//  So, currently a player can easily win the game.
//
// EXAMPLE
//  Suppose that the player runs the game like this:
//    go run main.go 9
//
//  Or like this:
//    go run main.go 2
//
//    Then the computer should pick a random number
//    between 0-10.
//
//  Or, if the player runs it like this:
//    go run main.go 15
//
//    Then the computer should pick a random number
//    between 0-15.
// ---------------------------------------------------------

const (
	errNoArgs    = "There are no arguments or too many args."
	errWrongArgs = "%s isn't a valid number."
	win          = "You win!!"
	lost         = "You Lose!!"
	maxTurns     = 20
)

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println(errNoArgs)
		return
	}
	num, err := strconv.Atoi(args[0])
	if err != nil || num < 0 {
		fmt.Printf(errWrongArgs, args[0])
		return
	}
	rand.Seed(time.Now().UnixNano())
	max := 10
	if num > 10 {
		max = num
	}
	for i := 0; i < maxTurns; i++ {
		if rand.Intn(max+1) == num {
			fmt.Println(win)
			return
		}
	}
	fmt.Println(lost)

}
