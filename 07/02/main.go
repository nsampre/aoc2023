package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

var text []string

const (
	highCard = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

type card struct {
	def   string
	typec int
	bid   int
}

var alphabet = []rune{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}
var power = make(map[rune]int)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	for i, v := range alphabet {
		power[v] = len(alphabet) - i
	}

	var cards []card
	for _, line := range text {
		ss := strings.Split(line, " ")
		card := card{def: ss[0]}
		card.bid, _ = strconv.Atoi(ss[1])

		if strings.ContainsRune(card.def, 'J') {
			firstJ := firstJ(card.def)
			card.typec = maxValueRecurse(firstJ, 0, card.def, 0)
		} else {
			card.typec = value(card.def)
		}
		cards = append(cards, card)
	}

	sort.Slice(cards, func(i, j int) bool {
		if cards[i].typec == cards[j].typec {
			for k := 0; k < len(cards[i].def); k++ {
				if power[rune(cards[i].def[k])] == power[rune(cards[j].def[k])] {
					continue
				}
				return power[rune(cards[i].def[k])] < power[rune(cards[j].def[k])]
			}
		}
		return cards[i].typec < cards[j].typec
	})

	total := int64(0)
	for rank, c := range cards {
		total += int64((rank + 1) * c.bid)
	}
	spew.Dump(total)
}

func maxValueRecurse(idx int, step int, input string, maxScore int) int {
	if idx == 5 {
		return maxScore
	}

	if step == len(alphabet) {
		for i := idx + 1; i < len(input); i++ {
			if input[i] == 'J' {
				return maxValueRecurse(i, 0, input, maxScore)
			}
		}
		return maxScore
	}

	out := []rune(input)
	out[idx] = alphabet[step]
	input = string(out)

	score := value(input)
	if score > maxScore {
		maxScore = score
	}

	for i := idx + 1; i < len(input); i++ {
		if out[i] == 'J' {
			m := maxValueRecurse(i, 0, input, maxScore)
			if m > maxScore {
				maxScore = m
			}
		}
	}

	return maxValueRecurse(idx, step+1, input, maxScore)
}

func value(hand string) int {
	attrs := make(map[rune]int)
	for _, c := range hand {
		if _, ok := attrs[c]; ok {
			attrs[c]++
		} else {
			attrs[c] = 1
		}
	}

	if len(attrs) == 5 {
		return highCard
	} else if len(attrs) == 4 {
		return onePair
	} else if len(attrs) == 3 {
		max := 0
		for _, v := range attrs {
			if v > max {
				max = v
			}
		}
		if max == 3 {
			return threeOfAKind
		}
		return twoPair
	} else if len(attrs) == 2 {
		max := 0
		for _, v := range attrs {
			if v > max {
				max = v
			}
		}
		if max == 4 {
			return fourOfAKind
		}
		return fullHouse
	} else if len(attrs) == 1 {
		return fiveOfAKind
	}

	spew.Dump(attrs)
	return 0
}

func firstJ(hand string) int {
	for i, c := range hand {
		if c == 'J' {
			return i
		}
	}
	return -1
}
