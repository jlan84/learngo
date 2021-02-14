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
	"strings"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Rune Manipulator
//
//  Please read the comments inside the following code.
//
// EXPECTED OUTPUT
//  Please run the solution.
// ---------------------------------------------------------

func main() {
	words := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}
	fmt.Printf("%-15s%-15s%-15s\n%s\n", "Word", "byte length", "rune length",
		strings.Repeat("-", 50))

	for _, word := range words {
		fmt.Printf("%-15s%-15d%-15d\n", word, len(word), utf8.RuneCountInString(word))
		for _, l := range word {
			fmt.Printf("bytes: % x ", l)
			fmt.Printf("rune: %x ", l)
		}
	}

	// Print the byte and rune length of the strings
	// Hint: Use len and utf8.RuneCountInString

	// Print the bytes of the strings in hexadecimal
	// Hint: Use % x verb

	// Print the runes of the strings in hexadecimal
	// Hint: Use % x verb

	// Print the runes of the strings as rune literals
	// Hint: Use for range

	// Print the first rune and its byte size of the strings
	// Hint: Use utf8.DecodeRuneInString

	// Print the last rune of the strings
	// Hint: Use utf8.DecodeLastRuneInString

	// Slice and print the first two runes of the strings

	// Slice and print the last two runes of the strings

	// Convert the string to []rune
	// Print the first and last two runes
}
