package days

import (
	"advent2023/lib"
	"errors"
	"math"
	"strconv"
	"strings"
)

// Template file

func day4_parseLine(line string) [][]int {
	card := strings.Trim(strings.Split(line, ": ")[0], " ")
	// lib.Debug("%v", card)

	nums := strings.Trim(strings.Split(line, ": ")[1], " ")

	wsnums := strings.Split(strings.Split(nums, " | ")[0], " ")
	// lib.Debug("%v(%v)", wsnums, len(wsnums))
	asnums := strings.Split(strings.Split(nums, " | ")[1], " ")
	// lib.Debug("%v(%v)", asnums, len(asnums))

	var wnums, anums, cnums []int

	pcard := strings.Split(card, " ")
	cnum, err := strconv.Atoi(pcard[len(pcard)-1])
	if err != nil {
		lib.Perror("Unable to convert '%v', '%v': %v", card, strings.Split(card, " ")[1], err)
	} else {
		cnums = append(cnums, cnum)
	}

	for _, wsnum := range wsnums {
		wnum, err := strconv.Atoi(wsnum)
		if err != nil {
			// lib.Debug("Unable to convert '%v': %v", wsnum, err)
			continue
		}
		wnums = append(wnums, wnum)
	}

	for _, asnum := range asnums {
		anum, err := strconv.Atoi(asnum)
		if err != nil {
			// lib.Debug("Unable to convert '%v': %v", asnum, err)
			continue
		}
		anums = append(anums, anum)
	}

	return [][]int{cnums, wnums, anums}
}

func day4_countcontains(alist []int, blist []int) []int {
	var res []int

	for _, a := range alist {
		for _, b := range blist {
			if a == b {
				res = append(res, a)
			}
		}
	}

	return res
}

func day4_calcLineSum(line string) int {
	pline := day4_parseLine(line)
	lib.Debug("%v", pline)

	count := len(day4_countcontains(pline[1], pline[2]))
	if count == 0 {
		return 0
	}
	return int(math.Pow(2.0, float64(count-1)))
}

func day4_calcLineCount(line string) int {
	pline := day4_parseLine(line)
	lib.Debug("%v", pline)

	return len(day4_countcontains(pline[1], pline[2]))
}

func Day04(part int, data []string) (int, error) {
	switch part {
	case 1:
		sum := 0
		for _, line := range data {
			sum += day4_calcLineSum(line)
		}
		return sum, nil
	case 2:
		sum := 0
		cards := make([]int, len(data))
		for c := range cards {
			cards[c] = 1
		}
		for i, line := range data {
			count := day4_calcLineCount(line)
			for j := 1; j < count+1; j += 1 {
				cards[i+j] += cards[i]
			}
		}
		for c := range cards {
			sum += cards[c]
			lib.Debug("%v: %v", c+1, cards[c])
		}
		return sum, nil
	default:
		return 0, errors.New("not implemented")
	}
}
