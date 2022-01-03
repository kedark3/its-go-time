package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Warm Up
//
//  Starting with this exercise, you'll build a command-line
//  game store.
//
//  1. Declare the following structs:
//
//     + item: id (int), name (string), price (int)
//
//     + game: embed the item, genre (string)
//
//
//  2. Create a game slice using the following data:
//
//     id  name          price    genre
//
//     1   god of war    50       action adventure
//     2   x-com 2       30       strategy
//     3   minecraft     20       sandbox
//
//
//  3. Print all the games.
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
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

	for _, game := range games {
		fmt.Printf("#%d: %-15q %-20s $%d\n", game.id, game.name, "("+game.genre+")", game.price)
	}
}
