package main

import (
	"advent2023/lib"
	e "advent2023/lib"
	"flag"
	"fmt"
)

func main() {
	// Initialize
	pday := flag.Int("day", 1, "Day number to run")
	ppart := flag.Int("part", 1, "Part of day to run")
	ptest := flag.Bool("test", false, "Use the test data")
	pdbg := flag.Bool("debug", false, "Print debug strings")
	flag.Parse()
	day := *pday
	part := *ppart
	test := *ptest
	lib.Dbg = *pdbg

	// Setup
	days := [2]func(int, []string, bool) (int, error){day0, day1}

	// Validation
	if 0 > day || day >= len(days) {
		e.Perror("Bad day %v\n", day)
	}

	filename := fmt.Sprintf("data/d%02d.txt", day)
	data, err := e.ReadInTextFile(filename)
	if err != nil {
		e.Perror("Unable to open file: %v\n", err)
	}

	// Execution
	r, err := days[day](part, data, test)
	if err != nil {
		e.Perror("Day %d, part %d, error: %v\n", day, part, err)
		return
	}

	// Results
	e.Printf("Day %d, part %d = %d\n", day, part, r)
}
