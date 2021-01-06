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
// EXERCISE: Currency Converter
//
//   In this exercise, you're going to display currency exchange ratios
//   against USD.
//
//   1. Declare a few constants with iota. They're going to be the keys
//      of the array.
//
//   2. Create an array that contains the conversion ratios.
//
//      You should use keyed elements and the contants you've declared before.
//
//   3. Get the USD amount to be converted from the command line.
//
//   4. Handle the error cases for missing or invalid input.
//
//   5. Print the exchange ratios.
//
// EXPECTED OUTPUT
//   go run main.go
//     Please provide the amount to be converted.
//
//   go run main.go invalid
//     Invalid amount. It should be a number.
//
//   go run main.go 10.5
//     10.50 USD is 9.24 EUR
//     10.50 USD is 8.19 GBP
//     10.50 USD is 1186.71 JPY
//
//   go run main.go 1
//     1.00 USD is 0.88 EUR
//     1.00 USD is 0.78 GBP
//     1.00 USD is 113.02 JPY
// ---------------------------------------------------------

func main() {
	const (
		EUR = iota
		GBP
		JPY
	)
	conversions := [3]float64{
		EUR: 0.88,
		GBP: 0.78,
		JPY: 113.02,
	}
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Enter an amount of USD to be converted")
		return
	}
	USD, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Printf("%q is not a valid number\n", os.Args[1])
		return
	}
	fmt.Printf("%g USD is %.2f EUR\n", USD, USD*conversions[EUR])
	fmt.Printf("%g USD is %.2f GBP\n", USD, USD*conversions[GBP])
	fmt.Printf("%g USD is %.2f JPY\n", USD, USD*conversions[JPY])
}
