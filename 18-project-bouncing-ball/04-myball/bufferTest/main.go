package main

import "fmt"

func main() {
	const (
		height    = 10
		width     = 50
		cellEmpty = ' '
		cellBall  = '🥎'
	)
	var cell rune
	grid := make([][]bool, height)
	// buffer := make
	buffer := make([]rune, width*height)
	for y := range grid {
		grid[y] = make([]bool, width)
	}
	grid[0][0] = true
	for i := 0; i < 3; i++ {
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

	}
	fmt.Print(string(buffer))
}
