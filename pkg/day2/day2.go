package day2

import (
	"bytes"
	"fmt"

	"../input"
)

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

// Part2 Calculate answer to day 2 part 2
func Part2() string {
	a, b := findMatch()
	var buffer bytes.Buffer
	for i := 0; i != len(a); i++ {
		if a[i] == b[i] {
			buffer.WriteByte(a[i])
		}
	}
	return fmt.Sprintf(buffer.String())
}

func findMatch() (string, string) {
	input := input.Retrieve("2")
	for index, line := range input {
		for i := index + 1; i != len(input); i++ {
			if match(line, input[i]) {
				return line, input[i]
			}
		}
	}
	return "", ""
}

func match(a string, b string) bool {
	differences := 0
	for index, character := range a {
		if string(character) != string(b[index]) {
			differences++
		}
		if differences > 1 {
			return false
		}
	}
	return true
}
