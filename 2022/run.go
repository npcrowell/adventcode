package main

import (
	d01 "adventcode/2022/day01"
	d02 "adventcode/2022/day02"
	d03 "adventcode/2022/day03"
	d04 "adventcode/2022/day04"
	d05 "adventcode/2022/day05"
	d06 "adventcode/2022/day06"
	d07 "adventcode/2022/day07"
	d08 "adventcode/2022/day08"
	d09 "adventcode/2022/day09"
	d10 "adventcode/2022/day10"
	d11 "adventcode/2022/day11"
	d12 "adventcode/2022/day12"
	d13 "adventcode/2022/day13"
	d14 "adventcode/2022/day14"
	d15 "adventcode/2022/day15"
	d16 "adventcode/2022/day16"
	d17 "adventcode/2022/day17"
	d18 "adventcode/2022/day18"
	d19 "adventcode/2022/day19"
	d20 "adventcode/2022/day20"
	d21 "adventcode/2022/day21"
	d22 "adventcode/2022/day22"
	d23 "adventcode/2022/day23"
	d24 "adventcode/2022/day24"
	d25 "adventcode/2022/day25"
	"fmt"
)

func run(day int, part int, datafile string) {
	fmt.Printf("This is the day! (%v)\n", day)
	switch day {
	case 1:
		d01.Run()
	case 2:
		d02.Run(datafile)
	case 3:
		d03.Run(datafile)
	case 4:
		d04.Run(datafile)
	case 5:
		d05.Run(datafile, part)
	case 6:
		d06.Run(datafile, part)
	case 7:
		d07.Run(datafile, part)
	case 8:
		d08.Run(datafile, part)
	case 9:
		d09.Run(datafile, part)
	case 10:
		d10.Run(datafile, part)
	case 11:
		d11.Run(datafile, part)
	case 12:
		d12.Run(datafile, part)
	case 13:
		d13.Run(datafile, part)
	case 14:
		d14.Run(datafile, part)
	case 15:
		d15.Run(datafile, part)
	case 16:
		d16.Run(datafile, part)
	case 17:
		d17.Run(datafile, part)
	case 18:
		d18.Run(datafile, part)
	case 19:
		d19.Run(datafile, part)
	case 20:
		d20.Run(datafile, part)
	case 21:
		d21.Run(datafile, part)
	case 22:
		d22.Run(datafile, part)
	case 23:
		d23.Run(datafile, part)
	case 24:
		d24.Run(datafile, part)
	case 25:
		d25.Run(datafile, part)
	}
}
