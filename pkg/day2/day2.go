package day2

import (
	"fmt"

	"../input"
)

func characterCounts(str string) (two bool, three bool) {
	counts := make(map[string]int)
	for _, character := range str {
		counts[string(character)]++
	}
	two, three = false, false
	for _, count := range counts {
		switch count {
		case 2:
			two = true
		case 3:
			three = true
		}
	}
	return
}

// Part1 Calculate answer to day 2 part 1
func Part1() string {
	input := input.Retrieve("2")
	twos, threes := 0, 0
	for _, line := range input {
		two, three := characterCounts(line)
		if two {
			twos++
		}
		if three {
			threes++
		}
	}
	return fmt.Sprint(twos * threes)
}

func letterPosition(character string) int {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	letters := make(map[string]int)
	for index, letter := range alphabet {
		letters[string(letter)] = index + 1
	}
	return letters[string(character)]
}

func match(a string, b string) bool {
	isMatch := false
	for index, character := range a {
		if character == b[index] {
			isMatch = true
			continue
		}
	}
	return isMatch
}

// Part2 Calculate answer to day 2 part 2
func Part2() string {
	input := input.Retrieve("2")
	a, b := "", ""
	for index, line := range input {
		for i := index + 1; a == ""; i++ {
			if match(line, input[i]) {
				a, b = line, input[i]
			}
		}
	}
	return fmt.Sprintf("%v - %v", a, b)
}
