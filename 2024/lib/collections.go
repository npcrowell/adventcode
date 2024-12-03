package lib

import "advent2024/lib/grid"

func ConvertStringArrayToGrid(data []string) (*grid.Grid, error) {
	rows := 0
	cols := len(data)

	for _, line := range data {
		if len(line) > rows {
			rows = len(line)
		}
	}

	g := grid.NewGrid(rows, cols)

	for y, line := range data {
		for x, char := range line {
			g.Set(y, x, string(char))
		}
	}

	return g, nil
}
