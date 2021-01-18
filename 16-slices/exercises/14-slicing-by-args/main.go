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
// EXERCISE: Slicing by arguments
//
//   We've a []string, you will get arguments from the command-line,
//   then you will print only the elements that are requested.
//
//   1. Print the []string (it's in the code below)
//
//   2. Get the starting and stopping positions from the command-line
//
//   3. Print the []string slice by slicing it using the starting and stopping
//      positions
//
//   4. Handle the error cases (see the expected output below)
//
//   5. Add new elements to the []string slice literal.
//      Your program should work without changing the rest of the code.
//
//   6. Now, play with your program, get a deeper sense of how the slicing
//      works.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    Provide only the [starting] and [stopping] positions
//
//
//  (error because: we expect only two arguments)
//
//  go run main.go 1 2 4
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    Provide only the [starting] and [stopping] positions
//
//
//  (error because: starting index >= 0 && stopping pos <= len(slice) )
//
//  go run main.go -1 5
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    Wrong positions
//
//
//  (error because: starting <= stopping)
//
//  go run main.go 3 2
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    Wrong positions
//
//
//  go run main.go 0
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Normandy Verrikan Nexus Warsaw]
//
//
//  go run main.go 1
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Verrikan Nexus Warsaw]
//
//
//  go run main.go 2
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Nexus Warsaw]
//
//
//  go run main.go 3
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Warsaw]
//
//
//  go run main.go 4
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    []
//
//
//  go run main.go 0 3
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Normandy Verrikan Nexus]
//
//
//  go run main.go 1 3
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Verrikan Nexus]
//
//
//  go run main.go 1 2
//    ["Normandy" "Verrikan" "Nexus" "Warsaw"]
//
//    [Verrikan]
//
// ---------------------------------------------------------

func main() {
	const (
		errComment  = "Please provide a starting or a start and stop position"
		errComment2 = `Please make sure the start position is < the stop
					   position and neither are negative`
	)

	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}

	args := os.Args[1:]
	if len(args) < 1 || len(args) > 2 {
		fmt.Println(errComment)
		return
	} else if len(args) == 1 {
		start, err := strconv.Atoi(args[0])
		if err != nil || start < 0 {
			fmt.Println(errComment)
			return
		}
		slc := ships[start:]
		fmt.Printf("Query: %s\n", slc)
	} else {
		start, err1 := strconv.Atoi(args[0])
		stop, err2 := strconv.Atoi(args[1])
		if err1 != nil || start < 0 || stop < 0 {
			fmt.Println(errComment)
			return
		} else if err2 != nil {
			slc := ships[start:]
			fmt.Printf("Query: %s\n", slc)
		} else {
			slc := ships[start:stop]
			fmt.Printf("Query: %s\n", slc)
		}

	}
}
