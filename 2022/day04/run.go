package day04

import (
	e "adventcode/2022/lib"
	"fmt"
	"strconv"
	"strings"
)

type elf struct {
	ass []int
}

type elfpair struct {
	elfone elf
	elftwo elf
}

func (e1 elf) contains(e2 elf) bool {
	for _, a := range e2.ass {
		if !rangeContains(e1.ass, a) {
			return false
		}
	}
	return true
}

func (e1 elf) overlaps(e2 elf) bool {
	for _, a := range e2.ass {
		if rangeContains(e1.ass, a) {
			return true
		}
	}
	return false
}

func rangeContains(i []int, in int) bool {
	for _, v := range i {
		if v == in {
			return true
		}
	}

	return false
}

func rangeToList(r string) ([]int, error) {
	vals := strings.Split(r, "-")
	startval, err := strconv.Atoi(vals[0])
	if err != nil {
		return nil, err
	}
	endval, err := strconv.Atoi(vals[1])
	if err != nil {
		return nil, err
	}

	var res []int
	for i := startval; i <= endval; i++ {
		res = append(res, i)
	}

	return res, nil
}

func lineToElfPair(line string) (elfpair, error) {
	elves := strings.Split(line, ",")
	if len(elves) != 2 {
		return elfpair{}, fmt.Errorf("bad line: %v", line)
	}

	elfonerange, err := rangeToList(elves[0])
	if err != nil {
		return elfpair{}, err
	}
	elftworange, err := rangeToList(elves[1])
	if err != nil {
		return elfpair{}, err
	}

	return elfpair{elfone: elf{ass: elfonerange}, elftwo: elf{ass: elftworange}}, nil
}

func part1(data []string) int {
	// var elves []elfpair
	res := 0

	for _, line := range data {
		pair, err := lineToElfPair(line)
		if err != nil {
			e.Perror("%v", err)
			return -1
		}

		if pair.elfone.contains(pair.elftwo) ||
			pair.elftwo.contains(pair.elfone) {
			res += 1
		}

		// elves = append(elves, pair)
	}

	return res
}

func part2(data []string) int {
	// var elves []elfpair
	res := 0

	for _, line := range data {
		pair, err := lineToElfPair(line)
		if err != nil {
			e.Perror("%v", err)
			return -1
		}

		if pair.elfone.overlaps(pair.elftwo) ||
			pair.elftwo.overlaps(pair.elfone) {
			res += 1
		}

		// elves = append(elves, pair)
	}

	return res
}

func Run(datafile string) {
	e.Print("Let's Go Day 03!")

	data, err := e.ReadInTextFile(datafile)
	if err != nil {
		e.Perror("%v", err)
		return
	}

	// Last line of file is blank line
	// Remove empty lines
	data = e.RemoveBlankLines(data)

	e.Printf("Result: %v | %v\n", part1(data), part2(data))
}
