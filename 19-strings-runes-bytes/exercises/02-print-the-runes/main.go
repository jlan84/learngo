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
)

// ---------------------------------------------------------
// EXERCISE: Print the runes
//
//  1. Loop over the "console" word and print its runes one by one,
//     in decimals, hexadecimals and binary.
//
//  2. Manually put the runes of the "console" word to a byte slice, one by one.
//
//     As the elements of the byte slice use only the rune literals.
//
//     Print the byte slice.
//
//  3. Repeat the step 2 but this time, as the elements of the byte slice,
//     use only decimal numbers.
//
//  4. Repeat the step 2 but this time, as the elements of the byte slice,
//     use only hexadecimal numbers.
//
//
// EXPECTED OUTPUT
//   Run the solution to see the expected output.
// ---------------------------------------------------------

func main() {
	const word = "console"

	fmt.Printf("%-10s%-10s%-10s%s\n%s\n", "Letter", "Decimal", "Hex", "Binary",
		strings.Repeat("-", 50))

	for _, l := range word {
		fmt.Printf("%-10c%-10d%-10x%b\n", l, l, l, l)
	}

	bslice := string([]byte{'c', 'o', 'n', 's', 'o', 'l', 'e'})
	fmt.Println(bslice)

	bslice = string([]byte{99, 111, 110, 115, 111, 108, 101})
	fmt.Println(bslice)

	bslice = string([]byte{0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65})
	fmt.Println(bslice)
}
