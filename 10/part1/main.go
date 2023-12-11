package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

var text []string
var isLoopGrid [][]rune

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// line = strings.ReplaceAll(line, "F", "┌")
		// line = strings.ReplaceAll(line, "J", "┘")
		// line = strings.ReplaceAll(line, "L", "└")
		// line = strings.ReplaceAll(line, "7", "┐")
		text = append(text, line)
	}

	for _, line := range text {
		fmt.Println(line)
	}

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

var history = 0

func follow(sX, sY, x, y int, comingFrom string) {
	history++

	if x == sX && y == sY {
		trimGrid()
		printLoopGrid()

		fmt.Println("looped")
		spew.Dump(history / 2)
		return
	}

	// printGrid(x, y)

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

func printGrid(x, y int) {
	fmt.Println()
	for _, line := range text {
		for _, char := range line {
			// if i == y && j == x {
			// fmt.Printf("X")
			// } else {
			fmt.Printf("%c", char)
			// }
		}
		fmt.Println()
	}
	fmt.Println()
}

func printLoopGrid() {
	fmt.Println()
	for _, line := range isLoopGrid {
		line := string(line)
		line = strings.ReplaceAll(line, "F", "┌")
		line = strings.ReplaceAll(line, "J", "┘")
		line = strings.ReplaceAll(line, "L", "└")
		line = strings.ReplaceAll(line, "7", "┐")
		line = strings.ReplaceAll(line, "|", "|")
		for _, char := range line {
			// if char == 1 {
			// 	fmt.Printf("X")
			// } else if char == 2 {
			// 	fmt.Printf(" ")
			// } else {
			// 	fmt.Printf(".")
			// }
			fmt.Printf("%c", char)
		}
		fmt.Printf("P\n")
	}
	fmt.Println()
}

func trimGrid() {
	return
	for idY, line := range isLoopGrid {
		foundX := false
		for idX, char := range line {
			if char != ' ' {
				foundX = true
				break
			}
			if !foundX {
				isLoopGrid[idY][idX] = ' '
			}
		}
	}

	// for idY, line := range isLoopGrid {
	// 	for x := len(line) - 1; x >= 0; x-- {
	// 		foundX := false
	// 		if isLoopGrid[idY][x] == 1 {
	// 			foundX = true
	// 			break
	// 		}
	// 		if !foundX {
	// 			isLoopGrid[idY][x] = 2
	// 		}
	// 	}
	// }
}
