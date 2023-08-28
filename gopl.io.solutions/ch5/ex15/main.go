package main

import "fmt"

func max(ints ...int) int {
	max := 0
	for _, num := range ints {
		if num >= max {
			max = num
		}
	}
	return max
}

func min(ints ...int) int {
	min := ints[0]
	for _, num := range ints {
		if num <= min {
			min = num
		}
	}
	return min
}

func main() {
	fmt.Printf("Max is: %d\n", max(1, 2, 3, 4))
	fmt.Printf("Max is: %d\n", max(1, 2))
	fmt.Printf("Max is: %d\n", max(3, 4, 11))

	fmt.Printf("\n")

	fmt.Printf("Min is: %d\n", min(1, 2, 3, 4))
	fmt.Printf("Min is: %d\n", min(1, 2))
	fmt.Printf("Min is: %d\n", min(3, 4, 11))
}
