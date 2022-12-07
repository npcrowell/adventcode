package main

import (
	"adventcode/2021/days"
	"adventcode/2021/lib"
)

func main() {
	data, err := lib.ReadInTextFile("data/d03.txt")
	if err != nil {
		lib.Perror("%v", err)
		return
	}
	days.Run03(data)
}
