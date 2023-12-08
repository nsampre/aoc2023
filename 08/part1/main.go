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

	for _, line := range text[2:] {
		var left, right string

		spew.Dump(line)
		matches := LRregex.FindStringSubmatch(line)
		step := matches[1]
		left = matches[2]
		right = matches[3]

		instructions[step] = RL{left, right}
	}

	spew.Dump(instructions)

	step := "AAA"
	ins := instructions[step]
	count := 0

	for {
		letter := path[count%(len(path))]
		fmt.Printf("Letter: %c\n", letter)

		if letter == 'L' {
			step = ins.Left
		}
		if letter == 'R' {
			step = ins.Right
		}
		count++

		if step == "ZZZ" {
			break
		}

		fmt.Printf("%c: %s\n", letter, step)

		ins = instructions[step]
		spew.Dump(ins)
	}

	spew.Dump(count)
}
