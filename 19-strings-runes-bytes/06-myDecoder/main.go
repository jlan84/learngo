package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var start, stop int

	if args := os.Args[1:]; len(args) == 2 {
		start, _ = strconv.Atoi(args[0])
		stop, _ = strconv.Atoi(args[1])
	}

	if start == 0 || stop == 0 {
		start, stop = 'a', 'z'
	}

	fmt.Printf("%-10s%-10s%-10s%s\n%s\n", "literal", "dec", "hex", "encoded",
		strings.Repeat("-", 45))
	for i := start; i <= stop; i++ {
		fmt.Printf("%-10c%-10d%-10x% x\n", i, i, i, string(i))
	}
	
	

}
