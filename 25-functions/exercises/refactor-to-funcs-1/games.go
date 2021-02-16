package main

import "fmt"

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

func newGame(id, price int, name, genre string) (newGame game) {

	newGame = game{item: item{id, name, price}, genre: genre}

	return
}

func addGame(new game, games []game) []game {
	return append(games, new)
}

func load(id, price int, name, genre string, currentGames []game) (g []game) {
	g = addGame(newGame(id, price, name, genre), currentGames)
	return
}

func indexByID(games []game) (dic map[int]game) {
	dic = make(map[int]game)
	for _, g := range games {
		dic[g.id] = g
	}
	return
}

func printGames(g game) {

	fmt.Printf("#%d: %-15q %-20s $%d\n",
		g.id, g.name, "("+g.genre+")", g.price)
}
