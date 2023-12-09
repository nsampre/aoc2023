package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

var text []string

var numbersExpr = regexp.MustCompile(`-?\d+`)

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

	sum := int64(0)
	for _, line := range text {
		numbersString := numbersExpr.FindAllString(line, -1)
		numbers := make([]int64, len(numbersString))
		for i, n := range numbersString {
			numbers[len(numbers)-1-i], _ = strconv.ParseInt(n, 10, 64) // PART 2, just revert input
		}

		// spew.Dump(numbers)

		rows := make([][]int64, 1)
		rows[0] = numbers
		x := 0
		for {
			row := make([]int64, len(rows[x])-1)
			zero := true
			for i := 0; i < len(rows[x])-1; i++ {
				diff := rows[x][i+1] - rows[x][i]
				row[i] = diff
				if diff != 0 {
					zero = false
				}
			}
			rows = append(rows, row)
			x++
			if zero {
				break
			}
		}

		rows[len(rows)-1] = append(rows[len(rows)-1], 0)

		for i := len(rows) - 2; i >= 0; i-- {
			// fmt.Printf("row in: %v\n", rows[i])

			r := rows[i]
			bottomR := rows[i+1]

			diff := bottomR[len(bottomR)-1] + r[len(r)-1]

			r = append(r, diff)
			rows[i] = r

			// fmt.Printf("row out: %v\n", rows[i])

		}

		if rows[0][len(rows[0])-1] == 590384 {
			for _, r := range rows {
				fmt.Printf("row: %v\n", r)
			}
		}

		sum += rows[0][len(rows[0])-1]
		// spew.Dump(rows[0][len(rows[0])-1])
		fmt.Println(rows[0][len(rows[0])-1])
		// _ = sum
		// break
	}

	spew.Dump(sum)
}
