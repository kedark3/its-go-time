// ---------------------------------------------------------
// EXERCISE: Refactor the game store to funcs - step #1
//
//  Remember the game store program from the structs exercises?
//  Now, it's time to refactor it to funcs.
//
//  Create games.go file
//
//     1- Move the structs there
//
//     2- Add a func that creates and returns a game.
//
//        Name  : newGame
//        Inputs: id, price, name and genre
//        Output: game
//
//     3- Add a func that adds a `game` to `[]game` and returns `[]game`:
//
//        Name  : addGame
//        Inputs: []game, game
//        Output: []game
//
//     4- Add a func that uses newGame and addGame funcs to load up the game
//        store:
//
//        Name  : load
//        Inputs: none
//        Output: []game
//
//     5- Add a func that indexes games by ID:
//
//        Name  : indexByID
//        Inputs: []game
//        Output: map[int]game
//
//     6- Add a func that prints the given game:
//
//        Name  : printGame
//        Inputs: game
//
//
//  Go back to main.go and change the existing code with
//  the new funcs that you've created. There are hints
//  inside the code.
//
// ---------------------------------------------------------

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	games := load()
	byID := indexById(games)

	fmt.Printf("Inanc's game store has %d games.\n", len(games))

	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(`
  > list   : lists all the games
  > id N   : queries a game by id
  > quit   : quits
`)

		if !in.Scan() {
			break
		}

		fmt.Println()

		cmd := strings.Fields(in.Text())
		if len(cmd) == 0 {
			continue
		}

		switch cmd[0] {
		case "quit":
			fmt.Println("bye!")
			return

		case "list":
			for _, g := range games {
				printGame(g)

			}

		case "id":
			if len(cmd) != 2 {
				fmt.Println("wrong id")
				continue
			}

			id, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Println("wrong id")
				continue
			}

			g, ok := byID[id]
			if !ok {
				fmt.Println("sorry. I don't have the game")
				continue
			}

			printGame(g)
		}
	}
}
