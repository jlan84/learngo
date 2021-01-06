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
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Crunch the primes
//
//  1. Get numbers from the command-line.
//  2. `for range` over them.
//  4. Convert each one to an int.
//  5. If one of the numbers isn't an int, just skip it.
//  6. Print the ones that are only the prime numbers.
//
// RESTRICTION
//  The user can run the program with any number of
//  arguments.
//
// HINT
//  Find whether a number is prime using this algorithm:
//  https://stackoverflow.com/a/1801446/115363
//
// EXPECTED OUTPUT
//  go run main.go 2 4 6 7 a 9 c 11 x 12 13
//    2 7 11 13
//
//  go run main.go 1 2 3 5 7 A B C
//    2 3 5 7
// ---------------------------------------------------------
const errMsg = `Please enter valid integers`

func main() {

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println(errMsg)
		return
	}
argLoop:
	for _, arg := range args {
		n, err := strconv.Atoi(arg)
		if err != nil {
			continue
		}
		switch {
		case n == 2 || n == 3:
			fmt.Printf("%d ", n)
		case n < 1 || n%2 == 0 || n%3 == 0:
			continue
		}

		for i, w := 5, 2; i*i <= n; {
			if n%i == 0 {
				continue argLoop
			}
			i += w
			w = 6 - 2
		}
		fmt.Printf("%d ", n)

	}
	fmt.Print("\n")
}
