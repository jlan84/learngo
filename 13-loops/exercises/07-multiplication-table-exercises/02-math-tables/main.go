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
// EXERCISE: Math Tables
//
//  Create division, addition and subtraction tables
//
//  1. Get the math operation and
//     the size of the table from the user
//
//  2. Print the table accordingly
//
//  3. Correctly handle the division by zero error
//
//
// BONUS #1
//
//  Use strings.IndexAny function to detect
//    the valid operations.
//
//  Search on Google for: golang pkg strings IndexAny
//
// BONUS #2
//
//  Add remainder operation as well (remainder table using %).
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//    Usage: [op=*/+-] [size]
//
//  go run main.go "*"
//    Size is missing
//    Usage: [op=*/+-] [size]
//
//  go run main.go "%" 4
//    Invalid operator.
//    Valid ops one of: */+-
//
//  go run main.go "*" 4
//    *    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    1    2    3    4
//    2    0    2    4    6    8
//    3    0    3    6    9   12
//    4    0    4    8   12   16
//
//  go run main.go "/" 4
//    /    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    1    0    0    0
//    2    0    2    1    0    0
//    3    0    3    1    1    0
//    4    0    4    2    1    1
//
//  go run main.go "+" 4
//    +    0    1    2    3    4
//    0    0    1    2    3    4
//    1    1    2    3    4    5
//    2    2    3    4    5    6
//    3    3    4    5    6    7
//    4    4    5    6    7    8
//
//  go run main.go "-" 4
//    -    0    1    2    3    4
//    0    0   -1   -2   -3   -4
//    1    1    0   -1   -2   -3
//    2    2    1    0   -1   -2
//    3    3    2    1    0   -1
//    4    4    3    2    1    0
//
// BONUS:
//
//  go run main.go "%" 4
//    %    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    0    1    1    1
//    2    0    0    0    2    2
//    3    0    0    1    0    3
//    4    0    0    0    1    0
//
// NOTES:
//
//   When running the program in Windows Git Bash, you might need
//   to escape the characters like this.
//
//   This happens because those characters have special meaning.
//
//   Division:
//     go run main.go // 4
//
//   Multiplication and others:
//   (this is also necessary for non-Windows bashes):
//
//     go run main.go "*" 4
// ---------------------------------------------------------

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Please provide an integer and an operator")
		return
	}
	num, err := strconv.Atoi(args[2])
	op := args[1]
	fmt.Println(op)
	if err != nil || num < 0 {
		fmt.Println("This is not a valid integer")
		return
	}
	fmt.Printf("%s", op)
	for i := 0; i <= num; i++ {
		fmt.Printf("%4d", i)
	}
	switch op {

	case "*":
		for i := 0; i <= num; i++ {
			fmt.Printf("\n%d", i)
			for j := 0; j <= num; j++ {
				fmt.Printf("%4d", i*j)
			}
		}

	case "/":
		for i := 0; i <= num; i++ {
			fmt.Printf("\n%d", i)
			for j := 0; j <= num; j++ {
				if i == 0 {
					fmt.Printf("%4d", 0)
				} else {
					fmt.Printf("%4d", i*j)
				}
			}
		}

	case "+":
		for i := 0; i <= num; i++ {
			fmt.Printf("\n%d", i)
			for j := 0; j <= num; j++ {
				fmt.Printf("%4d", i+j)
			}
		}

	case "-":
		for i := 0; i <= num; i++ {
			fmt.Printf("\n%d", i)
			for j := 0; j <= num; j++ {
				fmt.Printf("%4d", i-j)
			}
		}
	case "%":
		for i := 0; i <= num; i++ {
			fmt.Printf("\n%d", i)
			for j := 0; j <= num; j++ {
				if i == 0 {
					fmt.Printf("%4d", 0)
				} else {
					fmt.Printf("%4d", j%i)
				}

			}
		}
	default:
		fmt.Println("Please enter a valid opertor")
	}
}
