package days

import (
	"advent2023/lib"
	"errors"
)

func day1_parseLine(line string, digits []string) int {
	subs := lib.GetSubstrings(line, digits, false)

	first, err := lib.StringToDigit(subs[0], 10)
	if err != nil {
		lib.Perror("Error converting %v to digit: %v", subs[0], err)
	}
	last, err := lib.StringToDigit(subs[len(subs)-1], 10)
	if err != nil {
		lib.Perror("Error converting %v to digit: %v", subs[len(subs)-1], err)
	}
	return first*10 + last
}

func Day01(part int, data []string) (int, error) {

	switch part {
	case 1:
		sum := 0
		for i, line := range data {
			val := day1_parseLine(line,
				[]string{
					"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
				})
			lib.Debug("%4d: (%v) %v", i, val, line)
			sum += val
		}
		return sum, nil
	case 2:
		sum := 0
		for i, line := range data {
			val := day1_parseLine(line,
				[]string{
					"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
					"one", "two", "three", "four", "five",
					"six", "seven", "eight", "nine",
				})
			lib.Debug("%4d: (%v) %v", i, val, line)
			sum += val
		}
		return sum, nil
	default:
		return 0, errors.New("not implemented")
	}
}
