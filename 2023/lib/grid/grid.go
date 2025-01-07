package grid

import (
	"errors"
	"fmt"
)

// Interface used as a generic
type T interface{}

// Grid is a dwo dimentional array that can be walked
// Left(), Right(), Up(), Down() cell by cell infinitely
type Grid struct {
	cells []T
	cols  int
	rows  int
}

func NewGrid(cols, rows int) *Grid {
	g := Grid{
		cells: make([]T, cols*rows),
		cols:  cols,
		rows:  rows,
	}
	for i := range g.cells {
		g.cells[i] = 0
	}
	return &g
}

// Public getter to get all cells as a single array,
func (g Grid) Cells() []T {
	return g.cells
}

// Get one dimension size
func (g Grid) Cols() int {
	return g.cols
}

// Get one dimension size
func (g Grid) Rows() int {
	return g.cols
}

// Get index of a Row by providing Index of cell
func (g Grid) Row(i int) int {
	return int(i / g.rows)
}

// Get index of a column by providing Index of a cell
func (g Grid) Column(i int) int {
	return i % g.cols
}

// Get index and a Row and Column by providing Index of a cell
func (g Grid) RowColumn(i int) (int, int) {
	return g.Row(i), g.Column(i)
}

// Get value of a cell by providing Index of a cell
func (g Grid) Value(i int) T {
	return g.cells[i]
}

// Get Index of a cell by providing row and column index
func (g Grid) Index(r int, c int) int {
	return (g.rows * r) + c
}

func (g Grid) IsFirstRow(i int) bool {
	return g.Row(i) == 0
}

func (g Grid) IsLastRow(i int) bool {
	return g.Row(i) == (g.rows - 1)
}

func (g Grid) IsFirstColumn(i int) bool {
	return g.Column(i) == 0
}

func (g Grid) IsLastColumn(i int) bool {
	return g.Column(i) == (g.cols - 1)
}

// Index of a cell to direction (method name) by providing Index of a cell
func (g Grid) Up(i int, wrap bool) int {
	if g.IsFirstRow(i) {
		if wrap {
			return g.Index(g.cols-1, g.Column(i))
		} else {
			return -1
		}
	}
	return g.Index(g.Row(i)-1, g.Column(i))
}

// Index of a cell to direction (method name) by providing Index of a cell
func (g Grid) Down(i int, wrap bool) int {
	if g.IsLastRow(i) {
		if wrap {
			return g.Index(0, g.Column(i))
		} else {
			return -1
		}
	}
	return g.Index(g.Row(i)+1, g.Column(i))
}

// Index of a cell to direction (method name) by providing Index of a cell
func (g Grid) Left(i int, wrap bool) int {
	if g.IsFirstColumn(i) {
		if wrap {
			return g.Index(g.Row(i), g.rows-1)
		} else {
			return -1
		}
	}
	return g.Index(g.Row(i), g.Column(i)-1)
}

// Index of a cell to direction (method name) by providing Index of a cell
func (g Grid) Right(i int, wrap bool) int {
	if g.IsLastColumn(i) {
		if wrap {
			return g.Index(g.Row(i), 0)
		} else {
			return -1
		}
	}
	return g.Index(g.Row(i), g.Column(i)+1)
}

// Set value of a cell where Index is first arg and value second
func (g Grid) Set(row, col int, v T) error {
	if row >= g.rows || col >= g.cols {
		return errors.New("index out of bounds")
	}
	g.cells[g.Index(row, col)] = v
	return nil
}

func (g Grid) String() string {
	out := ""
	maxlen := 0
	for _, c := range g.cells {
		checklen := len(fmt.Sprintf("%v", c))
		if checklen > maxlen {
			maxlen = checklen
		}
	}
	cellfmt := fmt.Sprintf(" %%%dv", maxlen)

	for i, c := range g.cells {
		if i > 0 && i%g.cols == 0 {
			out += "\n"
		}
		out += fmt.Sprintf(cellfmt, c)
	}
	return out
}

// Search for instances of item , starting at logical index start
// and finding count number of instances
// negative number for count returns all instances
// returns list of row, column
func (g Grid) Search(item T, sIndex int, fIndex int, count int) []int {
	var ind []int
	if sIndex >= len(g.cells) {
		return ind
	}

	if fIndex == -1 {
		fIndex = len(g.cells)
	}
	if count < 0 {
		count = len(g.cells)
	}

	for i, v := range g.cells[sIndex:fIndex] {
		if len(ind) >= count {
			return ind
		}
		if item == v {
			ind = append(ind, i)
		}
	}
	return ind
}

func (g Grid) GetRadius(index int, corners bool, wrap bool) []T {
	var radius []T

	if !g.IsFirstRow(index) {
		if !g.IsFirstColumn(index) && corners {
			radius = append(radius, g.Left(g.Up(index, wrap), wrap))
		}
		radius = append(radius, g.Up(index, false))
		if !g.IsLastColumn(index) && corners {
			radius = append(radius, g.Right(g.Up(index, wrap), wrap))
		}
	}

	if !g.IsFirstColumn(index) {
		radius = append(radius, g.Left(index, wrap))
	}
	radius = append(radius, index)
	if !g.IsLastColumn(index) {
		radius = append(radius, g.Right(index, wrap))
	}

	if !g.IsLastRow(index) {
		if !g.IsFirstColumn(index) && corners {
			radius = append(radius, g.Left(g.Down(index, wrap), wrap))
		}
		radius = append(radius, g.Down(index, wrap))
		if !g.IsLastColumn(index) && corners {
			radius = append(radius, g.Right(g.Down(index, wrap), wrap))
		}
	}

	return radius
}
