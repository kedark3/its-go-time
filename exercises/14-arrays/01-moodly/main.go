package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	moods := [...]string{"good", "bad", "happy", "awesome", "sad", "terrible"}
	if len(os.Args[1:]) != 1 {
		fmt.Println("Usage: go run main.go [your name]")
		return
	}
	name := os.Args[1]

	rand.Seed(time.Now().UnixNano())

	fmt.Printf("\n%s feels %s\n", name, moods[rand.Intn(5)])
}
