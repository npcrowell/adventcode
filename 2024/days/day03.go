package days

import (
	"advent2024/lib"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func d3_parsecomplex(data string) (sum int) {
	do := true

	// if do, cut the start of data until the next instance of "don't"
	for len(data) > 0 {
		if do {
			next := strings.Index(data, "don't()")
			if next == -1 {
				sum += d3_parsesimple(data)

				return
			}

			mystr := data[:next+len("don't()")]
			// lib.Printf("%v\n", mystr)
			sum += d3_parsesimple(mystr)
			next += len("don't")
			data = data[next:]
			do = false
		} else {
			next := strings.Index(data, "do()")
			if next == -1 {
				break
			}

			data = data[next:]
			do = true
		}
	}

	return
}

func d3_parsesimple(data string) int {
	m := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`).FindAllStringSubmatch(data, -1)
	// lib.Printf("Found %v matches\n", len(m))
	// lib.Printf("Matches: %v\n", m)

	sum := 0
	for _, v := range m {
		lib.Printf("%v ", v[0])
		a, _ := strconv.Atoi(v[1])
		b, _ := strconv.Atoi(v[2])
		sum += a * b
	}
	return sum
}

func Day03(part int, data []string) (int, error) {
	cdata := strings.Join(data, "")
	switch part {
	case 1:
		return d3_parsesimple(cdata), nil
		// return 0, nil
	case 2:
		return d3_parsecomplex(cdata), nil
	default:
		return 0, errors.New("not implemented")
	}
}
