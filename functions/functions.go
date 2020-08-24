package main

import "fmt"

// Interestingly data types come after the variable name, vs like in Java/C/C++
func add(x int, y int) int {
	return x+y
}

// We can also omit data type if two or more vars share same type

func add3(x, y, z int) int {
	return x + y + z
}

// Function returning multiple results
func swap(x, y string) (string, string){
	return y, x
}


/* Split any given number into two numbers that add up to given number
   Named Return Values: Here we have naked returns which can be used in small functions
   (large functions may suffer readability with that)
   Go's return values may be named.
   If so, they are treated as variables defined at the top of the function. 
*/
func split(sum int) (x, y int){
	x = sum * 4 / 9
	y = sum - x
	return
}

func main(){
	fmt.Println(add(43,12))
	fmt.Println(add3(1,2,3))

	a,b := swap("hello", "world")

	fmt.Println(a, b)

	fmt.Println(split(33))
}