package main

import (
	"fmt"

	BitVector "gopl.io.solutons/ch6/bitvector"
)

func main() {
	var x BitVector.IntSet
	nums := []int{1, 2, 3, 4, 7}

	fmt.Printf("Adding nums: %v to empty set %s\n", nums, x)
	x.AddAll(nums...)
	fmt.Println(x)

}
