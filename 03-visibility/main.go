package main

import "fmt"

func main() {
	// available only inside the package
	var s string

	fmt.Println(s)
}

// Foo visible outside of the package
func Foo() string {
	return "Foo"
}
