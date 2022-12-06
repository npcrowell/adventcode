package main

import (
	"flag"
	"fmt"
)

var (
	day  *int
	part *int
)

func init() {
	day = flag.Int("day", 1, "Day number to run")
	part = flag.Int("part", 1, "Part of day to run")
}

func main() {
	flag.Parse()
	datafile := fmt.Sprintf("data/d%02d.txt", *day)
	run(*day, *part, datafile)
}
