package day10

import (
	e "adventcode/2022/lib"
	"strconv"
	"strings"
)

//10 part 1: 11826 is too high
//10 part 1: 11720 is the correct answer

func filereport(cycle int, x int) {
	// return
	e.Print("%v: %v (%v)", cycle, x, cycle*x)
}

func part1(data []string) string {
	nstr := "null"
	x := 1
	cycle := 0
	var report []int

	for lineno, line := range data {
		instr := strings.Split(line, " ")
		if len(instr) < 1 {
			e.Perror("improper instruction: [%v]%v", lineno, line)
			return nstr
		}

		if !(instr[0] == "noop" && len(instr) == 1) &&
			!(instr[0] == "addx" && len(instr) == 2) {
			e.Perror("improper instruction: [%v]%v", lineno, line)
			return nstr
		}

		switch instr[0] {
		case "noop":
			cycle += 1
			report = append(report, cycle*x)
			filereport(cycle, x)
		case "addx":
			val, err := strconv.Atoi(instr[1])
			if err != nil {
				e.Perror("bad input: %v", line)
				return nstr
			}
			cycle += 1
			report = append(report, cycle*x)
			filereport(cycle, x)
			cycle += 1
			report = append(report, cycle*x)
			filereport(cycle, x)
			x += val
		}
	}
	diag := []int{19, 59, 99, 139, 179, 219}
	sum := 0

	for _, d := range diag {
		e.Print("[%v] %v", d, report[d])
		sum += report[d]
	}

	return strconv.Itoa(sum)
}

func setSpritePos(x int) []rune {
	spritestr := "........................................"
	var sprite []rune

	for i, char := range spritestr {
		if i >= x-1 && i <= x+1 {
			sprite = append(sprite, '#')
		} else {
			sprite = append(sprite, char)
		}
	}
	return sprite
}

func spriteToString(sprite []rune) string {
	var spritestr string

	for _, r := range sprite {
		if r == '.' {
			r = ' '
		}
		spritestr += string(r)
	}
	return spritestr
}

func part2(data []string) string {
	nstr := "null"
	var crtrows []string

	// crtrow := ""
	cycle := 0
	position := 0
	x := 1
	var sprite []rune
	var currentRow []rune

	for lineno, line := range data {
		instr := strings.Split(line, " ")
		if len(instr) < 1 {
			e.Perror("improper instruction: [%v]%v", lineno, line)
			return nstr
		}

		if !(instr[0] == "noop" && len(instr) == 1) &&
			!(instr[0] == "addx" && len(instr) == 2) {
			e.Perror("improper instruction: [%v]%v", lineno, line)
			return nstr
		}

		sprite = setSpritePos(x)

		switch instr[0] {
		case "noop":
			cycle += 1
			// e.Print("Sprite position : %v\n", spriteToString(sprite))
			// e.Print("Start cycle  %3v: begin executing %v", cycle, line)
			// e.Print("During cycle %3v: CRT draws pixel in position %v", cycle, position)
			currentRow = append(currentRow, sprite[position])
			// e.Print("Current CRT row : %v", spriteToString(currentRow))
			position += 1
			if position >= 40 {
				position = 0
				crtrows = append(crtrows, spriteToString(currentRow))
				currentRow = []rune{}
			}
			// e.Print("End of cycle %3v: finish executing %v (Register X is now %v)", cycle, line, x)
		case "addx":
			val, err := strconv.Atoi(instr[1])
			if err != nil {
				e.Perror("bad input: %v", line)
				return nstr
			}
			cycle += 1
			// e.Print("Sprite position : %v\n", spriteToString(sprite))
			// e.Print("Start cycle  %3v: begin executing %v", cycle, line)
			// e.Print("During cycle %3v: CRT draws pixel in position %v", cycle, position)
			currentRow = append(currentRow, sprite[position])
			// e.Print("Current CRT row : %v", spriteToString(currentRow))
			position += 1
			if position >= 40 {
				position = 0
				crtrows = append(crtrows, spriteToString(currentRow))
				currentRow = []rune{}
			}

			cycle += 1
			// e.Print("\nDuring cycle %3v: CRT draws pixel in position %v", cycle, position)
			currentRow = append(currentRow, sprite[position])
			// e.Print("Current CRT row : %v", spriteToString(currentRow))
			position += 1
			if position >= 40 {
				position = 0
				crtrows = append(crtrows, spriteToString(currentRow))
				currentRow = []rune{}
			}
			x += val
			// e.Print("End of cycle %3v: finish executing %v (Register X is now %v)", cycle, line, x)
		}

	}
	e.Print("")
	for _, row := range crtrows {
		e.Print("   %v", row)
	}
	e.Print("")
	// e.Print("%v", crtrows)

	return nstr

}

func Run(datafile string, part int) {
	e.Print("Day 10 is happening-ish!\n")

	data, err := e.ReadInTextFile(datafile)
	// data, err := e.ReadInTextFile("day10/testdata.txt")
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
