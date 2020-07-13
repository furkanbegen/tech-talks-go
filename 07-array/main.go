package main

import "fmt"

func main() {

	numbers := [3]int{}
	fruits := [2]string{"Elma", "Armut"}

	fmt.Printf("numbers: %d\n", numbers)
	fmt.Printf("fruites: %s\n", fruits)

	fmt.Printf("Fruits has %d elements\n", len(fruits))

	s1 := [2]string{"Elma", "Armut"}
	s2 := [2]string{"Elma", "Armut"}

	if s1 == s2 {
		fmt.Println("2 array eşit")
	}

}
