package main

import (
	"fmt"
	"time"
)

func main() {
	w, h := 50, 10
	board := make([][]bool, h)
	for row := range board {
		board[row] = make([]bool, w)
		fmt.Println(row)
	}

	ball := "⚽"

	var x, y, oldX, oldY int
	var reverseX, reverseY bool
	for i := 0; i < 1e3; i++ {
		board[x][y] = true
		oldX, oldY = x, y
		if x+1 < h && !reverseX {
			x++
		} else {
			x--
		}
		if y+1 < w && !reverseY {
			y++
		} else {
			y--
		}
		if x+1 == h {
			reverseX = true
		} else if x == 0 {
			reverseX = false
		}
		if y+1 == w {
			reverseY = true
		} else if y == 0 {
			reverseY = false
		}
		for i := range board {
			for j := range board[0] {
				if board[i][j] {
					fmt.Print(ball)
				} else {
					fmt.Print("   ")
				}
			}
			fmt.Println()
		}
		board[oldX][oldY] = false
		time.Sleep(time.Second / 32)
		fmt.Println("\033[H\033[2J")
	}

}
