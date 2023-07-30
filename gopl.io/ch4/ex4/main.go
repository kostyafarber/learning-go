package main

import "fmt"

func rotateLeft(s []int, r int) []int {
	h1 := s[:r]
	s = append(s, h1...)
	s = s[r:]
	return s
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	r := rotateLeft(s, 1)
	fmt.Println(r)
}
