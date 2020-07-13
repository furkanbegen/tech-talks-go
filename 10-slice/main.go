package main

import "fmt"

func main() {
	langs := []string{}
	fmt.Println("len:", len(langs)) // 0
	fmt.Println("cap:", cap(langs)) // 0

	langs = append(langs, "Go")
	fmt.Println(langs)

}
