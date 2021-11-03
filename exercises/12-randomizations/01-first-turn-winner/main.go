package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------

const (
	errUsage = `Please provide a number. e.g. go run main.go <number>`
	usage    = `Please enter a number and computer will pick a random number several times, if it matches your number, you win.
You will get bonus if your number is the first number computer guesses.`
	errNum   = `%s is an invalid number.`
	winBonus = `You won special bonus as well.`
	win      = `You won!!!`
	lost     = `you lost :(`
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println(errUsage)
		return
	}
	userNum, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf(errNum, args[0])
		return
	}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		if rand.Intn(userNum+1) == userNum {
			fmt.Println(win)
			if i == 0 {
				fmt.Println(winBonus)
			}
			return
		}
	}
	fmt.Println(lost)

}
