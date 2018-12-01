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
	var total int64
	for _, line := range input {
		number, _ := strconv.ParseInt(line, 10, 64)
		total = total + number
	}
	return total
}

func main() {
	fmt.Printf("Part 1 Answer: %v\n", part1())
}
