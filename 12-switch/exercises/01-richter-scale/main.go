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
// STORY
//  You're curious about the richter scales. When reporters
//  say: "There's been a 5.5 magnitude earthquake",
//
//  You want to know what that means.
//
//  So, you've decided to write a program to do that for you.
//
// EXERCISE: Richter Scale
//
//  1. Get the earthquake magnitude from the command-line.
//  2. Display its corresponding description.
//
// ---------------------------------------------------------
// MAGNITUDE                    DESCRIPTION
// ---------------------------------------------------------
// Less than 2.0                micro
// 2.0 and less than 3.0        very minor
// 3.0 and less than 4.0        minor
// 4.0 and less than 5.0        light
// 5.0 and less than 6.0        moderate
// 6.0 and less than 7.0        strong
// 7.0 and less than 8.0        major
// 8.0 and less than 10.0       great
// 10.0 or more                 massive
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me the magnitude of the earthquake.
//
//  go run main.go ABC
//    I couldn't get that, sorry.
//
//  go run main.go 0.5
//    0.50 is micro

// ---------------------------------------------------------

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Give me the magnitude of the earthquake.")
		return
	}

	mag, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("%q is not a magnitude\n", os.Args[1])
		return
	}

	switch {

	case mag >= 10:

		fmt.Printf("%g is massive\n", mag)

	case mag >= 8:
		fmt.Printf("%g is great\n", mag)

	case mag >= 7:
		fmt.Printf("%g is major\n", mag)

	case mag >= 6:
		fmt.Printf("%g is strong\n", mag)

	case mag >= 5:
		fmt.Printf("%g is moderate\n", mag)

	case mag >= 4:
		fmt.Printf("%g is light\n", mag)

	case mag >= 3:
		fmt.Printf("%g is minor\n", mag)

	case mag >= 2:
		fmt.Printf("%g is very minor\n", mag)

	case mag < 2:
		fmt.Printf("%g is micro\n", mag)

	}

}
