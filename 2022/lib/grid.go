package lib

import "fmt"

type Grid struct {
	g        [][]int
	rowCount uint
	colCount uint
}

/**
 * Grid Public Functions
 */

func NewGrid() *Grid {
	return &Grid{rowCount: 0, colCount: 0}
}

func (g *Grid) GetRows() [][]int {
	return g.g
}

func EmptyGrid(rows uint, cols uint) *Grid {
	g := NewGrid()
	g.AddRow(rows)
	g.AddColumn(cols)
	return g
}

func (g *Grid) Raw() [][]int {
	return g.g
}

func (g *Grid) GetBounds() (uint, uint) {
	g.Align()
	return g.rowCount, g.colCount
}

func (g *Grid) Align() {
	if !g.columnsAligned() {
		g.alignColumns()
	}

}

func (g *Grid) Row(index uint) ([]int, error) {
	if index > uint(len(g.g)) {
		err := fmt.Errorf("index [%v] exceeds row count [%v]",
			index, g.rowCount)
		return nil, err
	}

	return g.g[index], nil
}

func (g *Grid) Column(index uint) ([]int, error) {

	if len(g.g) == 0 {
		err := fmt.Errorf("no rows in grid")
		return nil, err
	}

	var r []int

	for _, row := range g.g {
		if index > uint(len(row)) {
			err := fmt.Errorf("index [%v] exceeds column count [%v]",
				index, g.colCount)
			return nil, err
		}
		r = append(r, row[index])
	}

	return r, nil
}

func (g *Grid) AddRow(count uint) {
	for i := 0; uint(i) < count; i++ {
		var newRow []int
		g.g = append(g.g, newRow)

		for j := 0; j < int(g.rowCount); j++ {
			k := len(g.g) - 1
			g.g[k] = append(g.g[k], 0)
		}
	}
	g.rowCount = uint(len(g.g))
	g.alignColumns()
}

func (g *Grid) AddColumn(count uint) {
	for r := range g.g {
		for i := 0; uint(i) < count; i += 1 {
			g.g[r] = append(g.g[r], 0)
			g.colCount = uint(len(g.g[0]))
		}
	}
}

func (g *Grid) AddFirstColumn() {
	for i := range g.g {
		var newRow []int
		newRow = append(newRow, 0)
		newRow = append(newRow, g.g[i]...)
		g.g[i] = newRow
	}
}

func (g *Grid) GetIndex(x uint, y uint) (int, error) {
	if x > uint(len(g.g)) {
		err := fmt.Errorf("index [%v,%v] out of bounds [%v]", x, y, len(g.g))
		return 0, err
	}
	if y > uint(len(g.g[x])) {
		err := fmt.Errorf("index [%v,%v] out of bounds [%v,%v]", x, y, len(g.g), len(g.g[x]))
		return 0, err
	}

	return g.g[x][y], nil
}

func (g *Grid) SetIndex(x uint, y uint, value int) error {
	if x > uint(len(g.g)) {
		err := fmt.Errorf("index [%v,%v] out of bounds [%v]", x, y, len(g.g))
		return err
	}
	if y > uint(len(g.g[x])) {
		err := fmt.Errorf("index [%v,%v] out of bounds [%v,%v]", x, y, len(g.g), len(g.g[x]))
		return err
	}

	g.g[x][y] = value

	return nil
}

func (g *Grid) Print() error {
	for _, row := range g.g {
		fmt.Printf("%v\n", row)
	}
	return nil
}

/**
 * Grid Private Functions
 */

func (g *Grid) columnsAligned() bool {

	if len(g.g) == 0 {
		g.colCount = 0
		return true
	}

	rowlen := len(g.g[0])

	for _, row := range g.g {
		if len(row) != rowlen {
			return false
		}
	}

	return true
}

func (g *Grid) alignColumns() {
	if len(g.g) == 0 {
		g.setColumnBound(0)
		g.setRowBound(0)
		return
	}

	rowlen := len(g.g[0])

	for _, row := range g.g {
		if len(row) > rowlen {
			rowlen = len(row)
		}
	}

	for r := range g.g {
		if len(g.g[r]) < rowlen {
			add := rowlen - len(g.g[r])
			for i := 0; i < add; i += 1 {
				g.g[r] = append(g.g[r], 0)
			}
		}
	}

	g.setColumnBound(uint(rowlen))
	g.setRowBound(uint(len(g.g)))
}

func (g *Grid) setRowBound(rowLen uint) {
	g.rowCount = rowLen
}

func (g *Grid) setColumnBound(colLen uint) {
	g.colCount = colLen
}
