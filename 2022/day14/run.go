package day14

import (
	e "adventcode/2022/lib"
	"strconv"
	"strings"
)

func printGrid(g *e.Grid) {
	for _, row := range g.GetRows() {
		for _, c := range row {
			if c == 0 {
				e.Printf(".")
			} else if c == 1 {
				e.Printf("#")
			} else {
				e.Printf("o")
			}
		}
		e.Print("")
	}
	e.Print("")
}

func establishGridBounds(data []string) (*e.Grid, int, int) {

	xmin := 999
	ymin := 999
	xmax := 0
	ymax := 0

	for _, line := range data {
		for _, coord := range strings.Split(line, " -> ") {
			c := strings.Split(coord, ",")
			x, _ := strconv.Atoi(c[0])
			y, _ := strconv.Atoi(c[1])

			if x < xmin {
				xmin = x
			}
			if x > xmax {
				xmax = x
			}
			if y < ymin {
				ymin = y
			}
			if y > ymax {
				ymax = y
			}
		}
	}
	xmin -= 1
	xmax += 1
	ymin = 0
	ymax += 1

	width := uint(xmax-xmin) + 1
	height := uint(ymax-ymin) + 1

	e.Print("xmin: %v", xmin)
	e.Print("xmax: %v", xmax)
	e.Print("ymin: %v", ymin)
	e.Print("ymax: %v", ymax)
	e.Print("width: %v", width)
	e.Print("height: %v", height)
	// xAdj := xmin
	// yAdj := ymin

	g := e.EmptyGrid(height, width)

	return g, (0 - xmin), (0 - ymin)

}

type coord struct {
	x int
	y int
}

func getCoordsBetween(a, b coord) []coord {
	var coords []coord
	coords = append(coords, a, b)

	if a.x == b.x {
		// e.Print("a.x (%v) == b.x (%v)", a.x, b.x)
		if a.y > b.y {
			// e.Print("a.y (%v) > b.y (%v)", a.y, b.y)
			for i := b.y + 1; i < a.y; i += 1 {
				coords = append(coords, coord{x: a.x, y: i})
			}
		} else if a.y < b.y {
			// e.Print("a.y (%v) < b.y (%v)", a.y, b.y)
			for i := a.y + 1; i < b.y; i += 1 {
				// e.Print("%v", i)
				coords = append(coords, coord{x: a.x, y: i})
			}
		}
	} else if a.y == b.y {
		if a.x > b.x {
			for i := b.x + 1; i < a.x; i += 1 {
				coords = append(coords, coord{x: i, y: a.y})
			}
		} else if a.x < b.x {
			for i := a.x + 1; i < b.x; i += 1 {
				coords = append(coords, coord{x: i, y: a.y})
			}
		}
	}
	// e.Print("a: %v", a)
	// e.Print("b: %v", b)
	// e.Print("out: %v", coords)
	return coords
}

func stringToCoord(a string) coord {
	c := strings.Split(a, ",")
	x, _ := strconv.Atoi(c[0])
	y, _ := strconv.Atoi(c[1])

	return coord{x: x, y: y}
}

func prepareGrid(data []string) (*e.Grid, int, int) {

	g, xAdj, yAdj := establishGridBounds(data)

	e.Print("Adjust xcoord by %v", xAdj)
	e.Print("Adjust ycoord by %v", yAdj)

	for _, line := range data {
		// var clist []coord
		cline := strings.Split(line, " -> ")
		for i, c := range cline {
			co := stringToCoord(c)
			if i < len(cline)-1 {
				// e.Debug("(%v,%v) -> (%v,%v)", co.x, co.y, stringToCoord(cline[i+1]).x, stringToCoord(cline[i+1]).y)
				clist := getCoordsBetween(co, stringToCoord(cline[i+1]))
				for _, coor := range clist {
					newcoorx := uint(coor.x + xAdj)
					newcoory := uint(coor.y + yAdj)

					// e.Debug("(%v,%v)", coor.x, coor.y)
					// e.Debug("(%v,%v)", newcoorx, newcoory)
					err := g.SetIndex(newcoory, newcoorx, 1)
					if err != nil {
						e.Perror(err.Error())
					}
				}
			}
		}
	}

	return g, xAdj, yAdj

}

func part1(data []string) string {
	nstr := "null"
	g, xAdj, yAdj := prepareGrid(data)

	lowbound, _ := g.GetBounds()

	for i := 0; i < 1001; i += 1 {
		pos := coord{x: 500, y: 0}
		pos.x += xAdj
		pos.y += yAdj

		j := 0
		for {
			if uint(pos.y+1) >= lowbound {
				printGrid(g)
				return strconv.Itoa(i)
			}
			j += 1

			n, _ := g.GetIndex(uint(pos.y+1), uint(pos.x))
			m, _ := g.GetIndex(uint(pos.y+1), uint(pos.x-1))
			o, _ := g.GetIndex(uint(pos.y+1), uint(pos.x+1))
			if n == 0 {
				pos.y += 1
				continue
			} else if m == 0 {
				pos.y += 1
				pos.x -= 1
				continue
			} else if o == 0 {
				pos.y += 1
				pos.x += 1
				continue
			} else {
				break
			}
		}
		g.SetIndex(uint(pos.y), uint(pos.x), 2)

		// if i > 0 && i%100 == 0 {
		// 	printGrid(g)
		// }
	}

	return nstr
}

func part2(data []string) string {
	nstr := "null"
	g, xAdj, yAdj := prepareGrid(data)
	g.AddRow(1)

	lowbound, rbound := g.GetBounds()
	lowbound -= 1

	for i := 0; i < 100001; i += 1 {

		pos := coord{x: 500, y: 0}
		pos.x += xAdj
		pos.y += yAdj
		t, _ := g.GetIndex(uint(pos.y), uint(pos.x))
		if t != 0 {
			printGrid(g)
			return strconv.Itoa(i)
		}

		for {
			if pos.x+1 >= int(rbound) {
				g.AddColumn(1)
				_, rbound = g.GetBounds()
			} else if pos.x <= 0 {
				g.AddFirstColumn()
				xAdj += 1
				pos.x += 1
			}

			n, _ := g.GetIndex(uint(pos.y+1), uint(pos.x))
			m, _ := g.GetIndex(uint(pos.y+1), uint(pos.x-1))
			o, _ := g.GetIndex(uint(pos.y+1), uint(pos.x+1))

			if uint(pos.y+1) >= lowbound {
				break
			} else if n == 0 {
				pos.y += 1
				continue
			} else if m == 0 {
				pos.y += 1
				pos.x -= 1
				continue
			} else if o == 0 {
				pos.y += 1
				pos.x += 1
				continue
			} else {
				break
			}
		}
		g.SetIndex(uint(pos.y), uint(pos.x), 2)

		if i > 0 && i%10000 == 0 {
			printGrid(g)
		}
	}

	return nstr
}

func Run(datafile string, part int) {
	e.Print("Day 14 is about falling sand?!")

	data, err := e.ReadInTextFile(datafile)
	// data, err := e.ReadInTextFile("day14/testdata.txt")
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
