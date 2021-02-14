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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Decode
//
//  At the beginning of the file:
//
//  1. Load the initial data to the game store from json.
//     (see the data constant below)
//
//  2. Load the decoded values into the usual `game` values (to the games slice as well).
//
//     So the rest of the program can work intact.
//
//
// HINT
//
//  Move the jsonGame type to the top and reuse it both when
//  loading the initial data, and in the "save" command.
//
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
// ---------------------------------------------------------

const data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

type saved struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Price int    `json:"price"`
}

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

func main() {
	// use your solution from the previous exercise
	in := bufio.NewScanner(os.Stdin)

	// gameStore := []game{
	// 	{item: item{id: 1, name: "god of war", price: 50},
	// 		genre: "action adventure"},
	// 	{item: item{id: 2, name: "x-com 2", price: 30},
	// 		genre: "strategy"},
	// 	{item: item{id: 3, name: "minecraft", price: 20},
	// 		genre: "sandbox"},
	// }

	var (
		savedGames []saved
		gameStore  []game
	)
	dataBytes := []byte(data)
	err := json.Unmarshal(dataBytes, &savedGames)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, g := range savedGames {
		gameStore = append(gameStore, game{item{g.ID, g.Name, g.Price}, g.Genre})
	}

	fmt.Println(`> list : lists all of the games
> id <number>: returns the specific game with the id
> save: saves entered games to file 
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
		case "save":
			var encodable []saved

			for _, g := range gameStore {
				encodable = append(encodable,
					saved{g.id, g.name, g.genre, g.price})
			}
			file, _ := json.MarshalIndent(encodable, "", "\t")
			_ = ioutil.WriteFile("games.json", file, 0644)
			return
		case "quit":
			fmt.Println("Bye!")
			return

		}
	}

}
