package day10

import e "adventcode/2022/lib"

func part1(data []string) string { return "null" }
func part2(data []string) string { return "null" }

func Run(datafile string, part int) {
	e.Perror("Day 10 is not yet implemented!")
	return

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
