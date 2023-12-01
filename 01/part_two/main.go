package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		l := scanner.Text()
		lines = append(lines, l)
	}
	return lines
}

func fixText(s string) string {
	for true {
		minIndex := -1

		minIndexOne := strstrIndex(s, "one")
		minIndexTwo := strstrIndex(s, "two")
		minIndexThree := strstrIndex(s, "three")
		minIndexFour := strstrIndex(s, "four")
		minIndexFive := strstrIndex(s, "five")
		minIndexSix := strstrIndex(s, "six")
		minIndexSeven := strstrIndex(s, "seven")
		minIndexEight := strstrIndex(s, "eight")
		minIndexNine := strstrIndex(s, "nine")

		minIndex = minimumIndex([]int{minIndexOne, minIndexTwo, minIndexThree, minIndexFour, minIndexFive, minIndexSix, minIndexSeven, minIndexEight, minIndexNine})

		fmt.Println([]int{minIndexOne, minIndexTwo, minIndexThree, minIndexFour, minIndexFive, minIndexSix, minIndexSeven, minIndexEight, minIndexNine})
		fmt.Println(minIndex)

		if minIndex == 9999999999 {
			break
		}
		if minIndexOne == minIndex && minIndexOne != -1 {
			s = strings.Replace(s, "one", "1", 1)
		}
		if minIndexTwo == minIndex && minIndexTwo != -1 {
			s = strings.Replace(s, "two", "2", 1)
		}
		if minIndexThree == minIndex && minIndexThree != -1 {
			s = strings.Replace(s, "three", "3", 1)
		}
		if minIndexFour == minIndex && minIndexFour != -1 {
			s = strings.Replace(s, "four", "4", 1)
		}
		if minIndexFive == minIndex && minIndexFive != -1 {
			s = strings.Replace(s, "five", "5", 1)
		}
		if minIndexSix == minIndex && minIndexSix != -1 {
			s = strings.Replace(s, "six", "6", 1)
		}
		if minIndexSeven == minIndex && minIndexSeven != -1 {
			s = strings.Replace(s, "seven", "7", 1)
		}
		if minIndexEight == minIndex && minIndexEight != -1 {
			s = strings.Replace(s, "eight", "8", 1)
		}
		if minIndexNine == minIndex && minIndexNine != -1 {
			s = strings.Replace(s, "nine", "9", 1)
		}
	}

	return s
}

func fixTextR(s string) string {
	for true {
		minIndex := -1

		minIndexOne := strstrIndex(s, reverString("one"))
		minIndexTwo := strstrIndex(s, reverString("two"))
		minIndexThree := strstrIndex(s, reverString("three"))
		minIndexFour := strstrIndex(s, reverString("four"))
		minIndexFive := strstrIndex(s, reverString("five"))
		minIndexSix := strstrIndex(s, reverString("six"))
		minIndexSeven := strstrIndex(s, reverString("seven"))
		minIndexEight := strstrIndex(s, reverString("eight"))
		minIndexNine := strstrIndex(s, reverString("nine"))

		minIndex = minimumIndex([]int{minIndexOne, minIndexTwo, minIndexThree, minIndexFour, minIndexFive, minIndexSix, minIndexSeven, minIndexEight, minIndexNine})

		fmt.Println([]int{minIndexOne, minIndexTwo, minIndexThree, minIndexFour, minIndexFive, minIndexSix, minIndexSeven, minIndexEight, minIndexNine})
		fmt.Println(minIndex)

		if minIndex == 9999999999 {
			break
		}
		if minIndexOne == minIndex && minIndexOne != -1 {
			s = strings.Replace(s, reverString("one"), "1", 1)
		}
		if minIndexTwo == minIndex && minIndexTwo != -1 {
			s = strings.Replace(s, reverString("two"), "2", 1)
		}
		if minIndexThree == minIndex && minIndexThree != -1 {
			s = strings.Replace(s, reverString("three"), "3", 1)
		}
		if minIndexFour == minIndex && minIndexFour != -1 {
			s = strings.Replace(s, reverString("four"), "4", 1)
		}
		if minIndexFive == minIndex && minIndexFive != -1 {
			s = strings.Replace(s, reverString("five"), "5", 1)
		}
		if minIndexSix == minIndex && minIndexSix != -1 {
			s = strings.Replace(s, reverString("six"), "6", 1)
		}
		if minIndexSeven == minIndex && minIndexSeven != -1 {
			s = strings.Replace(s, reverString("seven"), "7", 1)
		}
		if minIndexEight == minIndex && minIndexEight != -1 {
			s = strings.Replace(s, reverString("eight"), "8", 1)
		}
		if minIndexNine == minIndex && minIndexNine != -1 {
			s = strings.Replace(s, reverString("nine"), "9", 1)
		}
	}

	return s
}

func minimumIndex(ints []int) int {
	min := 9999999999
	for _, v := range ints {
		if v < min && v != -1 {
			min = v
		}
	}
	return min
}

func strstrIndex(s string, substring string) int {
	// find substring index
	for i := 0; i < len(s); i++ {
		if i+len(substring) > len(s) {
			return -1
		}
		if s[i:i+len(substring)] == substring {
			return i
		}
	}
	return -1
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
	numbers := make([]int, len(lines))
	for idxL, line := range lines {
		for _, char := range fixText(line) {
			if char >= '0' && char <= '9' {
				numbers[idxL] = 10 * int(char-'0')
				break
			}
		}

		for _, char := range fixTextR(reverString(line)) {
			if char >= '0' && char <= '9' {
				numbers[idxL] += int(char - '0')
				break
			}
		}

		fmt.Println(numbers[idxL])
	}
	return numbers
}
