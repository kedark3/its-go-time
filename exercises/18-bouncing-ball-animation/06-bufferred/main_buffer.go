package main

import (
	"fmt"
	"time"

	"golang.org/x/term"
)

func main() {
	w, h := getSize()
	board := make([][]bool, h)
	for row := range board {
		board[row] = make([]bool, w)
		fmt.Println(row)
	}
	ball := '⚽'

	var x, y, oldX, oldY int
	var reverseX, reverseY bool
	// *2 for extra spaces
	// +1 for newlines
	buf := make([]rune, 0, (w*2+1)*h) // making buffer with right size beforehand avoids allocation of the backing array in append calls below
	for i := 0; i < 1e3; i++ {
		buf = buf[:0]
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
					buf = append(buf, ball) // using append here instead of fmt.Println would be less expensive
				} else {
					// same here like above, in exercise 07-noBuffer we use print statements which is less effective,
					//it has no buffer
					buf = append(buf, ' ')
				}
			}
			buf = append(buf, '\n')
		}
		fmt.Println(string(buf)) // here we finally print all at once
		board[oldX][oldY] = false
		time.Sleep(time.Second / 32)
		fmt.Println("\033[H\033[2J")
	}

}

func getSize() (int, int) {
	if !term.IsTerminal(0) {
		println("not in a term")
	}
	width, height, err := term.GetSize(0)
	if err != nil {
		return 0, 0
	}
	println("width:", width, "height:", height)
	time.Sleep(time.Second)
	return width, height
}
