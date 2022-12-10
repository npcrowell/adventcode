package day09

import (
	"fmt"
	"math"
)

type coord struct {
	x int
	y int
}

func newCoord() *coord {
	c := coord{x: 0, y: 0}
	return &c
}

func (c *coord) move(up int, right int) {
	c.x += up
	c.y += right
}

type knot struct {
	h    *coord
	t    *coord
	prev []string
}

func (k *knot) adjustT() {
	xdist := k.h.x - k.t.x
	ydist := k.h.y - k.t.y
	xcomp := math.Abs(float64(xdist))
	ycomp := math.Abs(float64(ydist))

	if xcomp > 1 && ycomp == 1 {
		k.t.x += xdist / int(xcomp)
		k.t.y += ydist / int(ycomp)
		return
	}
	if xcomp == 1 && ycomp > 1 {
		k.t.x += xdist / int(xcomp)
		k.t.y += ydist / int(ycomp)
		return
	}
	if xcomp > 1 {
		k.t.x += xdist / int(xcomp)
		return
	}
	if ycomp > 1 {
		k.t.y += ydist / int(ycomp)
		return
	}
}

func (k *knot) accountTPos() {
	// defer e.Print("> %v, %v (%v, %v)", k.t.x, k.t.y, k.h.x, k.h.y)

	curloc := fmt.Sprintf("%v,%v", k.t.x, k.t.y)

	for _, p := range k.prev {
		if p == curloc {

			return
		}
	}
	k.prev = append(k.prev, curloc)
}

func (k *knot) moveUp(dist int) {
	for i := 0; i < dist; i += 1 {
		k.h.y += 1
		k.adjustT()
		k.accountTPos()
	}
}

func (k *knot) moveLeft(dist int) {
	for i := 0; i < dist; i += 1 {
		k.h.x -= 1
		k.adjustT()
		k.accountTPos()
	}
}

func (k *knot) moveRight(dist int) {
	for i := 0; i < dist; i += 1 {
		k.h.x += 1
		k.adjustT()
		k.accountTPos()
	}
}

func (k *knot) moveDown(dist int) {
	for i := 0; i < dist; i += 1 {
		k.h.y -= 1
		k.adjustT()
		k.accountTPos()
	}
}
