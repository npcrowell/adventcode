package day06

import (
	e "adventcode/2022/lib"
	"strconv"
	"strings"
)

func part1(data []string) string {
	d := data[0]

	// found := false
	index := "null"
outLoop:
	for i := 4; i < len(d); i += 1 {
		substr := d[i-4 : i]
		for _, char := range substr {

			if strings.Count(substr, string(char)) > 1 {
				continue outLoop
			}
		}
		index = strconv.Itoa(i)
		break
	}
	return index
}

func part2(data []string) string {
	d := data[0]

	// found := false
	index := "null"
outLoop:
	for i := 14; i < len(d); i += 1 {
		substr := d[i-14 : i]
		for _, char := range substr {

			if strings.Count(substr, string(char)) > 1 {
				continue outLoop
			}
		}
		index = strconv.Itoa(i)
		break
	}
	return index
}

func Run(datafile string, part int) {
	e.Print("Day 06 let's fcckking gooooo!")
	// return

	data, err := e.ReadInTextFile(datafile)
	if err != nil {
		e.Perror("%v", err)
		return
	}

	switch part {
	case 1:
		e.Print("Part 1: %v", part1(data))
	case 2:
		e.Print("Part 2: %v", part2(data))
	default:
		e.Perror("Unregognized part number: %v", part)
	}
}
