package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
)

var text []string
var isLoopGrid [][]rune

func main() {
	// f, err := os.Open("grid2.example")
	f, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		text = append(text, line)
	}

	// for _, line := range text {
	// 	fmt.Println(line)
	// }

	isLoopGrid = make([][]rune, len(text))
	for i := range isLoopGrid {
		isLoopGrid[i] = make([]rune, len(text[0]))
		for j := range isLoopGrid[i] {
			isLoopGrid[i][j] = ' '
		}
	}

	sX, sY := findStartingPoint()

	follow(sX, sY, sX, sY+1, "TOP") // going down
}

func findStartingPoint() (int, int) {
	for i, line := range text {
		for j, char := range line {
			if char == 'S' {
				return j, i
			}
		}
	}
	panic("No starting point found")
	return -1, -1
}

func follow(sX, sY, x, y int, comingFrom string) {
	if x == sX && y == sY {

		isLoopGrid[y][x] = '7'
		spew.Dump(x, y)
		xStart := 0
		yStart := 0
		for {
			pprint := true
			printLoopGrid(pprint)
			flood(xStart, yStart)
			// findOpenings(xStart, yStart)
			// openWindows()
			reduce()
			trim()
			printLoopGrid(pprint)

			return
		}
	}

	currentX, currentY := x, y
	currentStep := rune(text[currentY][currentX])
	isLoopGrid[y][x] = currentStep

	switch true {
	case currentStep == 'F' && comingFrom == "RIGHT":
		follow(sX, sY, currentX, currentY+1, "TOP")
	case currentStep == 'F' && comingFrom == "BOTTOM":
		follow(sX, sY, currentX+1, currentY, "LEFT")

	case currentStep == '7' && comingFrom == "LEFT":
		follow(sX, sY, currentX, currentY+1, "TOP")
	case currentStep == '7' && comingFrom == "BOTTOM":
		follow(sX, sY, currentX-1, currentY, "RIGHT")

	case currentStep == 'J' && comingFrom == "LEFT":
		follow(sX, sY, currentX, currentY-1, "BOTTOM")
	case currentStep == 'J' && comingFrom == "TOP":
		follow(sX, sY, currentX-1, currentY, "RIGHT")

	case currentStep == 'L' && comingFrom == "RIGHT":
		follow(sX, sY, currentX, currentY-1, "BOTTOM")
	case currentStep == 'L' && comingFrom == "TOP":
		follow(sX, sY, currentX+1, currentY, "LEFT")

	case currentStep == '-' && comingFrom == "LEFT":
		follow(sX, sY, currentX+1, currentY, "LEFT")
	case currentStep == '-' && comingFrom == "RIGHT":
		follow(sX, sY, currentX-1, currentY, "RIGHT")

	case currentStep == '|' && comingFrom == "TOP":
		follow(sX, sY, currentX, currentY+1, "TOP")
	case currentStep == '|' && comingFrom == "BOTTOM":
		follow(sX, sY, currentX, currentY-1, "BOTTOM")

	default:
		fmt.Printf("%v\n", currentStep == 'L' && comingFrom == "TOP")
		fmt.Printf("%c, %c, %v\n", currentStep, 'L', comingFrom == "TOP")
		fmt.Printf("%d %d %d %d %c %v\n", sX, sY, x, y, currentStep, comingFrom)
		panic("Unknown character")
	}
}

func printLoopGrid(pprint bool) {
	points := 0
	fmt.Println()
	for _, line := range isLoopGrid {
		line := string(line)
		if pprint {
			line = strings.ReplaceAll(line, "F", "┌")
			line = strings.ReplaceAll(line, "J", "┘")
			line = strings.ReplaceAll(line, "L", "└")
			line = strings.ReplaceAll(line, "7", "┐")
			line = strings.ReplaceAll(line, "|", "|")
			line = strings.ReplaceAll(line, " ", ".")
		}
		for _, char := range line {
			if char == '.' {
				points++
			}
			fmt.Printf("%c", char)
		}
		fmt.Printf("\n")
	}
	fmt.Println()
	spew.Dump(points)
}

func flood(x, y int) {
	isLoopGrid[y][x] = 'W'

	if y+1 < len(isLoopGrid) && isLoopGrid[y+1][x] == ' ' {
		flood(x, y+1)
	}
	if y-1 >= 0 && isLoopGrid[y-1][x] == ' ' {
		flood(x, y-1)
	}
	if x+1 < len(isLoopGrid[0]) && isLoopGrid[y][x+1] == ' ' {
		flood(x+1, y)
	}
	if x-1 >= 0 && isLoopGrid[y][x-1] == ' ' {
		flood(x-1, y)
	}

	// findOpenings(x, y)
}

