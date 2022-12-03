package day02

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	e "adventcode/2022/lib"
)

var resMap map[string]int

func getSingleResult(line string) (int, int, error) {
	if _, ok := resMap[line]; !ok {
		return 0, 0, fmt.Errorf("invalid line: /%v/", line)
	}

	return 0, resMap[line], nil
}

func makeMapRound1() {
	resMap = make(map[string]int)

	resMap["A X"] = 3 + 1 // Rock vs. rock
	resMap["A Y"] = 6 + 2 // Rock vs. paper
	resMap["A Z"] = 0 + 3 // Rock vs. scissors
	resMap["B X"] = 0 + 1 // Paper vs. rock
	resMap["B Y"] = 3 + 2 // Paper vs. paper
	resMap["B Z"] = 6 + 3 // Paper vs. scissors
	resMap["C X"] = 6 + 1 // Scissors vs. rock
	resMap["C Y"] = 0 + 2 // Scissors vs. paper
	resMap["C Z"] = 3 + 3 // Scissors vs. scissors
}

func makeMapRound2() {
	resMap = make(map[string]int)

	resMap["A X"] = 0 + 3 // Rock vs. lose (scissors)
	resMap["A Y"] = 3 + 1 // Rock vs. draw (rock)
	resMap["A Z"] = 6 + 2 // Rock vs.  win (paper)
	resMap["B X"] = 0 + 1 // Paper vs. lose (rock)
	resMap["B Y"] = 3 + 2 // Paper vs. draw (paper)
	resMap["B Z"] = 6 + 3 // Paper vs. win (scissors)
	resMap["C X"] = 0 + 2 // Scissors vs. lose (paper)
	resMap["C Y"] = 3 + 3 // Scissors vs. draw (scissors)
	resMap["C Z"] = 6 + 1 // Scissors vs. win (rock)
}

func Run(datafile string) {
	makeMapRound2()

	f, err := os.Open(datafile)
	if err != nil {
		e.Perror("Unable to open data file: %v", err)
		return
	}
	defer f.Close()

	totalUserScore := 0

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " \n")
		elf, user, err := getSingleResult(line)
		if err != nil {
			e.Perror("%v", err)
			time.Sleep(1 * time.Second)
		}
		totalUserScore += user
		e.Print("Elf: %v | User: %v/%v", elf, user, totalUserScore)

	}
}
