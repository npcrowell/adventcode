package day05

import (
	e "adventcode/2022/lib"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func initStacks(data []string) (map[string]e.Queue, []string, error) {
	stacks := make(map[string]e.Queue)
	divline := -1
	for i, line := range data {
		if line == "" {
			divline = i
			break
		}
	}
	if divline == -1 {
		return nil, data, fmt.Errorf("failed to find divider line in file")
	}

	stackdata := data[:divline]

	queuenums := strings.Split(stackdata[len(stackdata)-1], " ")
	for _, num := range queuenums {
		stacks[num] = e.Queue{}
	}

	// e.Print("'%v'", queuenums)

	for j, line := range stackdata {
		if j >= divline-1 {
			continue
		}
		for i := 0; i < len(line); i += 4 {
			queuenum := string(stackdata[divline-1][i])
			// e.Printf("'%v'", queuenum)
			next := line[i : i+3]
			val := strings.Trim(next, " \n\t[]")
			if strings.ReplaceAll(val, " ", "") == "" {
				continue
			}
			stacks[queuenum] = stacks[queuenum].PreInsert(val)
		}
	}
	delete(stacks, "")

	// for label := range stacks {
	// 	e.Printf("[%v]", label)
	// }
	// e.Print("")

	return stacks, data[divline+1:], nil
}

// func main() {
// 	matched, err := regexp.Match(`foo.*`, []byte(`seafood`))
// 	fmt.Println(matched, err)
// 	matched, err = regexp.Match(`bar.*`, []byte(`seafood`))
// 	fmt.Println(matched, err)
// 	matched, err = regexp.Match(`a(b`, []byte(`seafood`))
// 	fmt.Println(matched, err)

// }

func getInstruction(line string, reg *regexp.Regexp) (int, string, string, error) {
	matches := reg.FindAllStringSubmatch(line, -1)
	if len(matches) != 1 {
		return -1, "", "", fmt.Errorf("no (%v) matches found for '%v'", len(matches), line)
	}

	move, err := strconv.Atoi(matches[0][1])
	if err != nil {
		return -1, "", "", err
	}

	from := matches[0][2]
	to := matches[0][3]

	return move, from, to, nil
}

func getSortedKeys(stacks map[string]e.Queue) []string {
	var keys []string

	for k := range stacks {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func printStacks(stacks map[string]e.Queue) {
	keys := getSortedKeys(stacks)

	for _, key := range keys {
		e.Printf("%2v: %v\n", key, stacks[key])
	}
}

func part1(data []string) string {

	stacks, data, err := initStacks(data)
	if err != nil {
		e.Perror("%v", err)
		return "null"
	}
	printStacks(stacks)
	e.Print(".....")
	// time.Sleep(1 * time.Minute)

	reg, err := regexp.CompilePOSIX("move ([0-9]+) from ([0-9]+) to ([0-9]+)")
	if err != nil {
		e.Perror("Unable to compile regex pattern: %v", err)
	}

	for _, line := range data {
		move, from, to, err := getInstruction(line, reg)
		if err != nil {
			e.Perror("%v", err)
			return ""
		}

		for i := 0; i < move; i += 1 {
			item, fromqueue, err := stacks[from].Pop()
			stacks[from] = fromqueue
			if err != nil {
				e.Perror("%v", err)
				return ""
			}
			stacks[to] = stacks[to].Push(item)
		}
		// e.Print("%v", line)
		// printStacks(stacks)
	}

	res := ""
	keys := getSortedKeys(stacks)
	for _, key := range keys {
		q := stacks[key]
		item, _, err := q.Pop()
		if err != nil {
			e.Perror("%v", err)
		}
		res += item
	}
	return res

}

func part2(data []string) string {

	stacks, data, err := initStacks(data)
	if err != nil {
		e.Perror("%v", err)
		return "null"
	}
	printStacks(stacks)
	e.Print(".....")

	reg, err := regexp.CompilePOSIX("move ([0-9]+) from ([0-9]+) to ([0-9]+)")
	if err != nil {
		e.Perror("Unable to compile regex pattern: %v", err)
	}

	for _, line := range data {
		move, from, to, err := getInstruction(line, reg)
		if err != nil {
			e.Perror("%v", err)
			return ""
		}

		items, fromqueue, err := stacks[from].PopPlus(move)
		stacks[from] = fromqueue
		if err != nil {
			e.Perror("%v", err)
			return ""
		}
		stacks[to] = stacks[to].PushPlus(items)
		// e.Print("%v", line)
		// printStacks(stacks)
	}

	res := ""
	keys := getSortedKeys(stacks)
	for _, key := range keys {
		q := stacks[key]
		item, _, err := q.Pop()
		if err != nil {
			e.Perror("%v", err)
		}
		res += item
	}
	return res
}

func Run(datafile string) {
	e.Print("Let's Go Day 05!")

	data, err := e.ReadInTextFile(datafile)
	if err != nil {
		e.Perror("%v", err)
		return
	}

	e.Print("Part 1: %v", part1(data))

	e.Print("Part 2: %v", part2(data))

}
