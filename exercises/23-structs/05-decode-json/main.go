package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Decode
//
//  At the beginning of the file:
//
//  1. Load the initial data to the game store from json.
//     (see the data constant below)
//
//  2. Load the decoded values into the usual `game` values (to the games slice as well).
//
//     So the rest of the program can work intact.
//
//
// HINT
//
//  Move the jsonGame type to the top and reuse it both when
//  loading the initial data, and in the "save" command.
//
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
// ---------------------------------------------------------

const data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

type Item struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Price int    `json:"price,omitempty"`
}

type game struct {
	Item
	Genre string `json:"genre,omitempty"`
}

func main() {
	// use your solution from the previous exercise
	// games := []game{
	// 	{
	// 		Item{Id: 1, Name: "god of war", Price: 50}, // <--  we can use this form of declaration
	// 		"action adventure",
	// 	},
	// 	{
	// 		Item{2, "x-com 2", 30}, // <-- or This
	// 		"strategy",
	// 	},
	// 	{ // <--  or This
	// 		Item:  Item{3, "minecraft", 20},
	// 		Genre: "sandbox",
	// 	},
	// }
	var games []game
	err := json.Unmarshal([]byte(data), &games)

	if err != nil {
		fmt.Println(err)
		return
	}

	// index the games by id
	byID := make(map[int]game)
	for _, g := range games {
		byID[g.Id] = g
	}

	in := bufio.NewScanner(os.Stdin)
	fmt.Printf(`
> list   : lists all the games
> id N   : find game by id e.g. id 10
> save   : exports the data to json and quits
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
		case t == "save":
			j, err := json.MarshalIndent(games, "", "\t")
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(string(j))

		default:
			fmt.Println("Use one of the following commands: 'list', 'quit'")

		}
	}

}

func printGame(game game) {
	fmt.Printf("#%d: %-15q %-20s $%d\n", game.Id, game.Name, "("+game.Genre+")", game.Price)
}
