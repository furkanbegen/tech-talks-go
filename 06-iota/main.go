package main

import "fmt"

const (
	apple int = iota
	orange
	banana
)

func main() {
	fmt.Printf("The value of Apple is %v\n", apple)
	fmt.Printf("The value of Orange is %v\n", orange)
	fmt.Printf("The value of Banana is %v\n", banana)
}
