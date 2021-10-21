package main

import (
	"fmt"
	"math/big"
	"sync"
)

func factorial(number *big.Int) *big.Int {
	one := big.NewInt(1)
	var answer = big.NewInt(1)
	for i := 1; i <= int(number); i++ {
		answer *= uint(i)
	}
	return one.Mul(number, factorial(one.Sub(number, one)))
}

var wg = sync.WaitGroup{}

func main() {
	one := big.NewInt(1)
	num := big.NewInt(1)
	for i := 0; i < 900; i++ {
		wg.Add(1)
		fmt.Println(num, factorial(num.Add(num, one)))
	}
}
