package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Time in Unix Epoch, seeding into Random", time.Now().Unix())
	rand.Seed(time.Now().Unix())
	fmt.Println("My favorite (random) number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Printf("Value of Pi is: %g", math.Pi)
}
