package days

import (
	"errors"
	"strconv"
	"strings"
)

func day01_sort(data []string) ([]int, []int, error) {
	var col1 []int
	var col2 []int

	for _, line := range data {
		cols := strings.Fields(line)

		_col1, err := strconv.Atoi(cols[0])
		if err != nil {
			return nil, nil, err
		}

		_col2, err := strconv.Atoi(cols[1])
		if err != nil {
			return nil, nil, err
		}

		col1 = append(col1, _col1)
		col2 = append(col2, _col2)
	}

	for i := 0; i < len(col1); i++ {
		for j := i + 1; j < len(col1); j++ {
			if col1[i] > col1[j] {
				col1[i], col1[j] = col1[j], col1[i]
			}
			if col2[i] > col2[j] {
				col2[i], col2[j] = col2[j], col2[i]
			}
		}
	}

	return col1, col2, nil
}

func day01_count_contains(val int, col []int) int {
	count := 0
	for _, v := range col {
		if v == val {
			count++
		}
	}
	return count
}

func day01_distance(col1 []int, col2 []int) int {
	dist := 0
	for i := 0; i < len(col1); i++ {
		if col1[i] > col2[i] {
			dist += col1[i] - col2[i]
		} else {
			dist += col2[i] - col1[i]
		}
	}
	return dist
}

func Day01(part int, data []string) (int, error) {
	col1, col2, err := day01_sort(data)
	if err != nil {
		return 0, err
	}

	switch part {
	case 1:
		// lib.Printf("Column 1: %v\n", col1)
		// lib.Printf("Column 2: %v\n", col2)
		return day01_distance(col1, col2), nil
	case 2:
		// lib.Printf("Column 1: %v\n", col1)
		// lib.Printf("Column 2: %v\n", col2)
		prods := 0
		for i := 0; i < len(col1); i++ {
			prods += day01_count_contains(col1[i], col2) * col1[i]
		}
		return prods, nil
	default:
		return 0, errors.New("not implemented")
	}
}
