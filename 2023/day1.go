package main

import (
	"advent2023/lib"
	"errors"
	"strconv"
)

func day1_parseLineForDigits(line string) int {
	var first, last int

	for i := 0; i < len(line); i += 1 {
		f, err := strconv.Atoi(string(line[i]))
		if err == nil {
			first = f
			break
		}
	}

	for i := len(line) - 1; i >= 0; i -= 1 {
		l, err := strconv.Atoi(string(line[i]))
		if err == nil {
			last = l
			break
		}
	}

	return first*10 + last
}

func day1_parseLine(line string) int {
	subs := lib.GetSubstrings(
		line,
		[]string{
			"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
			"one", "two", "three", "four", "five",
			"six", "seven", "eight", "nine",
		},
		false,
	)
	// lib.Printf("%v\n", subs)

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

var d1p1_testdata []string = []string{
	"1abc2",
	"pqr3stu8vwx",
	"a1b2c3d4e5f",
	"treb7uchet",
}

var d1p2_testdata []string = []string{
	"two1nine",
	"eightwothree",
	"abcone2threexyz",
	"xtwone3four",
	"4nineeightseven2",
	"zoneight234",
	"7pqrstsixteen",
}

func day1(part int, data []string, useTestData bool) (int, error) {

	switch part {
	case 1:
		if useTestData {
			data = d1p1_testdata
		}
		sum := 0
		for i, line := range data {
			val := day1_parseLineForDigits(line)
			lib.Debug("%4d: (%v) %v", i, val, line)
			sum += val
		}
		return sum, nil
	case 2:
		if useTestData {
			data = d1p2_testdata
		}
		sum := 0
		for i, line := range data {
			val := day1_parseLine(line)
			lib.Debug("%4d: (%v) %v", i, val, line)
			sum += val
		}
		return sum, nil
	default:
		return 0, errors.New("not implemented")
	}
}
