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
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Quit
//
//  Create a program that quits when a user types the
//  same word twice.
//
//
// RESTRICTION
//
//  The program should work case insensitive.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//
//   hey
//   HEY
//   TWICE!
//
//  go run main.go
//
//   hey
//   hi
//   HEY
//   TWICE!
// ---------------------------------------------------------

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	wordDic := make(map[string]bool)

	for in.Scan() {
		word := strings.ToLower(in.Text())

		if wordDic[word] {
			fmt.Println("TWICE!")
			return
		}
		wordDic[word] = true
	}
}
