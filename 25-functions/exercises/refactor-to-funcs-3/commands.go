// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func runCmd(input string, games []game, byID map[int]game) bool {
	fmt.Println()

	cmd := strings.Fields(input)
	if len(cmd) == 0 {
		return true
	}

	switch cmd[0] {
	case "quit":
		return cmdQuit()

	case "list":
		return cmdList(games)

	case "id":
		return cmdByID(cmd, games, byID)
	case "save":
		return saveCmd(games)
	}
	return true
}

func cmdQuit() bool {
	fmt.Println("bye!")
	return false
}

func cmdList(games []game) bool {
	for _, g := range games {
		printGame(g)
	}
	return true
}

func cmdByID(cmd []string, games []game, byID map[int]game) bool {
	if len(cmd) != 2 {
		fmt.Println("wrong id")
		return true
	}

	id, err := strconv.Atoi(cmd[1])
	if err != nil {
		fmt.Println("wrong id")
		return true
	}

	g, ok := byID[id]
	if !ok {
		fmt.Println("sorry. i don't have the game")
		return true
	}

	printGame(g)

	return true
}

func saveCmd(games []game) bool {
	var newGames []export
	for _, g := range games {
		newGames = append(newGames, export{g.id, g.name, g.price, g.genre})
	}
	fmt.Println("Saving games and quitting")
	file, err := json.MarshalIndent(newGames, "", "\t")
	if err != nil {
		fmt.Println(err)
		return false
	}
	err = ioutil.WriteFile("games.json", file, 0644)
	return false
}
