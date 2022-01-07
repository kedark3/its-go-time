package main

import "fmt"

//     1- Move the structs there

type Item struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type game struct {
	Item
	Genre string `json:"genre"`
}

//     2- Add a func that creates and returns a game.
//
//        Name  : newGame
//        Inputs: id, price, name and genre
//        Output: game
func newGame(id, price int, name, genre string) game {
	return game{
		Item{Id: id, Name: name, Price: price},
		genre,
	}
}

//
//     3- Add a func that adds a `game` to `[]game` and returns `[]game`:
//
//        Name  : addGame
//        Inputs: []game, game
//        Output: []game
//

func addGame(games []game, g game) []game {
	return append(games, g)
}

//     4- Add a func that uses newGame and addGame funcs to load up the game
//        store:
//
//        Name  : load
//        Inputs: none
//        Output: []game
//

func load() (games []game) {
	games = addGame(games, newGame(1, 50, "god of war", "action adventure"))
	games = addGame(games, newGame(2, 40, "x-com 2", "strategy"))
	games = addGame(games, newGame(3, 20, "minecraft", "sandbox"))
	return
}

//     5- Add a func that indexes games by ID:
//
//        Name  : indexByID
//        Inputs: []game
//        Output: map[int]game
//

func indexById(games []game) (gById map[int]game) {
	gById = make(map[int]game)
	for _, g := range games {
		gById[g.Id] = g
	}
	return
}

//     6- Add a func that prints the given game:
//
//        Name  : printGame
//        Inputs: game

func printGame(game game) {
	fmt.Printf("#%d: %-15q %-20s $%d\n", game.Id, game.Name, "("+game.Genre+")", game.Price)
}
