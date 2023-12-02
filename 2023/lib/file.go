package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadInTextFile(filename string) ([]string, error) {

	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to open file (%v): %v", filename, err)
	}
	defer f.Close()
	var ret []string
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " \n")
		ret = append(ret, line)
	}
	return ret, nil
}

func RemoveBlankLines(data []string) []string {
	var outData []string

	for _, d := range data {
		if strings.Trim(d, " \n\t\r") != "" {
			outData = append(outData, d)
		}
	}

	return outData
}
