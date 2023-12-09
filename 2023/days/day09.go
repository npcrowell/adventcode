package days

import (
	"advent2023/lib"
	"errors"
	"strconv"
	"strings"
)

func day9_parseline(line string) (ret []int) {
	sline := strings.Split(line, " ")
	for _, snum := range sline {
		num, err := strconv.Atoi(snum)
		if err != nil {
			continue
		}
		ret = append(ret, num)
	}
	return
}

func day9_gendiffs(vals []int) (int, int) {
	var diffs []int
	countzeroes := 0
	for i := range vals[1:] {
		diff := vals[i+1] - vals[i]
		if diff == 0 {
			countzeroes += 1
		}
		diffs = append(diffs, diff)
	}

	prev, next := 0, 0
	if countzeroes != len(diffs) {
		p, n := day9_gendiffs(diffs)
		prev = diffs[0] - p
		next = diffs[len(diffs)-1] + n
	}
	lib.Debug("%v %v %v", prev, diffs, next)
	return prev, next
}

func Day09(part int, data []string) (int, error) {
	switch part {
	case 1:
		fallthrough
	case 2:
		sum := 0
		nsum := 0
		for _, line := range data {
			r := day9_parseline(line)
			lib.Debug("%v", r)
			p, n := day9_gendiffs(r)
			prev := r[0] - p
			next := r[len(r)-1] + n
			lib.Debug("%v %v %v\n", prev, r, next)
			sum += next
			nsum += prev
		}
		if part == 1 {
			return sum, nil
		}
		return nsum, nil
	default:
		return 0, errors.New("not implemented")
	}
}
