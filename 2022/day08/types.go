package day08

import "fmt"

type tree struct {
	x     uint
	y     uint
	v     int
	up    []int
	down  []int
	left  []int
	right []int
}

func viewDistanceRange(v int, r []int, interval int) int {
	count := 0
	if len(r) == 0 {
		return count
	}

	startval := 0
	endval := len(r)

	if interval < 0 {
		startval = len(r) - 1
		endval = -1
	}

	for i := startval; i != endval; i += interval {
		count += 1
		if r[i] >= v {
			break
		}

	}
	return count
}

func (t tree) ViewingDistance() int {
	view := viewDistanceRange(t.v, t.up, -1)
	view *= viewDistanceRange(t.v, t.down, 1)
	view *= viewDistanceRange(t.v, t.left, -1)
	view *= viewDistanceRange(t.v, t.right, 1)

	return view
}

func (t tree) Print() {
	fmt.Printf("(%v, %v) [%v] %v:%v, %v:%v\n",
		t.x, t.y, t.v,
		len(t.up), len(t.down),
		len(t.left), len(t.right))
}

func (t tree) IsVisible() bool {

	if t.isVisibleUp() ||
		t.isVisibleDown() ||
		t.isVisibleLeft() ||
		t.isVisibleRight() {
		return true
	}
	return false
}

func (t tree) isVisibleUp() bool {
	if len(t.up) == 0 {
		return true
	}

	for _, x := range t.up {
		if x >= t.v {
			return false
		}
	}
	return true
}
func (t tree) isVisibleDown() bool {
	if len(t.down) == 0 {
		return true
	}
	for _, x := range t.down {
		if x >= t.v {
			return false
		}
	}
	return true
}

func (t tree) isVisibleLeft() bool {
	if len(t.left) == 0 {
		return true
	}
	for _, x := range t.left {
		if x >= t.v {
			return false
		}
	}
	return true
}
func (t tree) isVisibleRight() bool {
	if len(t.right) == 0 {
		return true
	}
	for _, x := range t.right {
		if x >= t.v {
			return false
		}
	}
	return true
}