var iteration = 0

func reduce() {
	iteration++
	if iteration > 40000 {
		return
	}
	found := false
	for y, line := range isLoopGrid {
		for x := range line {
			if isLoopGrid[y][x] == 'F' && isLoopGrid[y][x+1] == '-' && isLoopGrid[y+1][x] == '|' && isLoopGrid[y+1][x+1] == 'Q' {
				isLoopGrid[y][x] = 'Q'
				isLoopGrid[y][x+1] = 'F'
				isLoopGrid[y+1][x] = 'F'
				isLoopGrid[y+1][x+1] = 'J'
				found = true
			}
			if isLoopGrid[y][x] == 'F' && isLoopGrid[y][x+1] == '-' && isLoopGrid[y+1][x] == 'J' && isLoopGrid[y+1][x+1] == 'Q' {
				isLoopGrid[y][x] = 'Q'
				isLoopGrid[y][x+1] = 'F'
				isLoopGrid[y+1][x] = '-'
				isLoopGrid[y+1][x+1] = 'J'
				found = true
			}
			if isLoopGrid[y][x] == 'F' && isLoopGrid[y][x+1] == 'J' && isLoopGrid[y+1][x] == 'J' && isLoopGrid[y+1][x+1] == 'Q' {
				isLoopGrid[y][x] = 'Q'
				isLoopGrid[y][x+1] = '|'
				isLoopGrid[y+1][x] = '-'
				isLoopGrid[y+1][x+1] = 'J'
				found = true
			}
			if isLoopGrid[y][x] == '|' && isLoopGrid[y+1][x] == 'J' && isLoopGrid[y][x-1] == 'Q' && isLoopGrid[y+1][x-1] == '-' {
				// isLoopGrid[y][x] = 'J'
				// isLoopGrid[y+1][x] = 'Q'
				// isLoopGrid[y][x-1] = 'F'
				// isLoopGrid[y+1][x-1] = 'J'
				// found = true
			}
			if isLoopGrid[y][x] == 'F' && isLoopGrid[y+1][x] == 'J' && isLoopGrid[y][x+1] == 'J' && isLoopGrid[y+1][x+1] == 'Q' {
				isLoopGrid[y][x] = 'Q'
				isLoopGrid[y+1][x] = '-'
				isLoopGrid[y][x+1] = '|'
				isLoopGrid[y+1][x+1] = 'J'
				found = true
			}
			if isLoopGrid[y][x] == 'F' && isLoopGrid[y][x+1] == 'J' && isLoopGrid[y+1][x] == '|' && isLoopGrid[y+1][x+1] == 'Q' {
				isLoopGrid[y][x] = 'Q'
				isLoopGrid[y][x+1] = '|'
				isLoopGrid[y+1][x] = 'F'
				isLoopGrid[y+1][x+1] = 'J'
				found = true
			}
			if isLoopGrid[y][x] == '|' && isLoopGrid[y][x+1] == 'Q' && isLoopGrid[y+1][x] == 'L' && isLoopGrid[y+1][x+1] == '-' {
				isLoopGrid[y][x] = 'L'
				isLoopGrid[y][x+1] = '7'
				isLoopGrid[y+1][x] = 'Q'
				isLoopGrid[y+1][x+1] = 'L'
				found = true
			}
			if isLoopGrid[y][x] == '7' && isLoopGrid[y][x+1] == 'Q' && isLoopGrid[y+1][x] == 'L' && isLoopGrid[y+1][x+1] == '-' {
				isLoopGrid[y][x] = '-'
				isLoopGrid[y][x+1] = '7'
				isLoopGrid[y+1][x] = 'Q'
				isLoopGrid[y+1][x+1] = 'L'
				found = true
			}

			if isLoopGrid[y][x] == 'F' && isLoopGrid[y][x+1] == '7' && verticalBars(isLoopGrid[y+1][x], isLoopGrid[y+1][x+1]) {
				// checked
				isLoopGrid[y][x] = 'Q'
				isLoopGrid[y][x+1] = 'Q'
				isLoopGrid[y+1][x] = 'F'
				isLoopGrid[y+1][x+1] = '7'
				found = true
			}
			if isLoopGrid[y][x] == 'F' && isLoopGrid[y][x+1] == '7' && verticalBarsOpenFromTopToLeft(isLoopGrid[y+1][x], isLoopGrid[y+1][x+1]) {
				isLoopGrid[y][x] = 'Q'
				isLoopGrid[y][x+1] = 'Q'
				isLoopGrid[y+1][x] = '-'
				isLoopGrid[y+1][x+1] = '7'
				found = true
			}
			if isLoopGrid[y][x] == 'F' && isLoopGrid[y][x+1] == '7' && verticalBarsOpenFromTopToRight(isLoopGrid[y+1][x], isLoopGrid[y+1][x+1]) {
				// checkd
				isLoopGrid[y][x] = 'Q'
				isLoopGrid[y][x+1] = 'Q'
				isLoopGrid[y+1][x] = 'F'
				isLoopGrid[y+1][x+1] = '-'
				found = true
			}

			if isLoopGrid[y][x] == 'F' && isLoopGrid[y][x+1] == '7' && verticalAndOpenFromTop(isLoopGrid[y+1][x], isLoopGrid[y+1][x+1]) {
				isLoopGrid[y][x] = 'Q'
				isLoopGrid[y][x+1] = 'Q'
				isLoopGrid[y+1][x] = '-'
				isLoopGrid[y+1][x+1] = '-'
				found = true
			}
			if isLoopGrid[y][x] == 'L' && isLoopGrid[y][x+1] == 'J' && verticalAndOpenFromBottom(isLoopGrid[y-1][x], isLoopGrid[y-1][x+1]) {
				isLoopGrid[y][x] = 'Q'
				isLoopGrid[y][x+1] = 'Q'
				isLoopGrid[y-1][x] = '-'
				isLoopGrid[y-1][x+1] = '-'
				found = true
			}
			if isLoopGrid[y][x] == 'L' && isLoopGrid[y][x+1] == 'J' && verticalBars(isLoopGrid[y-1][x], isLoopGrid[y-1][x+1]) {
				isLoopGrid[y][x] = 'Q'
				isLoopGrid[y][x+1] = 'Q'
				isLoopGrid[y-1][x] = 'L'
				isLoopGrid[y-1][x+1] = 'J'
				found = true
			}
			if isLoopGrid[y][x] == 'L' && isLoopGrid[y][x+1] == 'J' && verticalAndOpenFromBottomToLeft(isLoopGrid[y-1][x], isLoopGrid[y-1][x+1]) {
				isLoopGrid[y][x] = 'Q'
				isLoopGrid[y][x+1] = 'Q'
				isLoopGrid[y-1][x] = '-'
				isLoopGrid[y-1][x+1] = 'J'
				found = true
			}
			if isLoopGrid[y][x] == 'L' && isLoopGrid[y][x+1] == 'J' && verticalAndOpenFromBottomToRight(isLoopGrid[y-1][x], isLoopGrid[y-1][x+1]) {
				isLoopGrid[y][x] = 'Q'
				isLoopGrid[y][x+1] = 'Q'
				isLoopGrid[y-1][x] = 'L'
				isLoopGrid[y-1][x+1] = '-'
				found = true
			}

			if isLoopGrid[y][x] == 'F' && isLoopGrid[y+1][x] == 'L' && horizontalBars(isLoopGrid[y][x+1], isLoopGrid[y+1][x+1]) {
				// checked
				isLoopGrid[y][x] = 'Q'
				isLoopGrid[y+1][x] = 'Q'
				isLoopGrid[y][x+1] = 'F'
				isLoopGrid[y+1][x+1] = 'L'
				found = true
			}

			if isLoopGrid[y][x] == 'F' && isLoopGrid[y+1][x] == 'L' && horizontalOpensFromLeft(isLoopGrid[y][x+1], isLoopGrid[y+1][x+1]) {
				isLoopGrid[y][x] = 'Q'
				isLoopGrid[y+1][x] = 'Q'
				isLoopGrid[y][x+1] = '|'
				isLoopGrid[y+1][x+1] = '|'
				found = true
			}
			if isLoopGrid[y][x] == 'F' && isLoopGrid[y+1][x] == 'L' && horizontalOpensFromLeftToBottom(isLoopGrid[y][x+1], isLoopGrid[y+1][x+1]) {
				isLoopGrid[y][x] = 'Q'
				isLoopGrid[y+1][x] = 'Q'
				isLoopGrid[y][x+1] = 'F'
				isLoopGrid[y+1][x+1] = '|'
				found = true
			}

			if isLoopGrid[y][x] == 'F' && isLoopGrid[y+1][x] == 'L' && horizontalOpensFromLeftToTop(isLoopGrid[y][x+1], isLoopGrid[y+1][x+1]) {
				isLoopGrid[y][x] = 'Q'
				isLoopGrid[y+1][x] = 'Q'
				isLoopGrid[y][x+1] = '|'
				isLoopGrid[y+1][x+1] = 'L'
				found = true
			}

			if isLoopGrid[y][x] == '7' && isLoopGrid[y+1][x] == 'J' && isLoopGrid[y][x-1] == 'L' && isLoopGrid[y+1][x-1] == '-' {
				isLoopGrid[y][x] = 'Q'
				isLoopGrid[y+1][x] = 'Q'
				isLoopGrid[y][x-1] = '|'
				isLoopGrid[y+1][x-1] = 'J'
				found = true
			} else if isLoopGrid[y][x] == '7' && isLoopGrid[y+1][x] == 'J' && isLoopGrid[y][x-1] == '-' && isLoopGrid[y+1][x-1] == '-' {
				isLoopGrid[y][x] = 'Q'
				isLoopGrid[y+1][x] = 'Q'
				isLoopGrid[y][x-1] = '7'
				isLoopGrid[y+1][x-1] = 'J'
				found = true
			} else if isLoopGrid[y][x] == '7' && isLoopGrid[y+1][x] == 'J' && isLoopGrid[y][x-1] == '-' && isLoopGrid[y+1][x-1] == 'F' {
				isLoopGrid[y][x] = 'Q'
				isLoopGrid[y+1][x] = 'Q'
				isLoopGrid[y][x-1] = '7'
				isLoopGrid[y+1][x-1] = '|'
				found = true
			} else if isLoopGrid[y][x] == '7' && isLoopGrid[y+1][x] == 'J' && isLoopGrid[y][x-1] == 'L' && isLoopGrid[y+1][x-1] == 'F' {
				isLoopGrid[y][x] = 'Q'
				isLoopGrid[y+1][x] = 'Q'
				isLoopGrid[y][x-1] = '|'
				isLoopGrid[y+1][x-1] = '|'
				found = true
			}

		}
	}
	if found {
		exec.Command("clear")
		printLoopGrid(true)
		time.Sleep(300 * time.Millisecond)
		reduce()
	}
}

