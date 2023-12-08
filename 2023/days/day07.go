package days

import (
	"advent2023/lib"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var cards = map[rune]string{
	'A': "A",
	'K': "K",
	'Q': "Q",
	'J': "J",
	'T': "T",
	'9': "9",
	'8': "8",
	'7': "7",
	'6': "6",
	'5': "5",
	'4': "4",
	'3': "3",
	'2': "2",
	'M': "J",
}

var cardscore = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
	'M': 0,
}

const (
	HIGHCARD     = iota
	ONEPAIR      = iota
	TWOPAIR      = iota
	THREEOFAKIND = iota
	FULLHOUSE    = iota
	FOUROFAKIND  = iota
	FIVEOFAKIND  = iota
)

type hand struct {
	cards [5]rune
	htype int
	bid   int
	score int
}

func getEnumName(e int) string {
	switch e {
	case HIGHCARD:
		return "HIGH CARD"
	case ONEPAIR:
		return "ONE PAIR"
	case TWOPAIR:
		return "TWO PAIR"
	case THREEOFAKIND:
		return "THREE OF A KIND"
	case FULLHOUSE:
		return "FULL HOUSE"
	case FOUROFAKIND:
		return "FOUR OF A KIND"
	case FIVEOFAKIND:
		return "FIVE OF A KIND"
	default:
		return fmt.Sprintf("unknown enum: %v", e)
	}
}

func (h hand) String() string {
	s := "["
	for _, c := range h.cards {
		s += cards[c]
	}
	s += fmt.Sprintf("] %12v", h.score)
	s += fmt.Sprintf(" ($%5v) ", h.bid)
	s += getEnumName(h.htype)
	return s
}

func countintcontains(A []int, b int) int {
	count := 0
	for _, a := range A {
		if a == b {
			count += 1
		}
	}
	return count
}

func countrunecontains(A []rune, b rune) int {
	count := 0
	for _, a := range A {
		if a == b {
			count += 1
		}
	}
	return count
}

func day5_adjustforwildcards(htype int, wildcards int) int {
	rval := htype

	switch wildcards {
	case 5:
		rval = FIVEOFAKIND
	case 4:
		rval = FIVEOFAKIND
	case 3:
		switch htype {
		case HIGHCARD:
			rval = FOUROFAKIND
		case ONEPAIR:
			rval = FIVEOFAKIND
		}
	case 2:
		switch htype {
		case HIGHCARD:
			rval = THREEOFAKIND
		case ONEPAIR:
			rval = FOUROFAKIND
		case THREEOFAKIND:
			rval = FIVEOFAKIND
		}
	case 1:
		switch htype {
		case HIGHCARD:
			rval = ONEPAIR
		case ONEPAIR:
			rval = THREEOFAKIND
		case TWOPAIR:
			rval = FULLHOUSE
		case THREEOFAKIND:
			rval = FOUROFAKIND
		case FOUROFAKIND:
			rval = FIVEOFAKIND
		}
	default:
		lib.Perror("Bad wildcard count %v", wildcards)
	}

	return rval
}

func day5_getHandType(cards [5]rune) int {
	compares := [5]int{0, 0, 0, 0, 0}

	for i, c := range cards {
		for _, k := range cards {
			if c == k && c != 'M' {
				compares[i] += 1
			}
		}
	}

	rval := HIGHCARD

	if countintcontains(compares[:], 5) > 0 {
		rval = FIVEOFAKIND
	} else if countintcontains(compares[:], 4) > 0 {
		rval = FOUROFAKIND
	} else if countintcontains(compares[:], 3) > 0 {
		if countintcontains(compares[:], 2) > 0 {
			rval = FULLHOUSE
		} else {
			rval = THREEOFAKIND
		}
	} else if countintcontains(compares[:], 2) > 0 {
		if countintcontains(compares[:], 2) > 2 {
			rval = TWOPAIR
		} else {
			rval = ONEPAIR
		}
	}

	wildcards := countrunecontains(cards[:], 'M')
	if wildcards > 0 {
		rval = day5_adjustforwildcards(rval, wildcards)
	}

	return rval
}

func day5_calculateHandScore(htype int, cards [5]rune) int {

	score := htype * 10000000000
	score += cardscore[cards[0]] * 100000000
	score += cardscore[cards[1]] * 1000000
	score += cardscore[cards[2]] * 10000
	score += cardscore[cards[3]] * 100
	score += cardscore[cards[4]]
	return score
}

func day5_newHand(inp string, part2 bool) hand {
	splinp := strings.Split(inp, " ")

	bid, err := strconv.Atoi(splinp[1])
	if err != nil {
		lib.Perror("unable to convert %v to int", splinp[1])
		bid = 0
	}

	var c [5]rune

	for i, r := range splinp[0] {
		if part2 && r == 'J' {
			r = 'M'
		}
		c[i] = r
	}

	htype := day5_getHandType(c)

	return hand{
		cards: c,
		htype: htype,
		bid:   bid,
		score: day5_calculateHandScore(htype, c),
	}
}

func day5_sortHands(hands []hand) []hand {
	hd := make([]hand, len(hands))

	for i, h := range hands {
		count := 0
		for j, k := range hands {
			if i == j {
				continue
			}
			if h.score > k.score {
				count += 1
			}
		}
		hd[count] = h
	}
	return hd
}

func Day07(part int, data []string) (int, error) {
	switch part {
	case 1:
		var hands []hand
		for _, line := range data {
			hands = append(hands, day5_newHand(line, false))
			hands = day5_sortHands(hands)
		}
		sum := 0
		for i, hand := range hands {
			winnings := (i + 1) * hand.bid
			sum += winnings
			lib.Debug("[%6v] %v", winnings, hand)
		}
		return sum, nil
	case 2:
		var hands []hand
		for _, line := range data {
			hands = append(hands, day5_newHand(line, true))
			hands = day5_sortHands(hands)
		}
		sum := 0
		for i, hand := range hands {
			winnings := (i + 1) * hand.bid
			sum += winnings
			lib.Debug("[%6v] %v", winnings, hand)
		}
		return sum, nil
	default:
		return 0, errors.New("not implemented")
	}
}
