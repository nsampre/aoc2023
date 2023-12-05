package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

var text []string
var numbersRegex = regexp.MustCompile(`\d+`)
var seeds []int64

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

	seedsText := numbersRegex.FindAllString(text[0], -1)
	for _, seed := range seedsText {
		n, _ := strconv.ParseInt(seed, 10, 64)
		seeds = append(seeds, n)
	}

	var seedToSoilMap [][]int64
	var soilToFertilizerMap [][]int64
	var fertilizerToWaterMap [][]int64
	var waterToLightMap [][]int64
	var lightToTemperatureMap [][]int64
	var temperatureToHumidityMap [][]int64
	var humidityToLocationMap [][]int64

	results := []int64{}

	for idxLine, line := range text[1:] {
		switch line {
		case "seed-to-soil map:":
			seedToSoilMap = xToXMap(idxLine)
		case "soil-to-fertilizer map:":
			soilToFertilizerMap = xToXMap(idxLine)
		case "fertilizer-to-water map:":
			fertilizerToWaterMap = xToXMap(idxLine)
		case "water-to-light map:":
			waterToLightMap = xToXMap(idxLine)
		case "light-to-temperature map:":
			lightToTemperatureMap = xToXMap(idxLine)
		case "temperature-to-humidity map:":
			temperatureToHumidityMap = xToXMap(idxLine)
		case "humidity-to-location map:":
			humidityToLocationMap = xToXMap(idxLine)
		}
	}

	seeds := multiplySeeds(seeds)
	for _, seed := range seeds {
		for _, mappings := range [][][]int64{seedToSoilMap, soilToFertilizerMap, fertilizerToWaterMap, waterToLightMap, lightToTemperatureMap, temperatureToHumidityMap, humidityToLocationMap} {
			for _, mapping := range mappings {
				if o := offset(seed, mapping); o != 0 {
					seed += o
					break
				}
			}
		}
		results = append(results, seed)
	}

	min := results[0]
	for _, result := range results {
		if result < min {
			min = result
		}
	}

	spew.Dump(min)
}

func multiplySeeds(seeds []int64) []int64 {
	newSeeds := []int64{}
	for i := int64(0); i < int64(len(seeds)); i += 2 {
		fmt.Println(seeds[i], seeds[i+1])
		for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
			newSeeds = append(newSeeds, j)
		}
	}
	return newSeeds
}

func offset(seed int64, mapping []int64) int64 {
	sourceRangeStart := mapping[1]
	sourceRangeEnd := mapping[1] + mapping[2]
	if seed >= sourceRangeStart && seed < sourceRangeEnd {
		return mapping[0] - sourceRangeStart
	}
	return int64(0)
}

func xToXMap(idxBegin int) [][]int64 {
	var mapping [][]int64

	for _, line := range text[idxBegin+2:] {
		if line == "" {
			break
		}

		numbers := []int64{}
		numbersText := numbersRegex.FindAllString(line, -1)
		for _, number := range numbersText {
			n, _ := strconv.ParseInt(number, 10, 64)
			numbers = append(numbers, n)
		}
		mapping = append(mapping, numbers)
	}

	sort.Slice(mapping, func(i, j int) bool {
		return mapping[i][1] < mapping[j][1]
	})

	return mapping
}
