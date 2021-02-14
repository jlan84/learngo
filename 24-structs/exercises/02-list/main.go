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
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: List
//
//  Now, it's time to add an interface to your program using
//  the bufio.Scanner. So the users can list the games, or
//  search for the games by id.
//
//  1. Scan for the input in a loop (use bufio.Scanner)
//
//  2. Print the available commands.
//
//  3. Implement the quit command: Quits from the loop.
//
//  4. Implement the list command: Lists all the games.
//
//
// EXPECTED OUTPUT
//  Please run the solution and try the program with list and
//  quit commands.
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

	fmt.Println("> list : lists all of the games\n> quit: quits")

	for in.Scan() {
		line := strings.ToLower(in.Text())

		switch line {
		case "list":
			fmt.Printf("%-15s%-15s%-15s%s\n%s\n", "id", "name", "price", "genre",
				strings.Repeat("-", 50))
			for _, v := range gameStore {
				fmt.Printf("%-15d%-15q$%-15d(%s)\n", v.id, v.name, v.price, v.genre)
			}
		case "quit":
			fmt.Println("Bye!")
			return

		}
	}

}
