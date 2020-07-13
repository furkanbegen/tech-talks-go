package main

import "fmt"

func main() {
	names := [4]string{"Go", "Java", "Php", "Ruby"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	//Range array in index ve value değerlerini döner
	for i, n := range names {
		fmt.Printf("%d - %s\n", i, n)
	}

	/*for i := 0; i < len(names); i++ {
		if i == 2 {
			// go to the start of the loop
			continue
		}
		// do work
		fmt.Println(names[i])
	}*/

	/*for i := 0; i < len(names); i++ {
		if i == 2 {
			// stop looping
			break
		}
		// do work
		fmt.Println(names[i])
	}*/

}
