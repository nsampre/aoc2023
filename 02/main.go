package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

var (
	r     = regexp.MustCompile(`Game (\d+):(.*)`)
	red   = regexp.MustCompile(`(\d+) red`)
	green = regexp.MustCompile(`(\d+) green`)
	blue  = regexp.MustCompile(`(\d+) blue`)
)

type Subset struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	Number        int
	Subsets       []string
	SubsetsParsed []Subset
	Power         int
}

func main() {
	lines := readLines("input.txt")

	games := extractGames(lines)

	spew.Dump(games)

	// countPossible(games) // part 1

	countPower(games) // part 2
}

func readLines(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func extractGames(lines []string) []Game {
	var games []Game
	for _, line := range lines {
		if r.MatchString(line) {
			number, _ := strconv.Atoi(r.FindStringSubmatch(line)[1])
			subsetsRaw := r.FindStringSubmatch(line)[2]
			subsets := strings.Split(subsetsRaw, ";")
			subsetParsed := []Subset{}
			for _, s := range subsets {
				redCount := 0
				greenCount := 0
				blueCount := 0
				if red.MatchString(s) {
					redCount, _ = strconv.Atoi(red.FindStringSubmatch(s)[1])
				}
				if green.MatchString(s) {
					greenCount, _ = strconv.Atoi(green.FindStringSubmatch(s)[1])
				}
				if blue.MatchString(s) {
					blueCount, _ = strconv.Atoi(blue.FindStringSubmatch(s)[1])
				}
				subsetParsed = append(subsetParsed, Subset{
					Red:   redCount,
					Green: greenCount,
					Blue:  blueCount,
				})

			}
			games = append(games, Game{
				Number:        number,
				Subsets:       subsets,
				SubsetsParsed: subsetParsed,
			})
		}
	}
	return games
}

func countPossible(games []Game) {
	count := 0
	for _, g := range games {
		possible := true
		for _, s := range g.SubsetsParsed {
			if s.Red > 12 || s.Green > 13 || s.Blue > 14 {
				possible = false
				break
			}
		}
		if possible {
			count += g.Number
		}
	}
	println(count)
}

func countPower(games []Game) {
	total := 0
	for _, g := range games {
		minRed := 0
		minBlue := 0
		minGreen := 0
		for _, s := range g.SubsetsParsed {
			if s.Red > minRed {
				minRed = s.Red
			}
			if s.Green > minGreen {
				minGreen = s.Green
			}
			if s.Blue > minBlue {
				minBlue = s.Blue
			}
		}
		total += minRed * minBlue * minGreen
	}
	println(total)
}
