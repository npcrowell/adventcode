package lib

import (
	"errors"
	"strconv"
	"strings"
)

func GetSubstrings(base string, searchsubs []string, caseSensitive bool) []string {
	var foundsubs []string
	var snap int

	for i := range base {
		for _, sub := range searchsubs {
			if snap = i + len(sub); snap > len(base) {
				break
			}
			bsub := base[i:snap]

			if caseSensitive {
				if bsub == sub {
					foundsubs = append(foundsubs, bsub)
				}
			} else {
				if strings.EqualFold(bsub, sub) {
					foundsubs = append(foundsubs, bsub)
				}
			}
		}
	}

	return foundsubs
}

func StringToDigit(digit string, base int) (int, error) {
	digits := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	switch base {
	case 10:
		val, err := strconv.Atoi(digit)
		if err == nil {
			return val, nil
		}
		if val, ok := digits[digit]; ok {
			return val, nil
		}
		return 0, errors.New("digit not recognized in base ten")
	case 16:
		return 0, errors.New("base 16 is not yet supported")
	default:
		return 0, errors.New("unrecognized base for conversion")
	}
}

func Max(a int, b ...int) int {
	max := a
	for _, i := range b {
		// Debug("max %v > %v ?", i, max)
		if i > max {
			max = i
		}
	}
	// Debug("max =  %v", max)
	return max
}

func Min(a int, b ...int) int {
	min := a
	for _, i := range b {
		// Debug("min %v < %v ?", i, min)
		if i < min {
			min = i
		}
	}
	// Debug("min =  %v", min)
	return min
}
