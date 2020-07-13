package main

import "fmt"

func main() {
	ages := map[string]int{
		"ali":  32,
		"veli": 34,
	}

	fmt.Println(ages)

	emails := map[string]string{
		"ali@finartz.com": "Ali",
		"veli@finarz.com": "Veli",
	}

	fmt.Println(emails)

	var emails2 map[string]string
	emails2 = make(map[string]string)
	emails2["ali@finartz.com"] = "Ali"
	emails2["veli@finartz.com"] = "Veli"

	fmt.Println(emails2)

}
