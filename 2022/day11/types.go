package day11

import (
	"adventcode/2022/lib"
	"fmt"
	"math/big"
)

type Monkey struct {
	id        int
	items     []*big.Int
	operation string
	operate   rune
	operand   string
	modulus   uint64
	tMonkey   *Monkey
	tnext     int
	fMonkey   *Monkey
	fnext     int
	icount    int
}

func newMonkey() *Monkey {
	return &Monkey{}
}

func (m *Monkey) Print() {
	lib.Print("id: %v (%p)", m.id, m)
	lib.Print("  items:    %v", m.items)
	lib.Print("  operate:  %c", m.operate)
	lib.Print("  operand:  %v", m.operand)
	lib.Print("  test:     %v", m.modulus)
	lib.Print("  if true:  %2v (%p)", m.tnext, m.tMonkey)
	lib.Print("  if false: %2v (%p)", m.fnext, m.fMonkey)
}

func (m *Monkey) ToString() string {
	s := fmt.Sprintf("id: %2v (%p)\n", m.id, m)
	s += fmt.Sprintf("  items:    %v\n", m.items)
	s += fmt.Sprintf("  operate:  %3c\n", m.operate)
	s += fmt.Sprintf("  operand:  %3v\n", m.operand)
	s += fmt.Sprintf("  test:     %3v\n", m.modulus)
	s += fmt.Sprintf("  if true:  %3v (%p)\n", m.tnext, m.tMonkey)
	s += fmt.Sprintf("  if false: %3v (%p)\n", m.fnext, m.fMonkey)
	s += fmt.Sprintf("  ICOUNT:   %3v\n", m.icount)
	return s
}
