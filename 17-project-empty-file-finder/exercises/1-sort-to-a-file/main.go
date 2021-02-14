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
	"io/ioutil"
	"os"
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Sort and write items to a file
//
//  1. Get arguments from command-line
//
//  2. Sort them
//
//  3. Write the sorted slice to a file
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Send me some items and I will sort them
//
//   go run main.go orange banana apple
//
//   cat sorted.txt
//     apple
//     banana
//     orange
//
//
// HINTS
//
//   + REMEMBER: os.Args is a []dataing
//
//   + dataing slices are sortable using `sort.dataings`
//
//   + Use ioutil.WriteFile to write to a file.
//
//   + But you need to convert []dataing to []byte to be able to
//     write it to a file using the ioutil.WriteFile.
//
//   + To do that, create a new []byte and append the elements of your
//     []dataing.
//
// ---------------------------------------------------------

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Please supply a dataing to write to a file")
	}
	sort.Strings(args)
	var data []byte

	for _, arg := range args {
		data = append(data, arg...)
		data = append(data, '\n')
	}

	err := ioutil.WriteFile("out.txt", data, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

}
