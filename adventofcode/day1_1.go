package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
* Find the Elf carrying the most Calories.
* How many total Calories is that Elf carrying?
*
 */

func main() {
	f, err := os.Open("day1.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	var largest_calories = 0
	var current_calories = 0

	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 {
			if current_calories > largest_calories {
				largest_calories = current_calories
			}

			current_calories = 0
		} else {
			s_int, err := strconv.Atoi(s)

			if err != nil {
				panic(err)
			}

			current_calories += s_int
		}
	}

	fmt.Printf("Max calories are: %d\n", largest_calories)
}
