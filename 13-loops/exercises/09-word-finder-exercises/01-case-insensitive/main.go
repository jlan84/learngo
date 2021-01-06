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
	"log"
	"os"
	"regexp"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Case Insensitive Search
//
//  Allow for case-insensitive searching
//
// EXAMPLE
//  Let's say that the user runs the program like this:
//    go run main.go LAZY
//
//  Or like this: go run main.go lAzY
//  Or like this: go run main.go lazy
//
//  For all cases above, the program should find
//  the "lazy" keyword.
// ---------------------------------------------------------
var sentence = `Hello, my name is Justin. I like to snowboard, golf, and hike.
I live in Colorado and my favorite things about Colorado are the mountains and 
the seasons.`

func main() {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedSentence := reg.ReplaceAllString(sentence, " ")
	fmt.Println(processedSentence)
	words := strings.Fields(processedSentence)
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Please enter at lease one word")
		return
	}

	for _, q := range args {
		for i, w := range words {
			if strings.ToLower(q) == strings.ToLower(w) {
				fmt.Printf("#%-2d: %s\n", i+1, w)
				break
			}
		}
	}

}
