package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

var text = []string{}
var numbersExpr = regexp.MustCompile(`\d+`)

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

	total := 0
	for idxLine, line := range text {
		matchingNumbersBeginEnd := numbersExpr.FindAllStringIndex(line, -1)
		matchingNumbersValues := numbersExpr.FindAllString(line, -1)
		for idxValue, beginEnd := range matchingNumbersBeginEnd {
			closeToSym := false
			for x := beginEnd[0] - 1; x <= beginEnd[1]; x++ {
				for y := idxLine - 1; y <= idxLine+1; y++ {
					if x >= 0 && x < len(line) && y >= 0 && y < len(text) {
						if !(rune(text[y][x]) >= '0' && rune(text[y][x]) <= '9') && rune(text[y][x]) != '.' {
							closeToSym = true
							num, _ := strconv.Atoi(matchingNumbersValues[idxValue])
							total += num
							break
						}
					}
				}
				if closeToSym {
					break
				}
			}
		}
	}
	spew.Dump(total)
}
