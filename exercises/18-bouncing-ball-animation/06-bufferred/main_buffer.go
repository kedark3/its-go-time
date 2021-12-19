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
	buf := make([]rune, 0, w*h)
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
					buf = append(buf, ball)
				} else {
					buf = append(buf, ' ')
				}
			}
			buf = append(buf, '\n')
		}
		fmt.Println(string(buf))
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
