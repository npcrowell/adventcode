package day12

import e "adventcode/2022/lib"

func part1(data []string) string {
	nstr := "null"

	return nstr
}

func part2(data []string) string {
	nstr := "null"

	if len(data) == 0 {
		e.Perror("Empty data set")
		return nstr
	}

	g := e.EmptyGrid(uint(len(data)), uint(len(data[0])))

	for i, line := range data {
		for j, char := range line {
			g.SetIndex(uint(i), uint(j), int(char))
		}
	}

	return nstr
}

func Run(datafile string, part int) {
	e.Print("Day 12 is not yet ready!")

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
