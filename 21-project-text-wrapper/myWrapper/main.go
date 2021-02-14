package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {

	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Enter [max line length] [some text]")
		return
	}

	var (
		lineLength, _ = strconv.Atoi(args[0])
		text          = args[1]
		lineWidth     int
	)

	for _, r := range text {
		fmt.Printf("%c", r)

		switch lineWidth++; {
		case lineWidth >= lineLength && r != '\n' && unicode.IsSpace(r):
			fmt.Println()
			fallthrough
		case r == '\n':
			lineWidth = 0
		}
	}
}
