package main

import (
	"bufio"
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: List
//
//  Now, it's time to add an interface to your program using
//  the bufio.Scanner. So the users can list the games, or
//  search for the games by id.
//
//  1. Scan for the input in a loop (use bufio.Scanner)
//
//  2. Print the available commands.
//
//  3. Implement the quit command: Quits from the loop.
//
//  4. Implement the list command: Lists all the games.
//
//
// EXPECTED OUTPUT
//  Please run the solution and try the program with list and
//  quit commands.
// ---------------------------------------------------------

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

func main() {
	// use your solution from the previous exercise
	games := []game{
		{
			item{id: 1, name: "god of war", price: 50}, // <--  we can use this form of declaration
			"action adventure",
		},
		{
			item{2, "x-com 2", 30}, // <-- or This
			"strategy",
		},
		{ // <--  or This
			item:  item{3, "minecraft", 20},
			genre: "sandbox",
		},
	}

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	fmt.Printf(`
> list   : lists all the games
> quit   : quits
`)
	for in.Scan() {

		switch in.Text() {

		case "list":
			fmt.Println()
			for _, game := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n", game.id, game.name, "("+game.genre+")", game.price)
			}
		case "quit":
			return
		default:
			fmt.Println("Use one of the following commands: 'list', 'quit'")

		}
	}

}
