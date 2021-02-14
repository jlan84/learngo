// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Unique Words 2
//
//  Use your solution from the previous "Unique Words"
//  exercise.
//
//  Before adding the words to your map, remove the
//  punctuation characters and numbers from them.
//
//
// BE CAREFUL
//
//  Now the shakespeare.txt contains upper and lower
//  case letters too.
//
//
// EXPECTED OUTPUT
//
//  go run main.go < shakespeare.txt
//
//   There are 100 words, 69 of them are unique.
//
// ---------------------------------------------------------

func main() {
	// This is the regular expression pattern you need to use:
	// [^A-Za-z]+
	//
	// Matches to any character but upper case and lower case letters

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	re := regexp.MustCompile(`[^a-zA-z]+`)

	var (
		uniqueWords map[string]int
		wordCount   int
	)
	uniqueWords = make(map[string]int)

	for in.Scan() {
		word := strings.ToLower(in.Text())
		word = re.ReplaceAllString(word, "")
		uniqueWords[word]++
		wordCount++
	}

	fmt.Printf("There are %d words, %d of them are unique.\n", wordCount,
		len(uniqueWords))
}
