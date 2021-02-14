package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
	"github.com/mattn/go-runewidth"
)

func main() {
	const (
		cellEmpty = ' '
		cellBall  = '🥎'
		speed     = 10
		maxFrames = 1200
		sleepTime = time.Second / 20
	)
	var (
		cell   rune
		px, py int
		vx, vy int
	)
	width, height := screen.Size()
	ballWidth := runewidth.RuneWidth(cellBall)
	width = width / ballWidth
	height -= 2
	vx = 1
	vy = 1
	grid := make([][]bool, height)
	buffer := make([]rune, width*height)

	for y := range grid {
		grid[y] = make([]bool, width)
	}

	for i := 0; i < maxFrames; i++ {
		screen.Clear()

		px += vx
		py += vy

		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		grid[py][px] = true

		buffer = buffer[:0]
		for y := range grid {
			for x := range grid[0] {
				cell = cellEmpty
				if grid[y][x] {
					cell = cellBall
				}
				buffer = append(buffer, cell, ' ')
			}
			buffer = append(buffer, '\n')
		}

		fmt.Print(string(buffer))
		time.Sleep(sleepTime)
		grid[py][px] = false
		screen.MoveTopLeft()
	}
}
