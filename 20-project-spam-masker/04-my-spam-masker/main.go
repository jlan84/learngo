package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	const (
		start      = 33
		end        = 126
		link       = "https://"
		linkLength = len(link)
		mask       = '*'
		prefix     = "https://"
	)

	fmt.Printf("%-10s%-10s%-10s%s\n%s\n", "literal", "dec", "hex", "encoded",
		strings.Repeat("=", 50))

	for i := start; i <= end; i++ {
		fmt.Printf("%-10c%-10d%-10x% x\n", i, i, i, i)
	}
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Enter some text")
		return
	}
	var (
		text   = args[0]
		length = len(text)
		buf    = make([]byte, 0, length)
		inside bool
	)

	for i := 0; i < length; i++ {
		if len(text[i:]) >= linkLength && text[i:i+linkLength] == link {
			inside = true
			buf = append(buf, link...)
			i = i + linkLength
		}
		c := text[i]
		switch c {
		case ' ', '\t', '\n':
			inside = false
		}
		if inside {
			c = mask
		}
		buf = append(buf, c)
	}

	fmt.Println(string(buf))
}
