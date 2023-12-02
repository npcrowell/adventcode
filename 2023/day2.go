package main

import (
	"errors"
)

// Template file

var d2p1_testdata []string = []string{
	"abcde",
	"fghij",
	"klmno",
	"pqrst",
	"uvwxy",
}

var d2p2_testdata []string = []string{
	"empty",
}

func day2(part int, data []string) (int, error) {

	switch part {
	case 1:
		return 0, nil
	case 2:
		return 0, nil
	default:
		return 0, errors.New("not implemented")
	}
}
