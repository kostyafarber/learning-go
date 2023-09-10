package main

import (
	"fmt"

	BitVector "gopl.io.solutons/ch6/bitvector"
)

type testOperationPair struct {
	nums1     []int
	nums2     []int
	operation string
}

func printResults(operation string, set BitVector.IntSet) {
	fmt.Printf("Running %s of both sets\n", operation)
	fmt.Printf("Set one is now the %s of set one and two: %v\n\n", operation, set)
}

func testOperations(i int, pair testOperationPair) {
	var set1 BitVector.IntSet
	var set2 BitVector.IntSet

	nums1 := pair.nums1
	nums2 := pair.nums2

	set1.AddAll(nums1...)
	set2.AddAll(nums2...)

	fmt.Printf("Testing operation: %s\n", pair.operation)

	fmt.Printf("Set %d contains elements: %v\n", i, set1)
	fmt.Printf("Set two contains elements: %v\n", set2)

	switch pair.operation {
	case "intersection":
		set1.IntersectWith(&set2)
		printResults(pair.operation, set1)

	case "difference":
		set1.DifferenceWith(&set2)
		printResults(pair.operation, set1)

	case "symmetric_difference":
		set1.SymmetricDifference(&set2)
		printResults(pair.operation, set1)

	}
}
func main() {

	tests := []testOperationPair{
		{[]int{1, 2, 3, 4, 5}, []int{2, 3, 6}, "intersection"},
		{[]int{1, 2, 3, 4, 5}, []int{2, 3, 6}, "difference"},
		{[]int{1, 2, 3}, []int{3, 4}, "symmetric_difference"},
	}

	for i, test := range tests {
		testOperations(i, test)
	}
}
