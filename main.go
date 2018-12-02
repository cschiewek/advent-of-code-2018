package main

import (
	"fmt"
	"os"

	"./pkg/day1"
	"./pkg/day2"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "1":
			fmt.Printf("Day 1 Part 1: %v\n", day1.Part1())
			fmt.Printf("Day 1 Part 1: %v\n", day1.Part2())
		case "2":
			fmt.Printf("Day 2 Part 1: %v\n", day2.Part1())
		default:
			println("Not a valid day")
		}
	} else {
		println("You must pass a day as a param")
	}
}
