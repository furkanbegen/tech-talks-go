package main

import "fmt"

func main() {
	l, c := drivers([]string{"Hamilton", "Vettel", "Verstappen"})
	fmt.Println(l)
	fmt.Println(c)
}

func drivers(s []string) (int, int) {
	return len(s), cap(s)
}
