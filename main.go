package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	const (
		width  = 50
		height = 10

		cellEmpty = ' '
		cellBall  = '⚾'

		maxFrames = 1200
		speed     = time.Second / 20
	)

	var (
		px, py int    // ball position
		vx, vy = 1, 1 // velocities

		cell rune // current cell (for caching)

	)

	// create the board
	board := make([][]bool, width)
	for colume := range board {
		board[colume] = make([]bool, height)

	}

	// create a drawing buffer
	buf := make([]rune, 0, width*height)

	// clear the screen once
	screen.Clear()

	for i := 0; i < maxFrames; i++ {

		// calculate the next ball position
		px += vx
		py += vy

		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		// remove the previous ball
		for y := range board[0] {
			for x := range board {
				board[x][y] = false

			}
		}

		// put the new ball
		board[px][py] = true

		// rewind the buffer (allow appending from the beginning)
		buf = buf[:0]

		// draw the board
		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall

				}
				//fmt.Print(string(cell))
				buf = append(buf, cell, ' ')

			}

			//fmt.Print("")
			buf = append(buf, '\n')

		}
		screen.MoveTopLeft()
		fmt.Print(string(buf))

		time.Sleep(speed)

	}
}
