package days

import (
	"advent2024/lib"
	"errors"
	"strconv"
	"strings"
)

func d2_parse(data []string) ([][]int, error) {
	var results [][]int

	for _, line := range data {
		//each line is a space delimited list of integers
		//split the line into a slice of integers
		//append the slice to a slice of slices

		result_str := strings.Fields(line)
		var result_int []int
		for _, r := range result_str {
			ri, err := strconv.Atoi(r)
			if err != nil {
				return nil, err
			}
			result_int = append(result_int, ri)
		}
		results = append(results, result_int)
	}
	return results, nil
}

func d2_is_safe(data []int) bool {
	// establish initial case
	if len(data) == 1 {
		return false
	}

	if data[0] > data[1] {
		for i, d := range data {
			if i == 0 {
				continue
			}
			if d >= data[i-1] || d < data[i-1]-3 {
				return false
			}
		}
	} else {
		for i, d := range data {
			if i == 0 {
				continue
			}
			if d <= data[i-1] || d > data[i-1]+3 {
				return false
			}
		}
	}

	return true
}

func d2_remove_int(data []int, index int) []int {
	//copy data to a new slice, ex=cluding the index
	var result []int
	for i, d := range data {
		if i == index {
			continue
		}
		result = append(result, d)
	}
	return result
}

func d2_count_safe(data [][]int) int {
	count := 0

loop1:
	for _, d := range data {

		for i := range d {
			d_tmp := d2_remove_int(d, i)
			// log.Printf("d_tmp: %v", d_tmp)
			if d2_is_safe(d_tmp) {
				lib.Debug("%v (safe)", d_tmp)
				count++
				continue loop1
			}
		}
		lib.Debug("%v (unsafe)", d)
	}

	return count
}

func Day02(part int, data []string) (int, error) {
	pdata, err := d2_parse(data)
	if err != nil {
		return 0, err
	}
	switch part {
	case 1:
		return d2_count_safe(pdata), nil

		// return 0, nil
	case 2:
		return d2_count_safe(pdata), nil
	default:
		return 0, errors.New("not implemented")
	}
}
