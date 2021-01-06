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
// EXERCISE: Housing Prices and Averages
//
//  Use the previous exercise to solve this exercise (Housing Prices).
//
//  Your task is to print the averages of the sizes, beds, baths, and prices.
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
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
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
		head, rows, location     []string
		size, beds, baths, price []int
		avgs                     []float64
	)

	head = strings.Split(header, separator)
	rows = strings.Split(data, "\n")
	for _, v := range rows {
		cols := strings.Split(v, separator)
		location = append(location, cols[0])
		n, _ := strconv.Atoi(cols[1])
		size = append(size, n)
		n, _ = strconv.Atoi(cols[2])
		beds = append(beds, n)
		n, _ = strconv.Atoi(cols[3])
		baths = append(baths, n)
		n, _ = strconv.Atoi(cols[4])
		price = append(price, n)

	}

	for _, v := range head {
		fmt.Printf("%-15s", v)
	}
	fmt.Print("\n", strings.Repeat("=", 75), "\n")
	for i := range location {
		fmt.Printf("%-15s%-15d%-15d%-15d%-15d\n", location[i], size[i], beds[i],
			baths[i], price[i])
	}
	fmt.Print("\n", strings.Repeat("=", 75), "\n")

	sum := 0.

	for _, v := range size {
		sum += float64(v)
	}
	avgs = append(avgs, sum/float64(len(size)))

	sum = 0.
	for _, v := range beds {
		sum += float64(v)
	}
	avgs = append(avgs, sum/float64(len(beds)))

	sum = 0.
	for _, v := range baths {
		sum += float64(v)
	}
	avgs = append(avgs, sum/float64(len(baths)))

	sum = 0.
	for _, v := range price {
		sum += float64(v)
	}
	avgs = append(avgs, sum/float64(len(price)))

	fmt.Print(strings.Repeat(" ", 15))
	for _, v := range avgs {
		fmt.Printf("%-15.2f", v)
	}
}
