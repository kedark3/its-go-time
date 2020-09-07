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
	fmt.Printf("Value of Pi is: %g\n", math.Pi)

	f1, f2, f3 := 65.1, 23.5, 76.34
	floatSum := f1 + f2 + f3
	fmt.Printf("Float Sum: %.10g\n", floatSum)
}
