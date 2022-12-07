package days

import (
	"adventcode/2021/lib"
	"fmt"
	"strconv"
	"strings"
)

func d2part1(data []string) string {
	hpos := 0
	dpos := 0

	for _, line := range data {
		l := strings.Split(line, " ")
		if len(l) != 2 {
			lib.Perror("Line split weird: %v", line)
			return "null"
		}
		dir := l[0]
		dist, err := strconv.Atoi(l[1])
		if err != nil {
			lib.Perror(err.Error())
			return "null"
		}
		switch dir {
		case "forward":
			hpos += dist
		case "down":
			dpos += dist
		case "up":
			dpos -= dist
		default:
			lib.Perror("Bad direction: %v", line)
		}
	}
	return strconv.Itoa(hpos * dpos)
}

func d2part2(data []string) string {
	hpos := 0
	dpos := 0
	aim := 0

	for _, line := range data {
		l := strings.Split(line, " ")
		if len(l) != 2 {
			lib.Perror("Line split weird: %v", line)
			return "null"
		}
		dir := l[0]
		dist, err := strconv.Atoi(l[1])
		if err != nil {
			lib.Perror(err.Error())
			return "null"
		}
		switch dir {
		case "forward":
			hpos += dist
			dpos += aim * dist
		case "down":
			aim += dist
		case "up":
			aim -= dist
		default:
			lib.Perror("Bad direction: %v", line)
		}
	}
	return strconv.Itoa(hpos * dpos)
}

func Run02(dataset []string) {
	fmt.Println("Let's go day 2!")
	lib.Print("Part 1: %v", d2part1(dataset))
	lib.Print("Part 2: %v", d2part2(dataset))
}
