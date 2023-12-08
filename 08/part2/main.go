package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/davecgh/go-spew/spew"
)

var text []string

type RL struct {
	Left  string
	Right string
}

var instructions = make(map[string]RL)

var LRregex = regexp.MustCompile(`(.*) = \((.*), (.*)\)`)

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

	path := text[0]

	endsWithA := []string{}

	for _, line := range text[2:] {
		var left, right string

		spew.Dump(line)
		matches := LRregex.FindStringSubmatch(line)
		step := matches[1]
		left = matches[2]
		right = matches[3]

		if step[len(step)-1] == 'A' {
			endsWithA = append(endsWithA, step)
		}

		instructions[step] = RL{left, right}
	}

	spew.Dump(instructions)

	var ins []RL
	var res []int

	for _, step := range endsWithA {
		ins = append(ins, instructions[step])
	}

	for i := 0; i < len(ins); i++ {
		count := 0
		for {
			letter := path[count%(len(path))]

			step := ""
			if letter == 'L' {
				step = ins[i].Left
			}
			if letter == 'R' {
				step = ins[i].Right
			}

			count++

			if step[2] == 'Z' {
				fmt.Printf("Circular path %d/%d ends on Z at %d\n", i+1, len(ins), count)
				res = append(res, count)
				break
			}

			ins[i] = instructions[step]
		}
	}

	spew.Dump(res)

	// Least common multiplier -> fast way
	lcmResult := res[0]
	for _, num := range res[1:] {
		lcmResult = lcm(lcmResult, num)
	}

	spew.Dump(lcmResult)
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func lcm(x, y int) int {
	return x * y / gcd(x, y)
}
