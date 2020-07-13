package main

import (
	"fmt"
)

var boilingF = 212.0 //package level variable

func main() {
	var one int
	var two int = 2

	s := "Furkan"
	n := 5

	h, j, k, l := true, "Furkan", 2.05, 41

	fmt.Printf("%d\n%d\n", one, two)
	fmt.Printf("%s\n%d\n", s, n)

	fmt.Println(h, j, k, l)

	//in, err := os.Open("somefile")
	//out, err := os.Create("someotherfile")
}
