// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Query By Id
//
//  Add a new command: "id". So the users can query the games
//  by id.
//
//  1. Before the loop, index the games by id (use a map).
//
//  2. Add the "id" command.
//     When a user types: id 2
//     It should print only the game with id: 2.
//
//  3. Handle the errors:
//
//     id
//     wrong id
//
//     id HEY
//     wrong id
//
//     id 10
//     sorry. i don't have the game
//
//     id 1
//     #1: "god of war" (action adventure) $50
//
//     id 2
//     #2: "x-com 2" (strategy) $40
//
//
// EXPECTED OUTPUT
//  Please also run the solution and try the program with
//  list, quit, and id commands to see it in action.
// ---------------------------------------------------------

func main() {
	// use your solution from the previous exercise
	in := bufio.NewScanner(os.Stdin)

	type item struct {
		id    int
		name  string
		price int
	}

	type game struct {
		item
		genre string
	}

	gameStore := []game{
		{item: item{id: 1, name: "god of war", price: 50},
			genre: "action adventure"},
		{item: item{id: 2, name: "x-com 2", price: 30},
			genre: "strategy"},
		{item: item{id: 3, name: "minecraft", price: 20},
			genre: "sandbox"},
	}

	fmt.Println(`> list : lists all of the games
> id <number>: returns the specific game with the id 
> quit: quits`)

	for in.Scan() {
		line := strings.ToLower(in.Text())
		words := strings.Fields(line)

		switch words[0] {
		case "list":
			fmt.Printf("%-15s%-15s%-15s%s\n%s\n", "id", "name", "price", "genre",
				strings.Repeat("-", 50))
			for _, v := range gameStore {
				fmt.Printf("%-15d%-15q$%-15d(%s)\n", v.id, v.name, v.price, v.genre)
			}
		case "id":
			id, err := strconv.Atoi(words[1])

			if err != nil || id+1 > len(gameStore) || id < 1 {
				fmt.Printf("%q is not an id", words[1])
				return
			}
			fmt.Printf("%-15s%-15s%-15s%s\n%s\n", "id", "name", "price", "genre",
				strings.Repeat("-", 50))
			fmt.Printf("%-15d%-15q$%-15d(%s)\n", gameStore[id-1].id,
				gameStore[id-1].name, gameStore[id-1].price, gameStore[id-1].genre)
		case "quit":
			fmt.Println("Bye!")
			return

		}
	}
}
