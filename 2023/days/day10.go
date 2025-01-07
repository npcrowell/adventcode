package days

import (
	"advent2023/lib"
	"advent2023/lib/grid"
	"errors"
	"fmt"
)

const (
	UP    int = iota
	DOWN  int = iota
	LEFT  int = iota
	RIGHT int = iota
	INVAL int = iota
)

func dirstr(dir int) string {
	switch dir {
	case UP:
		return "UP"
	case DOWN:
		return "DOWN"
	case LEFT:
		return "LEFT"
	case RIGHT:
		return "RIGHT"
	case INVAL:
		return "INVAL"
	default:
		lib.Perror("unrecognized direction: %v", dir)
		return fmt.Sprintf("UNRECOGNIZED (%v)", dir)
	}
}

func nextDir(g *grid.Grid, index int, edir int) (xdir int, xindex int) {
	// lib.Debug("Taking in %v, %v", index, dirstr(edir))
	val := g.Value(index)
	xdir = INVAL
	xindex = -1

	switch edir {
	case UP:
		switch val {
		case "|":
			xdir = UP
			xindex = g.Up(index, false)
		case "7":
			xdir = LEFT
			xindex = g.Left(index, false)
		case "F":
			xdir = RIGHT
			xindex = g.Right(index, false)
		default:
			return INVAL, -1
		}
	case DOWN:
		switch val {
		case "|":
			xdir = DOWN
			xindex = g.Down(index, false)
		case "L":
			xdir = RIGHT
			xindex = g.Right(index, false)
		case "J":
			xdir = LEFT
			xindex = g.Left(index, false)
		default:
			return INVAL, -1
		}
	case LEFT:
		switch val {
		case "-":
			xdir = LEFT
			xindex = g.Left(index, false)
		case "L":
			xdir = UP
			xindex = g.Up(index, false)
		case "F":
			xdir = DOWN
			xindex = g.Down(index, false)
		default:
			return INVAL, -1
		}
	case RIGHT:
		switch val {
		case "-":
			xdir = RIGHT
			xindex = g.Right(index, false)
		case "7":
			xdir = DOWN
			xindex = g.Down(index, false)
		case "J":
			xdir = UP
			xindex = g.Up(index, false)
		default:
			return INVAL, -1
		}
	default:
		lib.Perror("bad entry direction: %v", edir)
		return INVAL, -1
	}

	return
}

func getStartPositionAndOpts(g *grid.Grid) (int, []int) {
	loc := g.Search("S", 0, -1, -1)[0]
	lib.Debug("%v", loc)

	up := g.Up(loc, false)
	down := g.Down(loc, false)
	left := g.Left(loc, false)
	right := g.Right(loc, false)

	var opts []int

	if up != -1 {
		if nd, _ := nextDir(g, up, UP); nd != INVAL {
			opts = append(opts, UP)
			// lib.Debug("Up: (%v, %v)", dirstr(nd), ni)
		}
	}
	if right != -1 {
		if nd, _ := nextDir(g, right, RIGHT); nd != INVAL {
			opts = append(opts, RIGHT)
			// lib.Debug("Right: (%v, %v)", dirstr(nd), ni)
		}
	}
	if down != -1 {
		if nd, _ := nextDir(g, down, DOWN); nd != INVAL {
			opts = append(opts, DOWN)
			// lib.Debug("Down: (%v, %v)", dirstr(nd), ni)
		}
	}
	if left != -1 {
		if nd, _ := nextDir(g, left, LEFT); nd != INVAL {
			opts = append(opts, LEFT)
			// lib.Debug("Left: (%v, %v)", dirstr(nd), ni)
		}
	}

	return loc, opts
}

func move(g *grid.Grid, dir int, index int) int {
	switch dir {
	case UP:
		return g.Up(index, false)
	case DOWN:
		return g.Down(index, false)
	case LEFT:
		return g.Left(index, false)
	case RIGHT:
		return g.Right(index, false)
	default:
		return -1
	}
}

func constructLoop(g *grid.Grid, index int, dir int) int {
	index = move(g, dir, index)

	// var loop []int
	count := 0
	countmax := 100000

	for index != -1 && g.Value(index) != "S" {
		// lib.Debug("Prev hop: %v, %v, %v", index, g.Value(index), dirstr(dir))
		// loop = append(loop, index)
		dir, index = nextDir(g, index, dir)

		lib.Debug("Next hop: %v, %v", index, dirstr(dir))

		count += 1
		if count >= countmax {
			lib.Perror("possible infinite loop")
			return -1
		}
	}
	// lib.Debug("Loop: %v", loop)
	lib.Debug("Count: %v", count)
	return count
}

func Day10(part int, data []string) (int, error) {
	g, err := lib.ConvertStringArrayToGrid(data)
	if err != nil {
		return 0, err
	}

	switch part {
	case 3:
		fallthrough
	case 1:
		lib.Debug("\n%v", g)
		sindex, opts := getStartPositionAndOpts(g)
		// lib.Debug("Start index: %v", sindex)
		// lib.Debug("Start options: %v", opts)
		llen := constructLoop(g, sindex, opts[0])
		return llen/2 + 1, nil
	case 2:
		return 0, nil
	default:
		return 0, errors.New("not implemented")
	}
}
