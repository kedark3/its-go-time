package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Query By Id
//
//  Add a new command: "id". So the users can query the games
//  by id.
//
//  1. Before the loop, index the games by id (use a map).
//
//  2. Add the "id" command.
//     When a user types: id 2
//     It should print only the game with id: 2.
//
//  3. Handle the errors:
//
//     id
//     wrong id
//
//     id HEY
//     wrong id
//
//     id 10
//     sorry. I don't have the game
//
//     id 1
//     #1: "god of war" (action adventure) $50
//
//     id 2
//     #2: "x-com 2" (strategy) $40
//
//
// EXPECTED OUTPUT
//  Please also run the solution and try the program with
//  list, quit, and id commands to see it in action.
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

	// index the games by id
	byID := make(map[int]game)
	for _, g := range games {
		byID[g.id] = g
	}

	in := bufio.NewScanner(os.Stdin)
	fmt.Printf(`
> list   : lists all the games
> id     : find game by id e.g. id 10
> quit   : quits
`)
	for in.Scan() {
		t := in.Text()
		switch {

		case t == "list":
			fmt.Println()
			for _, game := range games {
				printGame(game)
			}
			fmt.Println()
		case t == "quit":
			fmt.Println("Bye!")
			return
		case strings.Contains(t, "id"):
			tokens := strings.Split(t, " ")
			if len(tokens) != 2 {
				fmt.Println("wrong id")
				continue
			} else if id, err := strconv.Atoi(tokens[1]); err != nil {
				fmt.Println("wrong id")
				continue
			} else {
				if _, ok := byID[id]; ok {
					printGame(byID[id])
				} else {
					fmt.Println("sorry. I don't have the game")
				}
			}

		default:
			fmt.Println("Use one of the following commands: 'list', 'quit'")

		}
	}

}

func printGame(game game) {
	fmt.Printf("#%d: %-15q %-20s $%d\n", game.id, game.name, "("+game.genre+")", game.price)
}
