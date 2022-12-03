package day01

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"

	e "adventcode/2022/lib"
)

func Run() {
	e.Print("Hello world")

	datafile := "data/d01.txt"

	f, err := os.Open(datafile)
	if err != nil {
		e.Perror("Unable to open data file: %v", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	max := 0
	cur := 0
	var caps []int

	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")

		if line == "" {
			if cur > max {
				max = cur
			}
			caps = append(caps, cur)
			cur = 0
			continue
		}

		val, err := strconv.Atoi(line)
		if err != nil {
			e.Perror("Unable to convert to integer (%v): %v", err, line)
			continue
		}
		cur += val
	}
	e.Print("Max: %v", max)
	sort.Ints(caps)
	e.Print("Top 3: %v", caps[:3])
	e.Print("Bottom 3: %v", caps[len(caps)-3:])
	sum := caps[len(caps)-3] + caps[len(caps)-2] + caps[len(caps)-1]
	e.Print("Sum: %v", sum)
}
