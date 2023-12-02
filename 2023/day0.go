package main

import "errors"

// Template file

var d0p1_testdata []string = []string{
	"empty",
}

var d0p2_testdata []string = []string{
	"empty",
}

func day0(part int, data []string, useTestData bool) (int, error) {
	switch part {
	case 1:
		if useTestData {
			data = d0p1_testdata
		}
		return 0, nil
	case 2:
		if useTestData {
			data = d0p2_testdata
		}
		return 0, nil
	default:
		return 0, errors.New("not implemented")
	}
}
