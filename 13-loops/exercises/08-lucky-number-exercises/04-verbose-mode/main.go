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
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
// ---------------------------------------------------------

const errMsg = `Please enter -v if you would like to print the numbers generated
on the screen and one positive integer`

func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(errMsg)
		return
	}

	switch len(args) {
	case 1:
		num, err := strconv.Atoi(args[0])
		if err != nil || num < 0 {
			fmt.Println(errMsg)
			return
		}
		for i := 0; i < 5; i++ {
			n := rand.Intn(num + 1)
			if n == num {
				fmt.Println("You Win!!")
				return
			}
		}
	case 2:
		verb := args[0]
		num, err := strconv.Atoi(args[1])
		if err != nil || num < 0 || verb != "-v" {
			fmt.Println(errMsg)
			return
		}
		for i := 0; i < 5; i++ {
			n := rand.Intn(num + 1)
			fmt.Printf("%2d", n)
			if n == num {
				fmt.Println("\nYou Win!!")
				return
			}
		}
	}
	fmt.Println("\nYou lose :(")
}
