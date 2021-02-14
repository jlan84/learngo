// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Warm Up
//
//  Starting with this exercise, you'll build a command-line
//  game store.
//
//  1. Declare the following structs:
//
//     + item: id (int), name (string), price (int)
//
//     + game: embed the item, genre (string)
//
//
//  2. Create a game slice using the following data:
//
//     id  name          price    genre
//
//     1   god of war    50       action adventure
//     2   x-com 2       30       strategy
//     3   minecraft     20       sandbox
//
//
//  3. Print all the games.
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
// ---------------------------------------------------------

func main() {
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

	fmt.Printf("%-15s%-15s%-15s%s\n%s\n", "id", "name", "price", "genre",
		strings.Repeat("-", 50))

	for _, v := range gameStore {
		fmt.Printf("%-15d%-15q$%-15d(%s)\n", v.id, v.name, v.price, v.genre)
	}

}
