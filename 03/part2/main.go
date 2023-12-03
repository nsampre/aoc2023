package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

type numberInText struct {
	beginPos int
	endPos   int
	value    int
}

type gear struct {
	neighbours []int
}

var text = []string{}
var numbersExpr = regexp.MustCompile(`\d+`)
var numbersInText [][]numberInText
var gearStore = make(map[[2]int]gear)

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

	numbersInText = make([][]numberInText, len(text))
	for lineIdx, line := range text {
		matchingNumbersBeginEnd := numbersExpr.FindAllStringIndex(line, -1)
		matchingNumbersValues := numbersExpr.FindAllString(line, -1)
		numbersInLine := make([]numberInText, len(matchingNumbersValues))
		for idxValue, beginEnd := range matchingNumbersBeginEnd {
			value, _ := strconv.Atoi(matchingNumbersValues[idxValue])
			numbersInLine[idxValue] = numberInText{
				beginPos: beginEnd[0],
				endPos:   beginEnd[1],
				value:    value,
			}
		}
		numbersInText[lineIdx] = append(numbersInText[lineIdx], numbersInLine...)
	}

	for idxLine, line := range text {
		if len(numbersInText[idxLine]) > 0 {
			for _, numberInText := range numbersInText[idxLine] {
				for x := numberInText.beginPos - 1; x <= numberInText.endPos; x++ {
					isCloseToGear := false
					for y := idxLine - 1; y <= idxLine+1; y++ {
						if x >= 0 && x < len(line) && y >= 0 && y < len(text) {
							if rune(text[y][x]) == '*' {
								isCloseToGear = true
								if g, exists := gearStore[[2]int{x, y}]; !exists {
									gearStore[[2]int{x, y}] = gear{
										neighbours: []int{numberInText.value},
									}
								} else {
									g.neighbours = append(g.neighbours, numberInText.value)
									gearStore[[2]int{x, y}] = g
								}
								break
							}
						}
					}
					if isCloseToGear {
						break
					}
				}
			}
		}
	}

	total := 0
	for _, g := range gearStore {
		if len(g.neighbours) == 2 {
			total += g.neighbours[0] * g.neighbours[1]
		}
	}
	spew.Dump(total)
}
