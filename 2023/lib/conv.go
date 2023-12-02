package lib

import (
	"errors"
	"strconv"
	"strings"
)

func GetSubstrings(base string, searchsubs []string, caseSensitive bool) []string {
	var foundsubs []string

	for i := range base {
		for _, sub := range searchsubs {
			snap := i + len(sub)
			if snap <= len(base) {
				// Printf("Searching for '%v' at '%v' [%v:%v] \n", sub, base[i:], i, snap)
				if caseSensitive {
					if base[i:snap] == sub {
						foundsubs = append(foundsubs, sub)
					}
				} else {
					if strings.EqualFold(base[i:snap], sub) {
						// Printf("Found '%v' in '%v'\n", sub, base)
						foundsubs = append(foundsubs, base[i:snap])
					}
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
