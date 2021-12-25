package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/inancgumus/screen"
)

// ---------------------------------------------------------
// EXERCISE: Single Dimensional
//
//  In this exercise you will understand why I use
//  a multi-dimensional board slice instead of a
//  single-dimensional one.
//
//  1. Remove this:
//     board := make([][]bool, width)
//
//  2. Use this:
//     board := make([]bool, width*height)
//
//  3. Adjust the rest of the operations in the code to work
//     with this single-dimensional slice.
//
//     You'll see how hard it becomes to work with it.
//
// ---------------------------------------------------------

func main() {
	const (
		cellEmpty = ' '
		cellBall  = '⚾'

		maxFrames = 1200
		speed     = time.Second / 10

		// initial velocities
		ivx, ivy = 8, 4
	)

	var (
		px, py int        // ball position
		ppx    int        // previous ball position
		vx, vy = ivx, ivy // velocities

		cell rune // current cell (for caching)
	)

	// you can get the width and height using the screen package easily:
	width, height := screen.Size()

	// create the board -- here for this challenge we remove the height dimension of the board
	// just using width for the row and adjusting the px and ppx such that we always remove the
	// ball from row[ppx] and put it on row[px] using bool. Then we print py number of "\n" chars
	// and then print the row
	row := make([]bool, width)

	// create a drawing buffer
	buf := make([]rune, 0, width)

	// clear the screen once
	screen.Clear()

	for i := 0; i < maxFrames; i++ {
		// calculate the next ball position
		px += vx
		py += vy

		// when the ball hits a border reverse its direction
		if px <= 0 || px >= width-ivx {
			vx *= -1
		}
		if py <= 0 || py >= height-ivy {
			vy *= -1
		}

		// remove the previous ball and put the new ball
		row[px], row[ppx] = true, false

		// save the previous positions
		ppx = px

		// rewind the buffer (allow appending from the beginning)
		buf = buf[:0]

		// draw the board into the buffer
		for x := range row {
			cell = cellEmpty
			if row[x] {
				cell = cellBall
			}
			buf = append(buf, cell)
		}

		// print the buffer
		screen.MoveTopLeft()
		screen.Clear()
		fmt.Print(strings.Repeat("\n", py))
		fmt.Print(string(buf))

		// slow down the animation
		time.Sleep(speed)
	}
}
