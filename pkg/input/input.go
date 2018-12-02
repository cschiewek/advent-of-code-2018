package input

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Retrieve Gets the input for the given day
func Retrieve(day string) (nonEmptyLines []string) {
	path := fmt.Sprintf("input/day%v.txt", day)
	input, _ := ioutil.ReadFile(path)
	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		if line != "" {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}
	return
}
