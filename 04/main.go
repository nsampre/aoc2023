package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

var text = []string{}

type scratchCard struct {
	winningNumbers []int
	mineNumbers    []int
	power          int
}

var cards []scratchCard

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	for _, line := range text {
		line = strings.TrimLeft(line, " ")
		line = strings.TrimRight(line, " ")
		line = strings.ReplaceAll(line, "  ", " ")

		numbers := strings.Split(line, ":")[1]

		winning := strings.Split(numbers, "|")[0]
		mines := strings.Split(numbers, "|")[1]

		winning = strings.TrimLeft(winning, " ")
		winning = strings.TrimRight(winning, " ")
		mines = strings.TrimLeft(mines, " ")
		mines = strings.TrimRight(mines, " ")

		winningNumbersRaw := strings.Split(winning, " ")
		mineNumbersRaw := strings.Split(mines, " ")

		winningNumbers := []int{}
		mineNumbers := []int{}

		for _, number := range winningNumbersRaw {
			n, _ := strconv.Atoi(number)
			winningNumbers = append(winningNumbers, n)
		}

		for _, number := range mineNumbersRaw {
			n, _ := strconv.Atoi(number)
			mineNumbers = append(mineNumbers, n)
		}

		cards = append(cards, scratchCard{
			winningNumbers: winningNumbers,
			mineNumbers:    mineNumbers,
			power:          1,
		})
	}

	count := 0
	for idxCard, card := range cards {
		for x := 0; x < card.power; x++ {
			winner := 0
			for _, mine := range card.mineNumbers {
				if numInList(mine, card.winningNumbers) {
					winner += 1
				}
			}
			if winner > 0 {
				for i := 1; i <= winner; i++ {
					if idxCard+i <= len(cards) {
						cards[idxCard+i].power += 1
					}
				}
			}
			count++
		}
	}
	spew.Dump(count)
}

func numInList(n int, list []int) bool {
	for _, i := range list {
		if i == n {
			return true
		}
	}
	return false
}
