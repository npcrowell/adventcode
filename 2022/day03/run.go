package day03

import (
	e "adventcode/2022/lib"
	"fmt"
	"strings"
	"time"
)

func getLetterValue(letter rune) int {
	if letter < 91 {
		return int(uint(letter) - 65 + 27)
	}
	if letter < 123 {
		return int(uint(letter) - 96)
	}
	return -1
}

func getComponentSharedItem(comp1 string, comp2 string) (rune, error) {
	for _, letter := range comp1 {
		if strings.Contains(comp2, string(letter)) {
			return letter, nil
		}
	}
	return -1, fmt.Errorf("no shared characters")
}

func part1(data []string) int {
	res := 0

	for i, line := range data {
		linelen := len(line)
		// e.Printf("%4v, (%v) %v\n", i, linelen, line)

		comp1 := line[:linelen/2]
		comp2 := line[linelen/2:]

		letter, err := getComponentSharedItem(comp1, comp2)
		if err != nil {
			e.Perror("%v", err)
			time.Sleep(1 * time.Second)
		}

		value := getLetterValue(letter)

		e.Printf("%4v, %c, %4v, %v\n", i, letter, letter, value)
		res += value
	}
	return res
}

func getTeamSharedItem(line []string) rune {
	if len(line) != 3 {
		e.Perror("Team size is incorrect (%v)", len(line))
		return -1
	}

	for _, l := range line[0] {
		letter := string(l)
		if strings.Contains(line[1], letter) &&
			strings.Contains(line[2], letter) {
			return l
		}
	}

	return -1
}

func part2(data []string) int {
	res := 0
	for i := 0; i < len(data); i += 3 {
		shareditem := getTeamSharedItem(data[i : i+3])
		res += getLetterValue(shareditem)
	}
	return res
}

func testPrint(letter rune) {
	e.Printf("%c, %4v, %4v\n", letter, letter, getLetterValue(letter))
}

func Run(wordlist string) {
	e.Print("Let's Go Day 03!")

	data, err := e.ReadInTextFile(wordlist)
	if err != nil {
		e.Perror("%v", err)
		return
	}

	// Last line of file is blank line
	// Remove empty lines
	data = e.RemoveBlankLines(data)

	e.Printf("Result: %v\n", part2(data))
}
