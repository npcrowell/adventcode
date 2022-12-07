package days

import (
	"adventcode/2021/lib"
	"fmt"
	"strconv"
)

func d3part1(data []string) string {
	var zerobitcount []int
	var onebitcount []int

	for i := range data[0] {
		if data[0][i] == '0' {
			zerobitcount = append(zerobitcount, 1)
			onebitcount = append(onebitcount, 0)
		} else {
			zerobitcount = append(zerobitcount, 0)
			onebitcount = append(onebitcount, 1)
		}
	}

	for _, line := range data[1:] {
		for i, bit := range line {
			if bit == '0' {
				zerobitcount[i] += 1
			} else {
				onebitcount[i] += 1
			}
		}
	}

	var num string
	var inum string

	for i := range zerobitcount {
		if zerobitcount[i] > onebitcount[i] {
			num += "0"
			inum += "1"
		} else {
			num += "1"
			inum += "0"
		}
	}

	res1, err := strconv.ParseInt(num, 2, 32)
	if err != nil {
		lib.Perror(err.Error())
		return "null"
	}

	res2, err := strconv.ParseInt(inum, 2, 32)
	if err != nil {
		lib.Perror(err.Error())
		return "null"
	}

	ret := strconv.FormatInt(res1*res2, 10)
	return ret
}

func getMostCommonBit(data []string, index int) string {
	zerobitcount := 0
	onebitcount := 0

	for _, line := range data {
		bit := line[index]
		if bit == '0' {
			zerobitcount += 1
		} else {
			onebitcount += 1
		}
	}
	if onebitcount >= zerobitcount {
		return "1"
	} else {
		return "0"
	}
}

func getLeastCommonBit(data []string, index int) string {
	zerobitcount := 0
	onebitcount := 0

	for _, line := range data {
		bit := line[index]
		if bit == '0' {
			zerobitcount += 1
		} else {
			onebitcount += 1
		}
	}
	if onebitcount >= zerobitcount {
		return "0"
	} else {
		return "1"
	}
}

func reducedata(data []string, index int, mcb string) []string {
	var tempdata []string

	for _, line := range data {
		if line[index] == mcb[0] {
			tempdata = append(tempdata, line)
		}
	}
	return tempdata
}

func dataToInt(data string) int {

	res, err := strconv.ParseInt(data, 2, 32)
	if err != nil {
		lib.Perror(err.Error())
	}

	return int(res)
}

func d3part2(data []string) string {
	tempdata := data

	index := 0
	for len(tempdata) > 1 {
		mcb := getMostCommonBit(tempdata, index)
		tempdata = reducedata(tempdata, index, mcb)
		index += 1
	}

	ogr := dataToInt(tempdata[0])

	// lib.Print(tempdata[0]) // oxygen generator rating

	tempdata = data
	index = 0
	for len(tempdata) > 1 {
		lcb := getLeastCommonBit(tempdata, index)
		tempdata = reducedata(tempdata, index, lcb)
		index += 1
	}

	// lib.Print(tempdata[0]) // co2 scrubber rating
	csr := dataToInt(tempdata[0])

	// lif
	lif := ogr * csr

	return strconv.Itoa(lif)
}

func Run03(dataset []string) {
	fmt.Println("Onward day 3!")
	lib.Print("Part 1: %v", d3part1(dataset))
	lib.Print("Part 2: %v", d3part2(dataset))
}
