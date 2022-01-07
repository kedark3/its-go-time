//     1- Add a func that runs the given command from the user:
//
//        Name  : runCmd
//        Inputs: input string, []game, map[int]game
//        Output: bool
//
//        This func returns true if it wants the program to
//        continue. When it returns false, the program will
//        terminate. So, all the commands that it calls need
//        to return true or false as well depending on whether
//        they want to cause the program to terminate or not.
package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func runCmd(input string, games []game, gById map[int]game) bool {
	cmd := strings.Fields(input)
	if len(cmd) == 0 {
		return true
	}

	switch cmd[0] {
	case "quit":
		return cmdQuit()

	case "list":
		return cmdList(games)
	case "id":
		return cmdById(cmd, games, gById)
	case "save":
		return cmdSave(games)
	}

	return true
}

//     2- Add a func that handles the quit command:
//
//        Name  : cmdQuit
//        Input : none
//        Output: bool

func cmdQuit() bool {
	fmt.Println("bye!")
	return false
}

//     3- Add a func that handles the list command:
//
//        Name  : cmdList
//        Inputs: []game
//        Output: bool

func cmdList(games []game) bool {
	for _, g := range games {
		printGame(g)
	}
	return true
}

//     4- Add a func that handles the id command:
//
//        Name  : cmdByID
//        Inputs: cmd []string, []game, map[int]game
//        Output: bool
//
//     5- Refactor the runCmd() with the cmdXXX funcs.
func cmdById(cmd []string, games []game, gById map[int]game) bool {
	if len(cmd) != 2 {
		fmt.Println("wrong id")
		return true
	}

	id, err := strconv.Atoi(cmd[1])
	if err != nil {
		fmt.Println("wrong id")
		return true
	}

	g, ok := gById[id]
	if !ok {
		fmt.Println("sorry. I don't have the game")
		return true
	}
	printGame(g)
	return true
}

//  6- Create a new command in a func that encodes the
//     game store data to json and terminates the program.
//     Lastly, add it to runCmd func.
//
//     For more information, see: "Encode" exercise from
//     the structs section.
//
//        Name  : cmdSave
//        Inputs: []game
//        Output: bool

func cmdSave(games []game) bool {
	j, err := json.MarshalIndent(games, "", "\t")
	if err != nil {
		fmt.Println(err)
		return true
	}
	fmt.Println(string(j))
	return cmdQuit()
}

//  7- Refactor the load() to load the games from the
//     `data` constant (it's in the games.go as well).
//
//     For more information, see: "Decode" exercise from
//     the structs section.
//
