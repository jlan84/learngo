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
	"strings"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Days in a Month
//
//  Refactor the previous exercise from the if statement
//  section.
//
//  "Print the number of days in a given month."
//
//  Use a switch statement instead of if statements.
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}

	year := time.Now().Year()
	leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	month := strings.ToLower(os.Args[1])

	months31 := []string{"january", "march", "may", "july", "august",
		"october", "december"}

	months30 := []string{"april", "june", "september", "november"}

	found31 := Find(months31, month)
	found30 := Find(months30, month)

	switch {
	case found31:
		fmt.Printf("%s has 31 days\n", os.Args[1])
	case found30:
		fmt.Printf("%s has 30 days\n", os.Args[1])
	case month == "february" && leap:
		fmt.Printf("%s has 29 days\n", os.Args[1])
	case month == "february":
		fmt.Printf("%s has 28 days\n", os.Args[1])
	default:
		fmt.Printf("%q is not a valid month\n", os.Args[1])
	}

	// if m := strings.ToLower(month); m == "april" ||
	// 	m == "june" ||
	// 	m == "september" ||
	// 	m == "november" {
	// 	days = 30
	// } else if m == "january" ||
	// 	m == "march" ||
	// 	m == "may" ||
	// 	m == "july" ||
	// 	m == "august" ||
	// 	m == "october" ||
	// 	m == "december" {
	// 	days = 31
	// } else if m == "february" {
	// 	if leap {
	// 		days = 29
	// 	}
	// } else {
	// 	fmt.Printf("%q is not a month.\n", month)
	// 	return
	// }

	// fmt.Printf("%q has %d days.\n", month, days)
}

// Find determines if the str is in the slice
func Find(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}
