package days

import (
	"advent2023/lib"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func day2_parseGame(line string) (int, []map[string]int, error) {
	var game []map[string]int
	idxptn, err := regexp.Compile(`Game (\d+):`)
	if err != nil {
		return 0, nil, err
	}

	index, err := strconv.Atoi(idxptn.FindStringSubmatch(line)[1])
	if err != nil {
		return 0, nil, err
	}

	gameline := strings.Split(line, ":")[1]
	games := strings.Split(gameline, ";")

	clrptn, err := regexp.Compile(`(\d+) (\w+)`)
	if err != nil {
		return 0, nil, err
	}

	for _, g := range games {
		tgame := make(map[string]int)

		for _, colordata := range strings.Split(g, ",") {
			match := clrptn.FindStringSubmatch(colordata)
			count := match[1]
			color := match[2]
			tgame[color], err = strconv.Atoi(count)
			if err != nil {
				return 0, nil, err
			}
			game = append(game, tgame)
		}
	}

	return index, game, nil
}

func day2_validateGame(line string, limits map[string]int) int {
	index, subgames, err := day2_parseGame(line)
	if err != nil {
		lib.Perror(err.Error())
		return 0
	}

	for _, game := range subgames {
		for color, gcount := range game {
			if lcount, ok := limits[color]; ok {
				if gcount > lcount {
					lib.Debug("ejected %v for %v: %v", index, color, game)
					return 0
				}
			}
		}
	}

	return index
}

func day2_computeGamePower(line string) int {
	_, subgames, err := day2_parseGame(line)
	if err != nil {
		lib.Perror(err.Error())
		return 0
	}

	minimums := map[string]int{
		"red":   0,
		"blue":  0,
		"green": 0,
	}

	for _, game := range subgames {
		for color, count := range game {
			if lcount, ok := minimums[color]; ok {
				if count > lcount {
					minimums[color] = count
				}
			}
		}
	}

	power := 1
	for _, count := range minimums {
		power *= count
	}

	lib.Debug("%v = %v", minimums, power)
	return power
}

func Day02(part int, data []string) (int, error) {

	switch part {
	case 1:
		sum := 0
		for _, line := range data {
			limits := map[string]int{
				"red":   12,
				"green": 13,
				"blue":  14,
			}
			sum += day2_validateGame(line, limits)
		}
		return sum, nil
	case 2:
		sum := 0
		for _, line := range data {
			power := day2_computeGamePower(line)
			sum += power
			// lib.Debug("%4v: %v", i+1, power)
		}
		return sum, nil
	default:
		return 0, errors.New("not implemented")
	}
}
