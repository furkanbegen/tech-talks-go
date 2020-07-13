package main

import "fmt"

// Matrix is a 3x3 array, really an array of arrays.
type Matrix [3][3]int

func main() {
	m := Matrix{
		{0, 0, 0},
		{1, 1, 1},
		{2, 2, 3},
	}

	fmt.Println(m)
}
