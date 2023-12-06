package days

import (
	"advent2023/lib"
	"errors"
	"strconv"
	"strings"
)

func day6_doRace(time int, distance int) int {
	lib.Debug("Time: %v | Distance: %v", time, distance)
	mid := time / 2

	i := mid
	for {
		traveltime := time - i // travel time
		totaldist := i * traveltime
		// lib.Debug("%v <= %v || %v <= 0", totaldist, distance, i)
		if totaldist <= distance || i <= 0 {
			lib.Debug("low: %v | totaldist: %v", i, totaldist)
			// lib.Debug("Breaking")
			break
		}
		i -= 1
	}
	low := i
	i = mid
	for {
		traveltime := time - i
		totaldist := i * traveltime
		// lib.Debug("%v <= %v || %v >= %v", totaldist, distance, i, time)
		if totaldist <= distance || i >= time {
			lib.Debug("high: %v | totaldist: %v", i, totaldist)
			// lib.Debug("Breaking")
			break
		}
		i += 1
	}
	high := i

	return high - low - 1
}

func day6_parsedata(data []string) ([]int, []int) {
	var times, dists []int
	for i, line := range data {
		nums := strings.Split(line, " ")
		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err == nil {
				switch i {
				case 0:
					times = append(times, n)
				case 1:
					dists = append(dists, n)
				default:
					lib.Perror("Unexpected number of lines: %v", i)
				}
			}
		}
	}
	return times, dists
}

func day6_parsedata2(data []string) (int, int) {
	var time, dist int
	for i, line := range data {
		num := strings.Split(strings.ReplaceAll(line, " ", ""), ":")[1]
		n, err := strconv.Atoi(num)
		if err == nil {
			switch i {
			case 0:
				time = n
			case 1:
				dist = n
			default:
				lib.Perror("Unexpected number of lines: %v", i)
			}
		}
	}
	return time, dist
}

func Day06(part int, data []string) (int, error) {

	switch part {
	case 1:
		times, dists := day6_parsedata(data)
		prod := 1
		for i := range times {
			count := day6_doRace(times[i], dists[i])
			lib.Debug("Race %v: %v\n", i, count)
			prod *= count
		}
		return prod, nil
	case 2:
		time, dist := day6_parsedata2(data)
		lib.Debug("Time: %v | Distance: %v", time, dist)
		count := day6_doRace(time, dist)
		return count, nil
	default:
		return 0, errors.New("not implemented")
	}
}
