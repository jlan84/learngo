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

const errMsg = `Please enter a positive integer for a guess and a difficulty
level from 1 - 10, 10 being the easiest`

// ---------------------------------------------------------
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
// ---------------------------------------------------------

func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println(errMsg)
		return
	}
	guess, err1 := strconv.Atoi(args[0])
	level, err2 := strconv.Atoi(args[1])
	if err1 != nil || err2 != nil {
		fmt.Println(errMsg)
		return
	}

	fmt.Print("Numbers:")
	for i := 0; i < level; i++ {
		n := rand.Intn(guess + 1)
		fmt.Printf("%4d ", n)
		if guess != n {
			continue
		}
		if guess == n {
			switch msg := rand.Intn(3); msg {
			case 0:
				fmt.Println("\nYou won the game!")
			case 1:
				fmt.Println("\nYou complete me!")
			case 2:
				fmt.Println("\nIf I wasn't a computer I'd give you a huuuuuge....")
			}
			return
		}
	}
	fmt.Println("\nYou lost :(")
}
