// ---------------------------------------------------------
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
// ---------------------------------------------------------

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	errUsage   = `Please provide a number. e.g. go run main.go [-v] <number>`
	errZeroNum = `Please enter number greater than 0`
	usage      = `Please enter a number and computer will pick a random number several times, if it matches your number, you win.
You will get bonus if your number is the first number computer guesses.`
	errNum   = `%s is an invalid number.`
	winBonus = `You won special bonus as well.`
	win      = `You won!!!`
	lost     = `you lost :(`
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println(errUsage)
		return
	}
	verbose, userNum, err := false, 0, error(nil)
	if args[0] == "-v" {
		verbose = true
		userNum, err = strconv.Atoi(args[1])
	} else {
		userNum, err = strconv.Atoi(args[0])

	}
	if err != nil {
		fmt.Printf(errNum, args[0])
		return
	}

	if userNum < 1 {
		fmt.Println(errZeroNum)
	}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		n := rand.Intn(userNum + 1)
		if verbose {
			fmt.Printf("%3d", n)
		}
		if n == userNum {
			fmt.Printf("%20s\n", win)
			if i == 0 {
				fmt.Printf("%20s\n", winBonus)
			}
			return
		}
	}
	fmt.Printf("%20s\n", lost)

}
