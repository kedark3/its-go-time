package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Dynamic Difficulty
//
//  Current game picks only 5 numbers (5 turns).
//
//  Make sure that the game adjust its own difficulty
//  depending on the guess number.
//
// RESTRICTION
//  Do not make the game to easy. Only adjust the
//  difficulty if the guess is above 10.
//
// EXPECTED OUTPUT
//  Suppose that the player runs the game like this:
//    go run main.go 5
//
//    Then the computer should pick 5 random numbers.
//
//  Or, if the player runs it like this:
//    go run main.go 25
//
//    Then the computer may pick 11 random numbers
//    instead.
//
//  Or, if the player runs it like this:
//    go run main.go 100
//
//    Then the computer may pick 30 random numbers
//    instead.
//
//  As you can see, greater guess number causes the
//  game to increase the game turns, which in turn
//  adjust the game's difficulty dynamically.
//
//  Because, greater guess number makes it harder to win.
//  But, this way, game's difficulty will be dynamic.
//  It will adjust its own difficulty depending on the
//  guess number.
// ---------------------------------------------------------

const (
	errNoArgs    = "There are no arguments or too many args."
	errWrongArgs = "%s isn't a valid number."
	win          = "You win!!"
	lost         = "You Lose!!"
)

func main() {

	maxTurns := 5
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
		maxTurns = rand.Intn(max / 2)
		fmt.Println(maxTurns)
	}

	for i := 0; i < maxTurns; i++ {
		if rand.Intn(max+1) == num {
			fmt.Println(win)
			return
		}
	}
	fmt.Println(lost)

}
