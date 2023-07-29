package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func popCount(b byte) int {
	count := 0
	for ; b != 0; count++ {
		b &= b - 1
	}
	return count
}

func main() {
	args := os.Args
	if len(args)-1 > 2 || len(args) == 1 {
		fmt.Fprintf(os.Stderr, "usage: %s s1 s2\n", args[0])
		os.Exit(1)
	}

	s1 := sha256.Sum256([]byte(os.Args[1]))
	s2 := sha256.Sum256([]byte(os.Args[2]))
	fmt.Printf("%x\n%x\n", s1, s2)

	diffBits := 0
	for i := range s1 {
		mask := s1[i] ^ s2[i]
		diffBits += popCount(mask)
	}
	fmt.Println(diffBits)
}
