package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args
	if len(args)-1 > 2 || len(args) == 1 {
		fmt.Fprintf(os.Stderr, "usage: %s s1 s2\n", args[0])
		os.Exit(1)
	}

	s1 := args[1]
	s2 := args[2]

	if len(s1) != len(s2) {
		fmt.Println(false)
		os.Exit(0)
	}

	fmt.Printf("%t\n", insertionSort(s1) == insertionSort(s2))
}

/*
i ← 1
while i < length(A)

	j ← i
	while j > 0 and A[j-1] > A[j]
	    swap A[j] and A[j-1]
	    j ← j - 1
	end while
	i ← i + 1

end while
*/
func insertionSort(s string) string {
	sorted := strings.Split(s, "")
	i := 1
	for i < len(s) {
		j := i
		for j > 0 && sorted[j-1] > sorted[j] {
			sorted[j], sorted[j-1] = sorted[j-1], sorted[j]
			j--
		}
		i++
	}
	return strings.Join(sorted, "")
}