func verticalBars(r1, r2 rune) bool {
	return r1 == '|' && r2 == '|'
}
func verticalAndOpenFromTop(r1, r2 rune) bool {
	return r1 == 'J' && r2 == 'L'
}
func verticalAndOpenFromBottom(r1, r2 rune) bool {
	return r1 == '7' && r2 == 'F'
}
func verticalBarsOpenFromTopToLeft(r1, r2 rune) bool {
	return r1 == 'J' && r2 == '|'
}
func verticalBarsOpenFromTopToRight(r1, r2 rune) bool {
	return r1 == '|' && r2 == 'L'
}
func verticalAndOpenFromBottomToLeft(r1, r2 rune) bool {
	return r1 == '7' && r2 == '|'
}
func verticalAndOpenFromBottomToRight(r1, r2 rune) bool {
	return r1 == '|' && r2 == 'F'
}
func horizontalBars(r1, r2 rune) bool {
	return r1 == '-' && r2 == '-'
}
func horizontalOpensFromLeft(r1, r2 rune) bool {
	return r1 == 'J' && r2 == '7'
}
func horizontalOpensFromLeftToTop(r1, r2 rune) bool {
	return r1 == 'J' && r2 == '-'
}
func horizontalOpensFromLeftToBottom(r1, r2 rune) bool {
	return r1 == '-' && r2 == '7'
}

func trim() {
	for y, line := range isLoopGrid {
		for x := range line {
			if x+1 < len(isLoopGrid[0]) {
				if isLoopGrid[y][x] == 'W' && (isLoopGrid[y][x+1] == 'Q' || isLoopGrid[y][x+1] == ' ') {
					isLoopGrid[y][x+1] = 'W'
				}
			}
			x = len(line) - x - 2
			if x-1 >= 0 {
				if isLoopGrid[y][x] == 'W' && (isLoopGrid[y][x-1] == 'Q' || isLoopGrid[y][x-1] == ' ') {
					isLoopGrid[y][x-1] = 'W'
				}
			}
		}
	}
}
