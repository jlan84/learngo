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
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Enough Picks
//
//  If the player's guess number is below 10;
//  then make sure that the computer generates a random
//  number between 0 and 10.
//
//  However, if the guess number is above 10; then the
//  computer should generate the numbers
//  between 0 and the guess number.
//
// WHY?
//  This way the game will be harder.
//
//  Because, in the current version of the game, if
//  the player's guess number is for example 3; then the
//  computer picks a random number between 0 and 3.
//
//  So, currently a player can easily win the game.
//
// EXAMPLE
//  Suppose that the player runs the game like this:
//    go run main.go 9
//
//  Or like this:
//    go run main.go 2
//
//    Then the computer should pick a random number
//    between 0-10.
//
//  Or, if the player runs it like this:
//    go run main.go 15
//
//    Then the computer should pick a random number
//    between 0-15.
// ---------------------------------------------------------
const errMsg = `Please enter one positive integer`

func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println(errMsg)
		return
	}
	var (
		mx      int
		verbose bool
	)
	num, err := strconv.Atoi(args[len(args)-1])

	if num <= 10 {
		mx = 11
	} else {
		mx = num + 1
	}

	if err != nil || num < 0 {
		fmt.Println(errMsg)
		return
	}

	if args[0] == "-v" {
		verbose = true
	}

	for i := 0; i < 5; i++ {
		n := rand.Intn(mx)
		if verbose {
			fmt.Printf("%4d", n)
		}
		if n == num {
			fmt.Println("\nYou win :)")
			return
		}
	}
	fmt.Println("\nYou Lose :(")
}
