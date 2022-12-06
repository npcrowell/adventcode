package main

import (
	d01 "adventcode/2022/day01"
	d02 "adventcode/2022/day02"
	d03 "adventcode/2022/day03"
	d04 "adventcode/2022/day04"
	d05 "adventcode/2022/day05"
	d06 "adventcode/2022/day06"
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
	}
}
