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
	lerr    error
}

func newParser() parser {
	return parser{sum: make(map[string]result)}
}
func err(p *parser) error {
	return p.lerr
}

func parse(p *parser, line string) (r result) {
	if p.lerr != nil {
		return
	}
	p.lines++
	fields := strings.Fields(line)
	if len(fields) != 2 {
		p.lerr = fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
		return
	}
	var err error
	r.domain = fields[0]

	// Sum the total visits per domain
	r.visits, err = strconv.Atoi(fields[1])
	if r.visits < 0 || err != nil {
		p.lerr = fmt.Errorf("wrong input: %q (line #%d)", fields[1], p.lines)
	}
	return
}

func sortPrint(p *parser) {
	sort.Strings(p.domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range p.domains {
		fmt.Printf("%-30s %10d\n", domain, p.sum[domain].visits)
	}

	// Print the total visits for all domains
	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)

}

func update(p *parser, r result) {
	if p.lerr != nil {
		return
	}
	if _, ok := p.sum[r.domain]; !ok {
		p.domains = append(p.domains, r.domain)
	}

	// Keep track of total and per domain visits
	p.total += r.visits

	// You cannot assign to composite values
	// p.sum[domain].visits += visits

	// create and assign a new copy of `visit`
	p.sum[r.domain] = result{
		domain: r.domain,
		visits: r.visits + p.sum[r.domain].visits,
	}
	// following line complains about being not addressable
	// map elements are not addressable like slice/array/struct elements/fields
	// It is because sometimes Go runtime can move the elements in the maps in the memory
	// without asking you, hence you can't get its address from the map
	// p.sum[parsed.domain].visits += parsed.visits

	// when you access the element you get its copy
	// if you run following piece of code you can't actually modify
	// p.sum[parsed.domain].visits by adding +1 to v
	// v := p.sum[parsed.domain].visits
	// v += 1
	// fmt.Println(v, p.sum[parsed.domain].visits)

	// and following says: compilerUnaddressableOperand
	// v := &p.sum[parsed.domain].visits
}
