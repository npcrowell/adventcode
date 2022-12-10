package day09

import (
	"adventcode/2022/lib"
	"fmt"
	"math"
)

type snake struct {
	head *section
	tail *section
}

func newSnake() *snake {
	head := newSection(nil, 0)
	s := &snake{head: head, tail: head}

	return s
}

func (s *snake) print() {
	sect := s.head
	if sect == nil {
		lib.Print("snake is empty")
		return
	}
}

func (s *snake) addSection() {
	sect := newSection(s.tail, s.tail.index+1)
	s.tail = sect
}

func (s *snake) moveUp(dist int) {
	for i := 0; i < dist; i += 1 {
		s.head.y += 1
		s.head.adjust()
	}
}

func (s *snake) moveLeft(dist int) {
	for i := 0; i < dist; i += 1 {
		s.head.x -= 1
		s.head.adjust()
	}
}

func (s *snake) moveRight(dist int) {
	for i := 0; i < dist; i += 1 {
		s.head.x += 1
		s.head.adjust()
	}
}

func (s *snake) moveDown(dist int) {
	for i := 0; i < dist; i += 1 {
		s.head.y -= 1
		s.head.adjust()
	}
}

type section struct {
	x      int
	y      int
	index  int
	parent *section
	child  *section
	prev   []string
}

func newSection(parent *section, index int) *section {
	sect := &section{x: 0, y: 0, index: index, parent: parent, child: nil}
	if parent != nil {
		sect.parent.child = sect
	}
	sect.recordPosition()
	return sect
}

func (s *section) adjust() {
	defer s.recordPosition()

	if s.parent == nil {
		if s.child != nil {
			s.child.adjust()
			return
		}
	}

	xdist := s.parent.x - s.x
	ydist := s.parent.y - s.y
	xcomp := int(math.Abs(float64(xdist)))
	ycomp := int(math.Abs(float64(ydist)))

	if xcomp <= 1 && ycomp <= 1 {
		return
	}
	if xcomp >= 1 && ycomp >= 1 {
		s.x += xdist / xcomp
		s.y += ydist / ycomp
		if s.child != nil {
			s.child.adjust()
		}
		return
	}
	if xcomp > 1 {
		s.x += xdist / xcomp
		if s.child != nil {
			s.child.adjust()
		}
		return
	}
	if ycomp > 1 {
		s.y += ydist / ycomp
		if s.child != nil {
			s.child.adjust()
		}
		return
	}

	if s.child != nil {
		s.child.adjust()
	}
}

func (s *section) recordPosition() {

	// if s.parent != nil {
	// defer lib.Print("%v)> %v, %v", s.index, s.x, s.y)

	curloc := fmt.Sprintf("%v,%v", s.x, s.y)

	for _, p := range s.prev {
		if p == curloc {
			return
		}
	}
	s.prev = append(s.prev, curloc)
	// }

	if s.child != nil {
		s.child.recordPosition()
	}
}
