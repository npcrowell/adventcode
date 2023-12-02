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

func (g Grid) isFirstRow(i int) bool {
	return g.Row(i) == 0
}

func (g Grid) isLastRow(i int) bool {
	return g.Row(i) == (g.rows - 1)
}

func (g Grid) isFirstColumn(i int) bool {
	return g.Column(i) == 0
}

func (g Grid) isLastColumn(i int) bool {
	return g.Column(i) == (g.cols - 1)
}

// Index of a cell to direction (method name) by providing Index of a cell
func (g Grid) Up(i int) int {
	if g.isFirstRow(i) {
		return g.Index(g.cols-1, g.Column(i))
	}
	return g.Index(g.Row(i)-1, g.Column(i))
}

// Index of a cell to direction (method name) by providing Index of a cell
func (g Grid) Down(i int) int {
	if g.isLastRow(i) {
		return g.Index(0, g.Column(i))
	}
	return g.Index(g.Row(i)+1, g.Column(i))
}

// Index of a cell to direction (method name) by providing Index of a cell
func (g Grid) Left(i int) int {
	if g.isFirstColumn(i) {
		return g.Index(g.Row(i), g.rows-1)
	}
	return g.Index(g.Row(i), g.Column(i)-1)
}

// Index of a cell to direction (method name) by providing Index of a cell
func (g Grid) Right(i int) int {
	if g.isLastColumn(i) {
		return g.Index(g.Row(i), 0)
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
