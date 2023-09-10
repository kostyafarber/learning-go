package main

import (
	"fmt"

	BitVector "gopl.io.solutons/ch6/bitvector"
)

func main() {
	var x = BitVector.IntSet{}
	x.Add(5)
	x.Add(20)
	x.Add(1440)
	x.Add(3)

	// length
	fmt.Printf("Length of x: %s is: %d\n", x, x.Len())

	var y BitVector.IntSet
	y.Add(1000000)

	fmt.Printf("Length of y: %s is: %d\n", y, y.Len())

	// remove
	fmt.Println("Removing digit 5 from x")
	x.Remove(5)
	fmt.Println(x)

	fmt.Println("Removing digit 1 (should have no effect)")
	x.Remove(1)
	fmt.Println(x)

	// clear
	fmt.Println("Clearing set y")
	y.Clear()
	fmt.Println(y)

	// copy
	fmt.Println("Copying x")
	arr := x.Copy()

	fmt.Printf("Copied %s into %v\n", x, arr)

}
