package day13

import (
	"adventcode/2022/lib"
	"fmt"
	"strconv"
)

const (
	number itype = iota
	list
)

type itype int

type item struct {
	t       itype
	val     string
	nval    int
	subitem []*item
}

func (i *item) String() string {
	if i.t == number {
		return strconv.Itoa(i.nval)

	} else {
		var s string
		s += "["
		for i, t := range i.subitem {
			if i > 0 {
				s += ","
			}
			s += t.String()
		}
		s += "]"
		return s
	}
}

const (
	left int = iota
	right
	equal
)

func (l *item) Compare(r *item) int {
	if l.t == number && r.t == number {
		if l.nval < r.nval {
			return left
		} else if l.nval > r.nval {
			return right
		} else {
			return equal
		}
	} else if l.t == list && r.t == number {
		temptype := &item{t: list, val: r.val, nval: 0}
		temptype.subitem = append(temptype.subitem, r)
		return l.Compare(temptype)
	} else if l.t == number && r.t == list {
		temptype := &item{t: list, val: l.val, nval: 0}
		temptype.subitem = append(temptype.subitem, l)
		return temptype.Compare(r)
	} else {
		for t := range l.subitem {
			if t >= len(r.subitem) {
				return right
			} else {
				res := l.subitem[t].Compare(r.subitem[t])
				if res == left || res == right {
					return res
				}
			}
		}
		if len(l.subitem) < len(r.subitem) {
			return left
		} else {
			return equal
		}
	}
}

func itemListStr(ilist []*item, label string) string {
	var s string

	if len(label) > 0 {
		s += label
		s += "\n"
	}

	for _, t := range ilist {
		s += fmt.Sprintf("%v\n", t)
	}
	return s
}

func itemsort(ilist []*item) []*item {
	var sortedlist []*item
	// lib.Print(itemListStr(ilist, "Before:"))

	// lib.Print("During:")
	for _, l := range ilist {
		// lib.Printf("\ninsert: %v", l)
		j := 0
		for _, r := range sortedlist {
			if l.Compare(r) == left {
				break
			}
			j += 1
		}
		if j == len(ilist) {
			j -= 1
		}
		// e.Print(" at pos %v", j)
		var prelist []*item
		var postlist []*item

		// prelist = sortedlist[:j]
		prelist = append(prelist, sortedlist[:j]...)
		if j+1 <= len(sortedlist) {
			// postlist = sortedlist[j:]
			postlist = append(postlist, sortedlist[j:]...)
		}
		// lib.Printf(itemListStr(prelist, ""))
		// lib.Printf(">%v\n", l)
		// lib.Printf(itemListStr(postlist, ""))
		// lib.Print("")

		sortedlist = append(prelist, l)
		sortedlist = append(sortedlist, postlist...)

		// lib.Print(itemListStr(sortedlist, "sorting:"))
	}
	lib.Print("")
	return sortedlist
}
