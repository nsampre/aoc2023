package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// read input file line by line

	lines := readFile("input")
	numbers := numbers(lines)

	total := 0
	for _, n := range numbers {
		total += n
	}

	fmt.Println(total)
}

func readFile(filename string) []string {
	// open file
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// read file line by line
	scanner := bufio.NewScanner(f)
	lines := []string{}
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		lines = append(lines, scanner.Text())
	}
	return lines
}

func reverString(s string) string {
	// reverse string
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {

		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func numbers(lines []string) []int {
	// convert lines to numbers
	numbers := make([]int, len(lines))
	for idxL, line := range lines {
		for _, char := range line {
			if char >= '0' && char <= '9' {
				numbers[idxL] = 10 * int(char-'0')
				break
			}
		}
		for _, char := range reverString(line) {
			if char >= '0' && char <= '9' {
				numbers[idxL] += int(char - '0')
				break
			}
		}
		fmt.Println(numbers[idxL])
	}
	return numbers
}
