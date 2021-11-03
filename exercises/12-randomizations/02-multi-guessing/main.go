package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

const (
	errWrongInput = `You need to pass 2 numbers. e.g. go run main.go <num1> <num2>`
	errNum        = `You did not enter valid numbers.`
	errNegNum     = `You entered negative numbers.`
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println(errWrongInput)
		return
	}

	num1, err1 := strconv.Atoi(args[0])
	num2, err2 := strconv.Atoi(args[1])
	if err1 != nil || err2 != nil {
		fmt.Println(errNum)
		return
	}

	if num1 < 0 || num2 < 0 {
		fmt.Println(errNegNum)
		return
	}
	rand.Seed(time.Now().UnixNano())
	guess := 0
	if num1 < num2 {
		guess = num2 + 1
	} else {
		guess = num1 + 1
	}

	for i := 0; i < 5; i++ {
		n := rand.Intn(guess)
		fmt.Println(n)
		if n != num1 && n != num2 {
			continue
		}
		if n == num1 || n == num2 {
			fmt.Printf("%d matched, you Win!!", n)
			return
		}
	}
	fmt.Println("You LOSE!!!")
}
