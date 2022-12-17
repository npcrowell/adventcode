package day13

import (
	e "adventcode/2022/lib"
	"fmt"
	"strconv"
)

func parseInt(line string) (int, string, error) {
	var s string
	// e.Print("Received line: '%v'", line)
	// defer e.Print("Parsed int: '%v'", s)

	for i, c := range line {
		if c >= '0' && c <= '9' {
			s += string(c)
			// e.Print("meanwhile: '%v'", s)
		} else if len(s) < 1 {
			return -1, "", fmt.Errorf("line did not beget int: %v", line)
		} else {
			parsedint, err := strconv.Atoi(s)
			if err != nil {
				return -1, "", err
			}
			// e.Debug("parseint Returning %v , \"%v\"", parsedint, line[i:])
			return parsedint, line[i:], nil
		}
	}

	parsedint, err := strconv.Atoi(s)
	if err != nil {
		return -1, "", err
	}

	// e.Debug("parseint Returning %v", parsedint)
	return parsedint, "", nil
}

func parseList(line string) (*item, string, error) {
	i := 1
	it := &item{t: list}
	// oline := line
	// e.Debug("parselist received: '%v'", line)

	for i < len(line) {
		c := line[i]
		// e.Print("LINE: %v", line[i:])

		if c >= '0' && c <= '9' {
			num, l, err := parseInt(line[i:])
			if err != nil {
				return nil, "", err
			}
			// e.Print("Parse int returned: '%v'", num)
			numitem := &item{t: number, nval: num}
			// e.Print("numitem created: '%v'", numitem.String())
			it.subitem = append(it.subitem, numitem)
			line = l
			i = 0
			continue

		} else if c == ']' {
			it.val = line[:i+1]
			if i+2 > len(line) {
				// e.Debug("parselist returning: %v", it)
				return it, "", nil
			}

			// e.Debug("parselist returning: %v", it)
			return it, line[i+1:], nil

		} else if c == '[' {
			litem, l, err := parseList(line[i:])
			if err != nil {
				return nil, "", err
			}
			line = l
			it.subitem = append(it.subitem, litem)
			// e.Debug("line from subparselist: '%v'", line)
			i = 0

		} else if c == ',' {
			i += 1
			continue
		}
	}

	return nil, "", fmt.Errorf("unended list: %v", line)
}

// func parse(line string) *item {
// 	if len(line) == 0 {
// 		return nil
// 	}

// 	_, err := strconv.Atoi(string(line[0]))
// 	var i *item
// 	var s string

// 	for len(line) > 0 {
// 		if line[0] == '[' {
// 			i, s = parseList(line)
// 		} else if err != nil {
// 			i, s = parseInt(line)
// 		} else {
// 			e.Perror("Unconditioned symbol at start of line: %s", line)
// 		}
// 	}

// 	return i
// }

func part1(data []string) string {
	nstr := "null"
	count := 0

	for i := 0; i < len(data); i += 3 {
		e.Print("== Pair %v ==", (i/3)+1)
		leftstr := data[i]
		litem, _, err := parseList(leftstr)
		// e.Print("left %v", left)

		if err != nil {
			e.Perror(err.Error())
			return nstr
		}
		// e.Print(l)
		e.Print("%v", litem)
		// e.Print("")

		right := data[i+1]
		ritem, _, err := parseList(right)
		// e.Print("right %v", right)

		if err != nil {
			e.Perror(err.Error())
			return nstr
		}
		// e.Print(l)
		e.Print("%v", ritem)

		res := litem.Compare(ritem)

		if res == equal {
			e.Print("Equal!")
		} else if res == left {
			e.Print("Correct Order!")
			count += (i / 3) + 1
		} else {
			e.Print("Incorrect Order!")
		}
		e.Print("")
		// return nstr
	}

	return strconv.Itoa(count)
}

func part2(data []string) string {
	nstr := "null"

	data = append(data, "[[2]]", "[[6]]")
	var ilist []*item

	for i := 0; i < len(data); i += 1 {
		if data[i] == "" {
			continue
		}

		it, _, err := parseList(data[i])
		if err != nil {
			e.Perror(err.Error())
			return nstr
		}
		ilist = append(ilist, it)
	}

	ilist = itemsort(ilist)

	// e.Print("After:")
	b1 := 0
	b2 := 0
	test1, _, _ := parseList("[[2]]")
	test2, _, _ := parseList("[[6]]")
	for j, i := range ilist {
		if i.Compare(test1) == equal {
			b1 = j + 1
		} else if i.Compare(test2) == equal {
			b2 = j + 1
		}
	}
	// e.Print("")

	return strconv.Itoa(b1 * b2)
}

func Run(datafile string, part int) {
	e.Print("Day 13 is skipping day 12!")

	// data, err := e.ReadInTextFile("day13/testdata.txt")
	data, err := e.ReadInTextFile(datafile)
	if err != nil {
		e.Perror("%v", err)
		return
	}
	data = append(data, "")

	switch part {
	case 1:
		e.Print("Part 1: %v", part1(data))
	case 2:
		e.Print("Part 2: %v", part2(data))
	default:
		e.Perror("Unregognized part number: %v", part)
	}
}
