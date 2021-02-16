package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// result stores the parsed result for a domain
type result struct {
	domain string
	visits int
	// add more metrics if needed
}

// parser keep tracks of the parsing
type parser struct {
	sum     map[string]result // metrics per domain
	domains []string          // unique domain names
	total   int               // total visits for all domains
	lines   int               // number of parsed lines (for the error messages)
}

func makeParser() (p parser) {
	p = parser{sum: make(map[string]result)}
	return
}

func parseText(line string, p parser) (pResult result, err error) {
	fields := strings.Fields(line)

	if len(fields) != 2 {
		err = fmt.Errorf("%q wrong input (line #%d", fields, p.lines)
		return
	}

	pResult.domain = fields[0]
	pResult.visits, err = strconv.Atoi(fields[1])

	if pResult.visits < 0 || err != nil {
		err = fmt.Errorf("%q wrong input (line #%d", fields[1], p.lines)
		return
	}

	return
}

func collectDomains(p parser, r result) parser {
	domain, visits := r.domain, r.visits

	if _, exists := p.sum[domain]; !exists {
		p.domains = append(p.domains, domain)
	}
	p.total += visits

	p.sum[domain] = result{
		domain: domain,
		visits: visits + p.sum[domain].visits,
	}
	return p
}

func printLog(p parser) {
	sort.Strings(p.domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range p.domains {
		parsed := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.visits)
	}

	// Print the total visits for all domains
	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)

	// Let's handle the error

}
