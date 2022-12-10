package day09

import (
	e "adventcode/2022/lib"
	"strconv"
	"strings"
)

func part1(data []string) string {
	nstr := "null"
	k := &knot{h: newCoord(), t: newCoord()}
	k.accountTPos()

	for _, line := range data {
		cmd := strings.Split(line, " ")
		if len(cmd) != 2 {
			e.Perror("wrong command format: %v", line)
			return nstr
		}

		dir := cmd[0]
		dist, err := strconv.Atoi(cmd[1])
		if err != nil {
			e.Perror("distance is bad format: %v", line)
			return nstr
		}

		switch dir {
		case "U":
			k.moveUp(dist)
		case "D":
			k.moveDown(dist)
		case "L":
			k.moveLeft(dist)
		case "R":
			k.moveRight(dist)
		default:
			e.Perror("bad command: %v", line)
			return nstr
		}

	}

	return strconv.Itoa(len(k.prev))
}

func part2(data []string) string {
	nstr := "null"
	snek := newSnake()

	// for i := 0; i < 1; i += 1 {
	for i := 0; i < 9; i += 1 {
		snek.addSection()
	}
	// snek.print()
	snek.head.adjust()

	for _, line := range data {
		cmd := strings.Split(line, " ")
		if len(cmd) != 2 {
			e.Perror("wrong command format: %v", line)
			return nstr
		}

		dir := cmd[0]
		dist, err := strconv.Atoi(cmd[1])
		if err != nil {
			e.Perror("distance is bad format: %v", line)
			return nstr
		}

		switch dir {
		case "U":
			snek.moveUp(dist)
		case "D":
			snek.moveDown(dist)
		case "L":
			snek.moveLeft(dist)
		case "R":
			snek.moveRight(dist)
		default:
			e.Perror("bad command: %v", line)
			return nstr
		}
		// e.Print("")
		snek.print()
	}

	return strconv.Itoa(len(snek.tail.prev))
}

func Run(datafile string, part int) {
	e.Print("Day 09 is your father!")

	// data, err := e.ReadInTextFile("day09/testdata.txt")
	// data, err := e.ReadInTextFile("day09/testdata2.txt")
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
