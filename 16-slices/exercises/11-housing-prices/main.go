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
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Housing Prices
//
//  We have received housing prices. Your task is to load the data into
//  appropriately typed slices then print them.
//
//  1. Check out the expected output
//
//
//  2. Check out the code below
//
//     1. header   : stores the column headers
//     2. data     : stores the real data related to each column
//     3. separator: you will use it to separate the data
//
//
//  3. Parse the data
//
//     1. Separate it into cols by using the newline character ("\n")
//
//     2. For each col, separate it by using the separator (",")
//
//
//  4. Load the data into distinct slices
//
//     1. Load the locations into a []string
//     2. Load the sizes into []int
//     3. Load the beds into []int
//     4. Load the baths into []int
//     5. Load the prices into []int
//
//
//  5. Print the header
//
//     1. Separate it by using the separator
//
//     2. Print each column 15 character wide ("%-15s")
//
//
//  6. Print the cols from the slices that you've created, line by line
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//
// HINTS
//
//  + strings.Split function can separate a string into a []string
//
// ---------------------------------------------------------

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)
	var (
		loc                      []string
		size, beds, baths, price []int
	)

	rows := strings.Split(data, "\n")

	for _, row := range rows {
		col := strings.Split(row, ",")

		loc = append(loc, col[0])

		n, _ := strconv.Atoi(col[1])
		size = append(size, n)

		n, _ = strconv.Atoi(col[2])
		beds = append(beds, n)

		n, _ = strconv.Atoi(col[3])
		baths = append(baths, n)

		n, _ = strconv.Atoi(col[4])
		price = append(price, n)

	}

	for _, v := range strings.Split(header, ",") {
		fmt.Printf("%-15s", v)
	}
	fmt.Println("\n", strings.Repeat("=", 75))

	for i, _ := range loc {
		fmt.Printf("%-15s%-15d%-15d%-15d%-15d", loc[i], size[i], beds[i], baths[i], price[i])
		fmt.Println()
	}
	fmt.Println(rows)
}
