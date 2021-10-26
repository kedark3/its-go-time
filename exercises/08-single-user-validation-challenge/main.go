package main

import (
	"fmt"
	"os"
)

// Since this program is expected to validate a fixed user and password always, make that constant
// It will be easier to update later in case we want to validate a different user
const (
	username = "jack"
	password = "1888"
)

func main() {
	var (
		args = os.Args
		l    = len(args)
	)

	if l != 3 { // either no username or password provided or only one argument is passed
		fmt.Println("Usage: go run main.go [username] [password]")
		return
	}

	if args[1] != username {
		fmt.Printf("Access denied for %q.\n", args[1]) // if we want we can add errUser, errPasswd and accessOK Const
	} else if args[1] == username && args[2] != password {
		fmt.Printf("Invalid password for %q.\n", args[1])
	} else if args[1] == username && args[2] == password {
		fmt.Printf("Access granted to %q.\n", username)
	}

}
