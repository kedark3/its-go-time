package main

import (
	"fmt"
	"time"
)

func main() {

	switch t := time.Now().Hour(); {
	case t >= 18:
		fmt.Println("Good Evening")
	case t >= 12:
		fmt.Println("Good Afternoon")
	case t >= 6:
		fmt.Println("Good Morning")
	default:
		fmt.Println("Good Night")
	}
}
