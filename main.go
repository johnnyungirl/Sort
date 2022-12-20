package main

import (
	"fmt"
)

func main() {
	var a [4]int                      // array with zero values
	b := [4]int{1, 2, 3}              // partially initialized array
	var c [4]int = [4]int{1, 2, 3, 4} // array initialization
	d := [...]int{5, 6, 7, 0}         // ... - means that array size equals the number of elements in the array literal

	fmt.Printf("a: length: %d, capacity: %d, data: %v\n", len(a), cap(a), a)
	fmt.Printf("b: length: %d, capacity: %d, data: %v\n", len(b), cap(b), b)
	fmt.Printf("c: length: %d, capacity: %d, data: %v\n", len(c), cap(c), c)
	fmt.Printf("d: length: %d, capacity: %d, data: %v\n", len(d), cap(d), d)
}
