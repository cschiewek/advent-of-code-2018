package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func retrieveInput() []string {
	input, _ := ioutil.ReadFile("day1.input")
	lines := strings.Split(string(input), "\n")
	return lines
}

func part1() int64 {
	input := retrieveInput()
	var frequency int64
	for _, line := range input {
		number, _ := strconv.ParseInt(line, 10, 64)
		frequency += number
	}
	return frequency
}

func part2() int64 {
	input := retrieveInput()
	var frequency int64
	frequencies := make(map[int64]int64)
	frequencies[0] = 1
	i := 0
	for {
		if input[i] != "" {
			number, _ := strconv.ParseInt(input[i], 10, 64)
			frequency += number
			frequencies[frequency]++
			if frequencies[frequency] == 2 {
				return frequency
			}
		}
		if i+1 == len(input) {
			i = 0
		} else {
			i++
		}
	}
}

func main() {
	fmt.Printf("Part 1 Answer: %v\n", part1())
	fmt.Printf("Part 2 Answer: %v\n", part2())
}
