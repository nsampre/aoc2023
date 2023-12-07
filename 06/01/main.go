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
var numbersRegex = regexp.MustCompile(`\d+`)

type race struct {
	Time     int
	Distance int
}

var races []race

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

	timesText := numbersRegex.FindAllString(text[0], -1)
	distancesText := numbersRegex.FindAllString(text[1], -1)
	for _, timeText := range timesText {
		n, _ := strconv.ParseInt(timeText, 10, 64)
		races = append(races, race{
			Time: int(n),
		})
	}
	for idx, distancesText := range distancesText {
		n, _ := strconv.ParseInt(distancesText, 10, 64)
		races[idx].Distance = int(n)
	}

	total := 1
	for _, race := range races {
		beaten := 0
		for t := 0; t < race.Time; t++ {
			hold := t
			remains := race.Time - hold

			distance := hold * remains
			if distance > race.Distance {
				beaten++
			}
		}
		fmt.Printf("Max distance: %d\n", beaten)
		total *= beaten
	}

	spew.Dump(total)
}
