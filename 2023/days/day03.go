package days

import (
	"advent2023/lib"
	"advent2023/lib/grid"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func day3_getSum(lines []string, findex int) (int, error) {
	numptn, err := regexp.Compile(`\D?(\d+)\D?`)
	if err != nil {
		return 0, err
	}
	tline := lines[findex]

	mindices := numptn.FindAllStringSubmatchIndex(tline, -1)

	sum := 0
	for _, loc := range mindices {
		num, _ := strconv.Atoi(tline[loc[2]:loc[3]])

		sindex := lib.Max(loc[2]-1, 0)
		lindex := lib.Min(loc[3]+1, len(tline))

		box := ""
		for _, line := range lines {
			box += line[sindex:lindex]
		}

		box = strings.ReplaceAll(box, ".", "")
		symptn, err := regexp.Compile(`\D`)
		if err != nil {
			return 0, err
		}

		res := symptn.FindString(box)

		if len(res) > 0 {
			sum += num
		}
	}

	return sum, nil
}

func containsInt(a []int, b int) bool {
	for _, c := range a {
		if b == c {
			return true
		}
	}
	return false
}

func isnum(sval string) bool {
	if _, err := lib.StringToDigit(sval, 10); err != nil {
		return false
	}
	return true
}

func day3_getAssNumsAmong(g *grid.Grid, indices []grid.T) []int {
	var anums []int
	var cnums []int

	for _, i := range indices {
		index := i.(int)
		if containsInt(cnums, index) {
			continue
		}
		cnums = append(cnums, index)

		sval := g.Value(index).(string)
		if !isnum(sval) {
			continue
		}

		sindex := index
		for ic := index; true; ic = g.Left(ic) {
			value := g.Value(ic).(string)
			if !isnum(value) {
				break
			}

			sindex = ic

			if g.IsFirstColumn(ic) {
				break
			}
		}
		nstr := ""

		for ic := sindex; true; ic = g.Right(ic) {
			value := g.Value(ic).(string)
			if !isnum(value) {
				break
			}

			nstr += value
			if !containsInt(cnums, ic) {
				cnums = append(cnums, ic)
			}

			if g.IsLastColumn(ic) {
				break
			}
		}

		val, err := strconv.Atoi(nstr)
		if err != nil {
			lib.Debug("Error converting %v to int", nstr)
			return []int{}
		}
		anums = append(anums, val)
	}

	return anums
}

func day3_handleAsGrid(g *grid.Grid) (int, error) {
	lib.Debug("\n%v\n", g)

	matchcoords := g.Search("*", g.Index(0, 0), -1, -1)
	ratios := 0

	for _, index := range matchcoords {
		radius := g.GetRadius(index, true)

		anums := day3_getAssNumsAmong(g, radius)
		if len(anums) == 2 {
			gearratio := anums[0] * anums[1]
			lib.Debug("%v * %v = %v", anums[0], anums[1], gearratio)
			ratios += gearratio
		}
	}

	return ratios, nil
}

func Day03(part int, data []string) (int, error) {
	switch part {
	case 1:
		sum := 0
		for mid := 0; mid < len(data); mid += 1 {
			min := int(lib.Max(0, mid-1))
			max := int(lib.Min(len(data)-1, mid+1))
			tsum, err := day3_getSum(data[min:max+1], mid-min)
			if err != nil {
				return sum, err
			}
			sum += tsum
		}

		return sum, nil
	case 2:
		g, err := lib.ConvertStringArrayToGrid(data)
		if err != nil {
			return 0, err
		}
		sum, err := day3_handleAsGrid(g)
		return sum, err
	default:
		return 0, errors.New("not implemented")
	}
}
