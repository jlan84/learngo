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
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const errMsg = `Please enter two positive integer numbers and a level of 
difficulty from 1 - 10, 10 being the easiest`

// ---------------------------------------------------------
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

func main() {
	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Println(errMsg)
		return
	}

	num1, err1 := strconv.Atoi(args[0])
	num2, err2 := strconv.Atoi(args[1])
	lvl, err3 := strconv.Atoi(args[2])
	mx := int(math.Max(float64(num1), float64(num2)))

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(errMsg)
		return
	}
	fmt.Print("Numbers: ")
	for i := 0; i < lvl; i++ {
		n := rand.Intn(mx + 1)
		fmt.Printf("%4d", n)
		if num1 != n && num2 != n {
			continue
		} else {
			fmt.Println("\nYou win!")
			return
		}

	}
	fmt.Println("\nYou lose :(")
}
