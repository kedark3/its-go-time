package main

import (
	"fmt"
	"math"
	"runtime"
)

// Interestingly data types come after the variable name, vs like in Java/C/C++
func add(x int, y int) int {
	return x + y
}

// We can also omit data type if two or more vars share same type

func add3(x, y, z int) int {
	return x + y + z
}

// Function returning multiple results
func swap(x, y string) (string, string) {
	return y, x
}

/* Split any given number into two numbers that add up to given number
   Named Return Values: Here we have naked returns which can be used in small functions
   (large functions may suffer readability with that)
   Go's return values may be named.
   If so, they are treated as variables defined at the top of the function.
*/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func fib(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return i + fib(i-1)
}

func factorial(i int) int {
	if i == 0 || i == 1 {
		return 1
	}
	return i * factorial(i-1)
}

// Sqrt : returns Sqrt of given number
func Sqrt(x float64) float64 {
	z := 3.0
	//	for i := 0; i < 10; i++ {
	//	fmt.Println(z)
	//z -= (z*z - x) / (2 * z)
	//}
	newZ := 0.0
	for {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-newZ) < 1e-15 {
			break
		}
		newZ = z
	}
	return z
}

func main() {
	fmt.Println(add(43, 12))
	fmt.Println(add3(1, 2, 3))

	a, b := swap("hello", "world")

	fmt.Println(a, b)

	fmt.Println(split(33))
	stars := "*"

	starsAddUp := stars
	for i := 0; i < 5; i++ {
		fmt.Println(starsAddUp)
		starsAddUp += "*"
	}

	starsSubtract := "*****"

	for i := 5; i > 0; i-- {
		fmt.Println(starsSubtract[:i])
	}
	sum := 1
	// This acts as while loop as well
	for sum < 100 {
		fmt.Println(sum)
		sum += sum
	}
	fmt.Println("Fibonacci of 5 is:", fib(5))
	fmt.Println("Fibonacci of 5 is:", fib(10))

	fmt.Println("Factorial of 5 is:", factorial(5))
	fmt.Println("Factorial of 10 is:", factorial(10))
	fmt.Println(Sqrt(9))

	/*
		Infinite loop would be:
		for{

		}
	*/

	// Switch Case - Good ol' C language - brings back my memories
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("My Penguine")
	default:
		fmt.Printf("%s.\n", os)
	}
	// Implement nostalgic calculator
	// for {
	// 	fmt.

	// }
}
