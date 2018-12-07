package day1

import (
	"strconv"

	"../input"
)

// Part1 Calculates the answer for part 1
func Part1() int64 {
	input := input.Retrieve("1")
	var frequency int64
	for _, line := range input {
		number, _ := strconv.ParseInt(line, 10, 64)
		frequency += number
	}
	return frequency
}

// Part2 Calculates the answer for part 2
func Part2() int64 {
	input := input.Retrieve("1")
	var frequency int64
	frequencies := make(map[int64]int64)
	frequencies[0] = 1
	i := 0
	for {
		number, _ := strconv.ParseInt(input[i], 10, 64)
		frequency += number
		frequencies[frequency]++
		if frequencies[frequency] == 2 {
			return frequency
		}
		if i+1 == len(input) {
			i = 0
		} else {
			i++
		}
	}
}
