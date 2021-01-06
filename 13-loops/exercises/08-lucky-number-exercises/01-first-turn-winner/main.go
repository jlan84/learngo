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

const (
	errMsg = `Please enter an integer for you guess and a difficulty level from
	1 - 10, 10 being the easiest.`
)

// ---------------------------------------------------------
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
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
	if err1 != nil && err2 != nil {
		fmt.Println(errMsg)
		return
	}
	if guess < 0 {
		fmt.Println("Please pick a positive number")
	}
	fmt.Print("Guesses:")
	for i := 0; i < level; i++ {
		n := rand.Intn(guess + 1)
		fmt.Printf("%4d", n)
		if n == guess {
			switch i {
			case 0:
				fmt.Print("\nAwesome! You got it on the first try!")
				fallthrough
			default:
				fmt.Println("\nYou win!")
			}
			return
		}
	}
	fmt.Println("\nYou Lost")
}
