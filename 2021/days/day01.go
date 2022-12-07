package days

import (
	"adventcode/2021/lib"
	"fmt"
	"strconv"
)

func d1part1(data []string) string {
	increaseCount := 0
	prev, err := strconv.Atoi(data[0])
	if err != nil {
		lib.Perror("%v", err)
		return "null"
	}

	for _, d := range data[1:] {
		t, err := strconv.Atoi(d)
		if err != nil {
			lib.Perror(err.Error())
			return "null"
		}
		if t > prev {
			increaseCount += 1
		}
		prev = t
	}
	return strconv.Itoa(increaseCount)
}

func sum(data []string) int {
	s := 0
	for _, d := range data {
		a, err := strconv.Atoi(d)
		if err != nil {
			lib.Perror(err.Error())
		}
		s += a
	}
	return s
}

func d1part2(data []string) string {
	count := 0
	pwindow := sum(data[0:3])

	for i := 1; i < (len(data) - 2); i += 1 {
		t := sum(data[i : i+3])
		if t > pwindow {
			count += 1
		}
		pwindow = t
		lib.Print("%v: %v, %v", data[i], pwindow, count)
	}
	return strconv.Itoa(count)
}

func Run01(dataset []string) {

	fmt.Println("Hello day 1!")
	lib.Print("Part 1: %v", d1part1(dataset))
	lib.Print("Part 2: %v", d1part2(dataset))
}
