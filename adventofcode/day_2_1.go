package main

import (
	"adventofcode/readinput"
	"fmt"
)

var scoreMap = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func determinePoints(pair string) int {
	switch pair[2] {
	case 'X':
		if pair[0] == 'C' {
			return 6
		} else if pair[0] == 'A' {
			return 3
		} else {
			return 0
		}
	case 'Y':
		if pair[0] == 'A' {
			return 6
		} else if pair[0] == 'B' {
			return 3
		} else {
			return 0
		}
	case 'Z':
		if pair[0] == 'B' {
			return 6
		} else if pair[0] == 'C' {
			return 3
		} else {
			return 0
		}
	}
	return 0
}

func main() {
	scanner := readinput.ReadInput("day2.txt")
	totalScore := 0
	for scanner.Scan() {
		s := scanner.Text()
		currentScore := determinePoints(s)
		totalScore += currentScore + scoreMap[string(s[2])]
	}
	fmt.Printf("Total score is: %d\n", totalScore)
}
