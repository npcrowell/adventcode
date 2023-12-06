package days

import (
	"advent2023/lib"
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var (
	seedline     *regexp.Regexp
	mapstartline *regexp.Regexp
	numline      *regexp.Regexp
)

func initRegex() {
	sl, err := regexp.Compile(`^seeds: \d+ `)
	if err != nil {
		lib.Perror("Unable to compile seedline: %v", err)
		return
	}
	msl, err := regexp.Compile(`^(\w+)-to-(\w+) map:`)
	if err != nil {
		lib.Perror("Unable to compile mapstartline: %v", err)
		return
	}
	nl, err := regexp.Compile(`^\d+ +`)
	if err != nil {
		lib.Perror("Unable to compile seedline: %v", err)
		return
	}
	seedline = sl
	mapstartline = msl
	numline = nl
}

func day5_parseSeeds(line string) ([]int, error) {
	var snums []int
	for _, sstr := range strings.Split(line, " ")[1:] {
		snum, err := strconv.Atoi(sstr)
		if err != nil {
			lib.Perror("Unable to convert '%v' to int", sstr)
			return nil, err
		}
		snums = append(snums, snum)
	}
	return snums, nil
}

func day5_parseSeeds2(line string) ([]int, error) {
	var snums []int
	for _, sstr := range strings.Split(line, " ")[1:] {
		snum, err := strconv.Atoi(sstr)
		if err != nil {
			lib.Perror("Unable to convert '%v' to int", sstr)
			return nil, err
		}
		snums = append(snums, snum)
	}

	var srnums []int
	for i := 0; i < len(snums); i += 2 {
		lib.Debug("Seed: %v", snums[i])
		for snum := snums[i]; snum < snums[i]+snums[i+1]; snum += 1 {
			srnums = append(srnums, snum)
		}
	}
	return srnums, nil
}

func day5_parseSingleSeed(lines []string, sval int) (string, string, int, int, error) {
	// lib.Debug(lines[0])
	mapst := ""
	mapdst := ""
	if mapstart := mapstartline.FindStringSubmatch(lines[0]); mapstart != nil {
		// lib.Debug("Found mapstartline: %v to %v", mapstart[1], mapstart[2])
		mapst = mapstart[1]
		mapdst = mapstart[2]
	}

	count := 1
	sdval := 0

	for c, line := range lines[1:] {
		// lib.Debug(line)
		if len(line) == 0 {
			count = c + 2
			break
		}
		snums := strings.Split(line, " ")
		dest, err := strconv.Atoi(snums[0])
		if err != nil {
			lib.Perror("Unable to convert '%v' to int", snums)
			return mapst, mapdst, sval, count, err
		}
		start, err := strconv.Atoi(snums[1])
		if err != nil {
			lib.Perror("Unable to convert '%v' to int", snums)
			return mapst, mapdst, sval, count, err
		}
		rlen, err := strconv.Atoi(snums[2])
		if err != nil {
			lib.Perror("Unable to convert '%v' to int", snums)
			return mapst, mapdst, sval, count, err
		}
		mod := dest - start

		// lib.Debug("(%v,%v,%v) start: %v, end: %v, diff: %v", start, dest, rlen, start, start+rlen, mod)

		if sdval == 0 {
			if sval >= start && sval < (start+rlen) {
				lib.Debug("%v->%v", sval, sval+mod)
				sdval = sval + mod
			}
		}
		count = c + 2
	}

	if sdval == 0 {
		sdval = sval
	}
	// lib.Debug("Returning %v", count)
	return mapst, mapdst, sdval, count, nil
}

func day5_parseNextMapping(lines []string, svals []int) (string, string, []int, int, error) {
	// lib.Debug(lines[0])
	mapst := ""
	mapdst := ""
	if mapstart := mapstartline.FindStringSubmatch(lines[0]); mapstart != nil {
		// lib.Debug("Found mapstartline: %v to %v", mapstart[1], mapstart[2])
		mapst = mapstart[1]
		mapdst = mapstart[2]
	}

	count := 1
	var sdvals []int
	for range svals {
		sdvals = append(sdvals, 0)
	}

	for c, line := range lines[1:] {
		// lib.Debug(line)
		if len(line) == 0 {
			count = c + 2
			break
		}
		snums := strings.Split(line, " ")
		dest, err := strconv.Atoi(snums[0])
		if err != nil {
			lib.Perror("Unable to convert '%v' to int", snums)
			return mapst, mapdst, svals, count, err
		}
		start, err := strconv.Atoi(snums[1])
		if err != nil {
			lib.Perror("Unable to convert '%v' to int", snums)
			return mapst, mapdst, svals, count, err
		}
		rlen, err := strconv.Atoi(snums[2])
		if err != nil {
			lib.Perror("Unable to convert '%v' to int", snums)
			return mapst, mapdst, svals, count, err
		}
		mod := dest - start

		// lib.Debug("(%v,%v,%v) start: %v, end: %v, diff: %v", start, dest, rlen, start, start+rlen, mod)

		for i, v := range sdvals {
			if v == 0 {
				if svals[i] >= start && svals[i] < (start+rlen) {
					lib.Debug("%v->%v", svals[i], svals[i]+mod)
					sdvals[i] = svals[i] + mod
				}
			}
		}
		count = c + 2
	}

	for i, v := range sdvals {
		if v == 0 {
			sdvals[i] = svals[i]
		}
	}
	// lib.Debug("Returning %v", count)
	return mapst, mapdst, sdvals, count, nil
}

func day5_CompileMaps(lines []string) [][][3]int {
	var maps [][][3]int
	var m [][3]int

	for _, line := range lines[1:] {

		if len(line) == 0 {
			maps = append(maps, m)
			m = nil
			continue
		} else if mapstart := mapstartline.FindStringSubmatch(line); mapstart != nil {
			continue
		}

		snums := strings.Split(line, " ")
		dest, _ := strconv.Atoi(snums[0])
		start, _ := strconv.Atoi(snums[1])
		rlen, _ := strconv.Atoi(snums[2])
		mod := dest - start
		m = append(m, [3]int{start, start + rlen, mod})
	}
	maps = append(maps, m)
	for _, m := range maps {
		lib.Debug("%v", m)
	}
	return maps
}

func day5_checkSeedAgainst(seed int, maps [][][3]int) int {
	startseed := seed
	dbgstr := fmt.Sprintf("%v", startseed)
	for _, t := range maps {
		dbgstr += "->"
		for _, m := range t {
			if m[0] <= seed && seed < m[1] {
				// lib.Printf("%v <= %v <= %v;  ", m[0], seed, m[1])
				seed += m[2]
				break
			}
		}
		dbgstr += fmt.Sprintf("%v", seed)
	}
	dbgstr += "  "
	lib.Debug("%v | %v->%v", dbgstr, startseed, seed)
	return seed
}

func day5_doSeedPool(seedstart int, seedend int, cmaps [][][3]int, minout chan<- int) { //, status chan<- int) {
	lib.Print("Starting seed pool %12v (%12v)", seedstart, seedend-seedstart)
	tenpct := ((seedend - seedstart) / 10) + 1
	min := math.MaxInt
	count := 0
	dcount := 0
	for snum := seedstart; snum < seedend; snum += 1 {
		count += 1
		if count%tenpct == 0 {
			dcount += 10
			lib.Printf("[%12v] %v%%\n", seedstart, dcount)
			// status <- dcount
		}
		min = lib.Min(min, day5_checkSeedAgainst(snum, cmaps))
	}
	lib.Psuccess("[%12v (%12v)] %12v", seedstart, seedend-seedstart, min)
	minout <- min
}

func Day05(part int, data []string) (int, error) {
	initRegex()
	cmaps := day5_CompileMaps(data[1:])
	switch part {
	case 1:
		sv, err := day5_parseSeeds(data[0])
		if err != nil {
			return 0, err
		}
		lib.Printf("Seeds (%v): %v\n", len(sv), sv)
		index := 2
		for index < len(data) {
			m1, m2, s, i, err := day5_parseNextMapping(data[index:], sv)
			if err != nil {
				return 0, nil
			}
			index += i
			sv = s

			lib.Printf("%15v->%15v: %v\n", m1, m2, sv)
			if m2 == "location" {
				break
			}
		}

		minseedval := lib.Min(sv[0], sv[1:]...)

		return minseedval, nil
	case 2:
		min := math.MaxInt
		var snums []int
		for _, sstr := range strings.Split(data[0], " ")[1:] {
			snum, err := strconv.Atoi(sstr)
			if err != nil {
				lib.Perror("Unable to convert '%v' to int", sstr)
				return 0, err
			}
			snums = append(snums, snum)
		}

		var outs []chan int
		// var status []chan int
		for i := 0; i < len(snums); i += 2 {
			ochan := make(chan int, 1)
			// stat := make(chan int, 10)
			outs = append(outs, ochan)
			// status = append(status, stat)
			start := snums[i]
			end := snums[i] + snums[i+1]
			go day5_doSeedPool(start, end, cmaps, ochan) //, stat)
		}

		for i := 0; i < len(outs); i += 1 {
			min = lib.Min(min, <-outs[i])
		}
		return min, nil
	default:
		return 0, errors.New("not implemented")
	}
}
