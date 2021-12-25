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
		speed     = time.Second / 30

		// initial velocities
		ivx, ivy = 4, 3
	)

	var (
		px, py int        // ball position
		vx, vy = ivx, ivy // velocities

	)

	// you can get the width and height using the screen package easily:
	width, height := screen.Size()

	// create a drawing buffer
	buf := make([]rune, 0, width)

	// clear the screen once
	screen.Clear()
	/*
		Since we need to use single dimension, we don't need board anymore. We just use the
		rune buffer to add `cellBall` at appropriate location horizontally with the help of px
		And for rest of the buffer we append `cellEmpty`
		We rewind buffer to length 0 every time we start over the loop to clear previous ball location
		so we no longer need ppx,ppy.
		We then move top left of the screen, clear the previous ball from screen and
		print `py` number of `\n` chars followed by printing out the `buf` to print
		the ball location.
	*/
	for i := 0; i < maxFrames; i++ {
		// calculate the next ball position
		px += vx
		py += vy

		// when the ball hits a border reverse its direction
		if px <= 0 || px >= width-ivx-1 { // adding -1 prevents some edge cases where it complains index out of range
			vx *= -1
		}
		if py <= 0 || py >= height-ivy {
			vy *= -1
		}
		// rewind the buffer (allow appending from the beginning)
		buf = buf[:0]
		buf = append(buf, []rune(strings.Repeat(string(cellEmpty), len(buf[:px])))...)
		buf = append(buf, []rune(strings.Repeat(string(cellEmpty), len(buf[px+1:cap(buf)])))...)
		buf[px] = cellBall
		// print the buffer
		screen.MoveTopLeft()
		screen.Clear()
		fmt.Print(strings.Repeat("\n", py))
		fmt.Print(string(buf))

		// slow down the animation
		time.Sleep(speed)
	}
}
