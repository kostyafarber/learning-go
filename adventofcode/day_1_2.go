package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("day1.txt")
	var calories []int

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	var current_calories = 0
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) == 0 {
			calories = append(calories, current_calories)
			current_calories = 0
		} else {
			s_int, err := strconv.Atoi(s)
			current_calories += s_int
			if err != nil {
				panic(err)
			}
		}
	}
	sort.Ints(calories)
	top_1 := calories[len(calories)-1]
	top_2 := calories[len(calories)-2]
	top_3 := calories[len(calories)-3]
	sum := top_1 + top_2 + top_3

	fmt.Printf("Largest three calories summed are: %d\n", sum)
}
