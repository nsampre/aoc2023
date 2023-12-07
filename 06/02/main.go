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

type raceDef struct {
	Time     int64
	Distance int64
}

var race raceDef

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
	t, _ := strconv.ParseInt(timesText[0], 10, 64)
	d, _ := strconv.ParseInt(distancesText[0], 10, 64)
	race = raceDef{
		Time:     int64(t),
		Distance: int64(d),
	}

	spew.Dump(race)

	//	Remains = (race.Time - x)
	//	Dist = (race.Time - x) * x
	//  Dist = x^2 - x*race.Time
	//  Dist > race.Distance => x^2 - x*race.Time > race.Distance
	//  Dist > race.Distance => x^2 - x*race.Time - race.Distance > 0

	// x1 = (-race.Time + sqrt(race.Time^2 - 4*race.Distance)) / 2
	// x2 = (-race.Time - sqrt(race.Time^2 - 4*race.Distance)) / 2

	// x1 = (-53717880 + sqrt(53717880^2 - 4*275118112151524)) / 2
	// x2 = (-53717880 - sqrt(53717880^2 - 4*275118112151524)) / 2

	// x1 = -5733492
	// x2 = -47984387

	// x1 - x2 = 42250895

	total := 1

	beaten := 0
	for x := int64(0); x < race.Time; t++ {
		remains := race.Time - x
		distance := x * remains

		if distance > race.Distance {
			beaten++
		}
	}
	fmt.Printf("Max distance: %d\n", beaten)
	total *= beaten
}


