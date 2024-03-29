// ---------------------------------------------------------
// EXERCISE: Refactor the game store to funcs - step #2
//
//  Let's continue the refactoring from the previous
//  exercise. This time, you're going to refactor the
//  command handling logic.
//
//
//  Create commands.go file
//
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
//
//     2- Add a func that handles the quit command:
//
//        Name  : cmdQuit
//        Input : none
//        Output: bool
//
//     3- Add a func that handles the list command:
//
//        Name  : cmdList
//        Inputs: []game
//        Output: bool
//
//     4- Add a func that handles the id command:
//
//        Name  : cmdByID
//        Inputs: cmd []string, []game, map[int]game
//        Output: bool
//
//     5- Refactor the runCmd() with the cmdXXX funcs.
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
)

func main() {

	games := load()
	byID := indexById(games)

	fmt.Printf("Inanc's game store has %d games.\n", len(games))

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {

		res := runCmd(in.Text(), games, byID)
		if !res {
			break
		} else {
			fmt.Printf(`
> list   : lists all the games
> id N   : queries a game by id
> save   : export json data and quit
> quit   : quits
`)
		}

	}
}
