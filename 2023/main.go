package main

import (
	"advent2023/data"
	"advent2023/lib"
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
	days := []func(int, []string) (int, error){
		day0, day1, day2,
	}

	// Validation, exit on failure
	if 0 > day || day >= len(days) {
		lib.Perror("Bad day %v", day)
		return
	}

	var d []string
	if test {
		tdatastr := fmt.Sprintf("day%dpart%d", day, part)
		d = data.Testdata[tdatastr]
	} else {
		// Load file
		filename := fmt.Sprintf("data/d%02d.txt", day)
		dt, err := lib.ReadInTextFile(filename)
		if err != nil {
			// There should be a file, but sometimes there might not be
			// Alert and continue execution
			lib.Perror("Unable to open file: %v", err)
		}
		d = dt
	}

	// Execution
	r, err := days[day](part, d)
	if err != nil {
		lib.Perror("Day %d, part %d, error: %v", day, part, err)
		return
	}

	// Results
	lib.Psuccess("Day %d, part %d = %d", day, part, r)
}
