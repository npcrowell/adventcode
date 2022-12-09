package day08

import (
	"adventcode/2022/lib"
	"strconv"
)

func part1(data []string) string {
	grid := lib.NewGrid()

	grid.AddRow(uint(len(data)))

	for _, line := range data {
		_, columnBound := grid.GetBounds()
		if len(line) > int(columnBound) {
			grid.AddColumn(uint(len(line)) - columnBound)
		}
	}

	for x, line := range data {
		for y, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				lib.Perror(err.Error())
				return "null"
			}
			err = grid.SetIndex(uint(x), uint(y), num)
			if err != nil {
				lib.Perror(err.Error())
				return "null"
			}
		}
	}

	var trees []tree

	r, c := grid.GetBounds()
	var x, y uint
	for x = 0; x < r; x += 1 {
		for y = 0; y < c; y += 1 {
			v, err := grid.GetIndex(x, y)
			if err != nil {
				lib.Perror(err.Error())
				return "null"
			}
			t := tree{x: x, y: y, v: v}
			row, err := grid.Row(x)
			if err != nil {
				lib.Perror(err.Error())
				return "null"
			}
			column, err := grid.Column(y)
			if err != nil {
				lib.Perror(err.Error())
				return "null"
			}
			t.up = row[:y]
			t.down = row[y+1:]
			t.left = column[:x]
			t.right = column[x+1:]

			trees = append(trees, t)
		}
	}

	// grid.Print()
	count := 0
	for _, t := range trees {
		if t.IsVisible() {
			count += 1
		}
		// t.Print()
	}

	return strconv.Itoa(count)
}

func part2(data []string) string {
	grid := lib.NewGrid()

	grid.AddRow(uint(len(data)))

	for _, line := range data {
		_, columnBound := grid.GetBounds()
		if len(line) > int(columnBound) {
			grid.AddColumn(uint(len(line)) - columnBound)
		}
	}

	for x, line := range data {
		for y, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				lib.Perror(err.Error())
				return "null"
			}
			err = grid.SetIndex(uint(x), uint(y), num)
			if err != nil {
				lib.Perror(err.Error())
				return "null"
			}
		}
	}

	var trees []tree

	r, c := grid.GetBounds()
	var x, y uint
	for x = 0; x < r; x += 1 {
		for y = 0; y < c; y += 1 {
			v, err := grid.GetIndex(x, y)
			if err != nil {
				lib.Perror(err.Error())
				return "null"
			}
			t := tree{x: x, y: y, v: v}
			row, err := grid.Row(x)
			if err != nil {
				lib.Perror(err.Error())
				return "null"
			}
			column, err := grid.Column(y)
			if err != nil {
				lib.Perror(err.Error())
				return "null"
			}
			t.up = row[:y]
			t.down = row[y+1:]
			t.left = column[:x]
			t.right = column[x+1:]

			trees = append(trees, t)
		}
	}

	winner := 0
	for _, t := range trees {
		vd := t.ViewingDistance()
		if vd > winner {
			winner = vd
		}
	}

	return strconv.Itoa(winner)
}

func Run(datafile string, part int) {
	lib.Print("Day 08 is ~n~o~t~~y~e~t~ implemented!")

	data, err := lib.ReadInTextFile(datafile)
	// data, err := lib.ReadInTextFile("day08/testdata.txt")
	if err != nil {
		lib.Perror("%v", err)
		return
	}

	switch part {
	case 1:
		lib.Print("Part 1: %v", part1(data))
	case 2:
		lib.Print("Part 2: %v", part2(data))
	default:
		lib.Perror("Unregognized part number: %v", part)
	}
}
