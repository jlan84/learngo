package main

import (
	"fmt"
	"strconv"
	"strings"
)

func runCmd(choice string, games []game, gameIdx map[int]game) bool {
	cmd := strings.Fields(choice)

	if len(cmd) == 0 {
		return true
	}

	switch cmd[0] {
	case "quit":
		return quitCmd()
	case "list":
		return listCmd(choice, games)
	case "id":
		return idCmd(choice, gameIdx)
	}
	return true
}

func quitCmd() bool {
	fmt.Println("Thanks for playing")
	return false
}

func idCmd(choice string, gameIdxs map[int]game) bool {

	cmd := strings.Fields(choice)
	if len(cmd) != 2 {
		fmt.Println("Please print the id #")
		return false
	}

	id, err := strconv.Atoi(cmd[1])

	if err != nil {
		fmt.Println("Sorry we do not have that game")
		return false
	}

	g, exists := gameIdxs[id]

	if !exists {
		fmt.Println("We do not have that game")
		return false
	}
	fmt.Printf("%-20s%-20s%-20s%s\n%s\n", "ID", "Name", "Price", "Genre",
		strings.Repeat("-", 50))
	fmt.Printf("%-20d%-20s%-20d%s\n", g.id, g.name, g.price, g.genre)

	return true
}

func listCmd(choice string, games []game) bool {
	fmt.Printf("%-20s%-20s%-20s%s\n%s\n", "ID", "Name", "Price", "Genre",
		strings.Repeat("-", 50))

	for _, g := range games {
		fmt.Printf("%-20d%-20s%-20d%s\n", g.id, g.name, g.price, g.genre)
	}

	return true

}
