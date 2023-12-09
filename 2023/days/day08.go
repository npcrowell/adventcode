package days

import (
	"advent2023/lib"
	"errors"
	"fmt"
	"regexp"
)

type day8_node struct {
	name  string
	left  string
	right string
}

func (node day8_node) String() string {
	return fmt.Sprintf("%v{%v %v}", node.name, node.left, node.right)
}

func (node day8_node) next(next rune) string {
	if next == 'R' {
		return node.right
	} else {
		return node.left
	}
}

type day8_path struct {
	d        []rune
	curIndex int
}

func (p day8_path) String() string {
	s := ""
	s += fmt.Sprintf("[%v] ", p.curIndex)
	for _, d := range p.d[p.curIndex:] {
		s += fmt.Sprintf("%v", string(d))
	}
	for _, d := range p.d[:p.curIndex] {
		s += fmt.Sprintf("%v", string(d))
	}
	return s
}
func (p *day8_path) next() rune {
	current := p.d[p.curIndex]
	p.curIndex += 1
	if p.curIndex >= len(p.d) {
		p.curIndex = 0
	}
	return current
}

// func (p *day8_path) index(newIndex int) {
// 	p.curIndex = newIndex % len(p.d)
// }

func day8_newPath(line string) day8_path {
	p := day8_path{curIndex: 0}

	for _, c := range line {
		if c == 'L' || c == 'R' {
			p.d = append(p.d, c)
		} else {
			lib.Perror("Unrecognized character in path: %v", c)
		}
	}
	return p
}

func day8_buildMap(lines []string) (map[string]day8_node, []string) {
	noderx, err := regexp.Compile(`(\w\w\w) = [(](\w\w\w), (\w\w\w)[)]`)
	if err != nil {
		lib.Perror("Unable to compile regular expression: %v", err)
		return nil, nil
	}

	nodemap := make(map[string]day8_node, len(lines))
	var start []string

	for _, line := range lines {
		res := noderx.FindAllStringSubmatch(line, -1)
		nodemap[res[0][1]] = day8_node{
			name:  res[0][1],
			left:  res[0][2],
			right: res[0][3],
		}
		if res[0][1][len(res[0][1])-1] == 'A' {
			start = append(start, res[0][1])
		}
	}

	return nodemap, start

}

func day8_checkCompletion(cnodes []string) bool {
	// lib.Print("Checking %v", cnodes)
	sp := false
	for _, cnode := range cnodes {
		if cnode[len(cnode)-1] != 'Z' {
			return false
		}
		if sp {
			lib.Print("Checking %v", cnodes)
		}
	}
	return true
}

func Day08(part int, data []string) (int, error) {
	path := day8_newPath(data[0])

	nodemap, cnodes := day8_buildMap(data[2:])
	lib.Debug("%v", cnodes)

	switch part {
	case 3:
		fallthrough
	case 1:

		cnode := "AAA"
		count := 0
		for cnode != "ZZZ" {
			k := path.next()

			oldnode := nodemap[cnode]
			cnode = nodemap[cnode].next(k)
			newnode := nodemap[cnode]

			lib.Debug("%v: %v => %v", string(k), oldnode, newnode)
			count += 1
		}

		return count, nil
	case 2:

		count := 0
		for !day8_checkCompletion(cnodes) {
			k := path.next()

			for i := range cnodes {
				oldnode := nodemap[cnodes[i]]
				cnodes[i] = nodemap[cnodes[i]].next(k)
				newnode := nodemap[cnodes[i]]

				lib.Debug("%v: %v => %v", string(k), oldnode, newnode)
			}
			lib.Debug("")
			count += 1
		}

		return count, nil
	default:
		return 0, errors.New("not implemented")
	}
}
