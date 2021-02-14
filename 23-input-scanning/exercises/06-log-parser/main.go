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
	"sort"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Log Parser from Stratch
//
//  You've watched the lecture. Now, try to create the same
//  log parser program on your own. Do not look at the lecture,
//  and the existing source code.
//
//
// EXPECTED OUTPUT
//
//  go run main.go < log.txt
//
//   DOMAIN                             VISITS
//   ---------------------------------------------
//   blog.golang.org                        30
//   golang.org                             10
//   learngoprogramming.com                 20
//
//   TOTAL                                  60
//
//
//  go run main.go < log_err_missing.txt
//
//   wrong input: [golang.org] (line #3)
//
//
//  go run main.go < log_err_negative.txt
//
//   wrong input: "-100" (line #3)
//
//
//  go run main.go < log_err_str.txt
//
//   wrong input: "FOUR" (line #3)
//
// ---------------------------------------------------------

func main() {
	var (
		domain  string
		counter int
		total   int
		domains []string
	)

	in := bufio.NewScanner(os.Stdin)
	logDic := make(map[string]int)

	for in.Scan() {
		counter++
		words := strings.Fields(in.Text())
		if len(words) < 2 {
			fmt.Printf("wrong input: %q (line #%d)\n", words[0], counter)
			return
		}
		domain = words[0]
		visits, err := strconv.Atoi(words[1])
		if err != nil || visits < 0 {
			fmt.Printf("wrong input: %q (line #%d)\n", words[1], counter)
			return
		}
		if _, exists := logDic[domain]; !exists {
			domains = append(domains, domain)
		}
		logDic[domain] += visits
		total += visits
	}

	fmt.Printf("%-30s%10s\n%s\n", "Domain", "Visits", strings.Repeat("-", 50))
	sort.Strings(domains)

	for _, v := range domains {
		fmt.Printf("%-30s%10d\n", v, logDic[v])
	}

	fmt.Printf("\n%-30s%10d\n", "Total", total)

	if err := in.Err(); err != nil {
		fmt.Printf("> err%s", err)
	}
}
