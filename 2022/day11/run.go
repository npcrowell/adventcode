package day11

import (
	e "adventcode/2022/lib"
	"fmt"
	"math/big"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// func getRegMatch(line string, regexp *regexp.Regexp) []error {
// 	return nil, nil
// }

func getSortedKeys(stacks map[string]string) []string {
	var keys []string

	for k := range stacks {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func getNextMonkey(data []string) (*Monkey, error) {
	reMonkeyId, err := regexp.CompilePOSIX("^Monkey ([0-9]+):")
	if err != nil {
		e.Perror("Unable to compile regex pattern: %v", err)
	}
	reStarting, err := regexp.CompilePOSIX("^Starting items: ([0-9, ]+)")
	if err != nil {
		e.Perror("Unable to compile regex pattern: %v", err)
	}
	reOperation, err := regexp.CompilePOSIX("^Operation: new = (.+)")
	if err != nil {
		e.Perror("Unable to compile regex pattern: %v", err)
	}
	reTest, err := regexp.CompilePOSIX("^Test: divisible by ([0-9]+)")
	if err != nil {
		e.Perror("Unable to compile regex pattern: %v", err)
	}
	reIfTrue, err := regexp.CompilePOSIX("^If true: throw to monkey ([0-9]+)")
	if err != nil {
		e.Perror("Unable to compile regex pattern: %v", err)
	}
	reIfFalse, err := regexp.CompilePOSIX("^If false: throw to monkey ([0-9]+)")
	if err != nil {
		e.Perror("Unable to compile regex pattern: %v", err)
	}

	// monkeyData := make(map[string]string)
	monkey := newMonkey()

	for _, line := range data {
		// e.Print(line)
		monkeyId := reMonkeyId.FindAllStringSubmatch(line, -1)
		if len(monkeyId) >= 1 {
			monkey.id, err = strconv.Atoi(monkeyId[0][1])
			if err != nil {
				return nil, fmt.Errorf(err.Error())
			}
			continue
		}

		starting := reStarting.FindAllStringSubmatch(line, -1)
		if len(starting) >= 1 {
			s := strings.Split(starting[0][1], ", ")
			for _, st := range s {
				sta, err := strconv.Atoi(st)
				if err != nil {
					return nil, fmt.Errorf(err.Error())
				}
				monkey.items = append(monkey.items, big.NewInt(int64(sta)))
			}
			// monkeyData["1 starting"] = starting[0][1]
			continue
		}

		operation := reOperation.FindAllStringSubmatch(line, -1)
		if len(operation) >= 1 {
			operands := strings.Split(operation[0][1], " ")
			monkey.operate = rune(operands[1][0])
			monkey.operand = operands[len(operands)-1]
			// monkeyData["2 operation"] = operation[0][1]
			continue
		}

		test := reTest.FindAllStringSubmatch(line, -1)
		if len(test) >= 1 {
			mod, err := strconv.Atoi(test[0][1])
			if err != nil {
				return nil, fmt.Errorf(err.Error())
			}

			monkey.modulus = uint64(mod)
			// monkeyData["3 test"] =
			continue
		}

		ifTrue := reIfTrue.FindAllStringSubmatch(line, -1)
		if len(ifTrue) >= 1 {
			mon, err := strconv.Atoi(ifTrue[0][1])
			if err != nil {
				return nil, fmt.Errorf(err.Error())
			}
			monkey.tnext = mon
			// monkeyData["4 true"] = ifTrue[0][1
			continue
		}
		ifFalse := reIfFalse.FindAllStringSubmatch(line, -1)
		if len(ifFalse) >= 1 {
			mon, err := strconv.Atoi(ifFalse[0][1])
			if err != nil {
				return nil, fmt.Errorf(err.Error())
			}
			monkey.fnext = mon
			// monkeyData["5 false"] = ifFalse[0][1]
			continue
		}
		e.Perror("No match for: %v", line)
	}

	// keys := getSortedKeys(monkeyData)

	// for _, key := range keys {
	// 	e.Printf("%2v: %v\n", key, monkeyData[key])
	// }
	return monkey, nil
}

func getMonkeys(data []string) ([]*Monkey, error) {
	var m []string
	var monkeys []*Monkey

	for _, line := range data {
		if line == "" {
			monkey, err := getNextMonkey(m)
			if err != nil {
				return nil, err
			}
			monkeys = append(monkeys, monkey)

			m = []string{}
		} else {
			m = append(m, line)
		}
	}
	return monkeys, nil
}

func getMagicNumber(monkeys []*Monkey) int {
	magic := 1
	for _, monkey := range monkeys {
		magic *= int(monkey.modulus)
	}
	return magic
}

func round(monkeys []*Monkey, modifier uint64, debug bool) ([]*Monkey, error) {
	magic := getMagicNumber(monkeys)
	for _, monkey := range monkeys {
		if debug {
			e.Print("Monkey %v:", monkey.id)
		}
		for len(monkey.items) > 0 {
			monkey.icount += 1

			item := monkey.items[0]
			monkey.items = monkey.items[1:]
			if debug {
				e.Print("  Monkey inpsects an item with a worry level of %v.", item)
			}

			var op *big.Int
			if monkey.operand == "old" {
				op = item
			} else {
				po, err := strconv.Atoi(monkey.operand)
				if err != nil {
					return nil, fmt.Errorf("unrecognized operand (%v): %v", monkey.operand, err)
				}
				op = big.NewInt(int64(po))
			}

			if monkey.operate == '*' {
				// val := item.Mul(item, op)
				var val big.Int
				val.Mul(item, op)
				if debug {
					// e.Print("    [%v] Worry %v * %v = %v", monkey.id, item, op, val)
				}
				// item *= op
				item = &val
				if debug {
					e.Print("    Worry level is multiplied by %v to %v.", op, item)
				}
			} else if monkey.operate == '+' {
				var val big.Int
				val.Add(item, op)
				if debug {
					// e.Print("    [%v] Worry %v + %v = %v", monkey.id, item, op, val)
				}
				// item *= op
				item = &val
				if debug {
					e.Print("    Worry level is increased by %v to %v.", op, item)
				}
			} else {
				return nil, fmt.Errorf("unrecognized operate (%c)", monkey.operate)
			}

			item.Div(item, big.NewInt(int64(modifier)))
			if debug {
				e.Print("    Monkey gets bored with item. Worry level is divided by 3 to %v.", item)
			}

			/* this changes the value of item */
			var tmp big.Int
			tmp.Mod(item, big.NewInt(int64(monkey.modulus)))
			item.Mod(item, big.NewInt(int64(magic)))
			// e.Print(tmp.String())
			if tmp.Text(10) == "0" {
				// if item%monkey.modulus == 0 {
				if debug {
					e.Print("    Current worry level is divisible by %v.", monkey.modulus)
				}
				monkey.tMonkey.items = append(monkey.tMonkey.items, item)
				if debug {
					e.Print("    Item with worry level %v is thrown to monkey %v.", item, monkey.tnext)
				}
			} else {
				if debug {
					e.Print("    Current worry level is not divisible by %v.", monkey.modulus)
				}
				monkey.fMonkey.items = append(monkey.fMonkey.items, item)
				if debug {
					e.Print("    Item with worry level %v is thrown to monkey %v.", item, monkey.fnext)
				}
			}

		}
	}
	return monkeys, nil
}

func printMonkeys(monkeys []*Monkey) {
	for _, monkey := range monkeys {
		e.Print("%v\n", monkey.ToString())
	}
}

func printItems(monkeys []*Monkey) {
	for _, monkey := range monkeys {
		e.Print("Monkey %v: %v", monkey.id, monkey.items)
	}
	e.Print("")
}

func printMonkeyCounts(monkeys []*Monkey, round int) {
	e.Print("== After round %4v ==", round)
	for _, monkey := range monkeys {
		e.Print("Monkey %v inspected items %v times.", monkey.id, monkey.icount)
	}
	e.Print("")
}

func getHighestCount(monkeys []*Monkey) int {
	var count []int
	for _, monkey := range monkeys {
		count = append(count, monkey.icount)
	}

	sort.Ints(count)
	highest := count[len(count)-1]
	secondh := count[len(count)-2]
	return highest * secondh
}

func part1(data []string) string {
	nstr := "null"

	monkeys, err := getMonkeys(data)
	if err != nil {
		e.Perror(err.Error())
		return nstr
	}

	for i, monkey := range monkeys {
		monkeys[i].fMonkey = monkeys[monkey.fnext]
		monkeys[i].tMonkey = monkeys[monkey.tnext]
	}

	// printMonkeys(monkeys)
	for i := 0; i < 20; i += 1 {
		e.Print("Round %v: ", i+1)

		monkeys, err = round(monkeys, 3, false)
		if err != nil {
			e.Perror(err.Error())
			return nstr
		}
		// e.Print("")

		printItems(monkeys)
		// printMonkeyCounts(monkeys, i+1)
	}
	// printMonkeys(monkeys)

	return strconv.Itoa(getHighestCount(monkeys))
}

func part2(data []string) string {
	nstr := "null"

	monkeys, err := getMonkeys(data)
	if err != nil {
		e.Perror(err.Error())
		return nstr
	}

	for i, monkey := range monkeys {
		monkeys[i].fMonkey = monkeys[monkey.fnext]
		monkeys[i].tMonkey = monkeys[monkey.tnext]
	}

	// printMonkeys(monkeys)
	for i := 0; i < 10000; i += 1 {
		// e.Print("Round %v: ", i+1)

		monkeys, err = round(monkeys, 1, false)
		if err != nil {
			e.Perror(err.Error())
			return nstr
		}
		// e.Print("")

		// printItems(monkeys)
		printMonkeyCounts(monkeys, i+1)
	}
	printMonkeys(monkeys)

	return strconv.Itoa(getHighestCount(monkeys))
}

func Run(datafile string, part int) {
	e.Print("Day 11 is perhaps a thing!\n")

	// data, err := e.ReadInTextFile("day11/testdata.txt")
	data, err := e.ReadInTextFile(datafile)
	if err != nil {
		e.Perror("%v", err)
		return
	}

	switch part {
	case 1:
		e.Print("Part 1: %v", part1(data))
	case 2:
		e.Print("Part 2: %v", part2(data))
	default:
		e.Perror("Unregognized part number: %v", part)
	}
}
