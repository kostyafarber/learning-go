package main

import (
	"fmt"

	BitVector "gopl.io.solutons/ch6/bitvector"
)

func main() {
	var set BitVector.IntSet
	nums := []int{2, 4, 13, 24, 95}
	fmt.Printf("Adding %v to the set\n", nums)
	set.AddAll(nums...)

	fmt.Printf("Testing method Elems by iterating with range on the set\n\n")

	for i, num := range set.Elems() {
		fmt.Printf("Number %d is: %d\n", i+1, num)
	}
}
