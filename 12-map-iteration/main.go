package main

import "fmt"

func main() {
	drivers := map[string]string{
		"Hamilton":   "Mercedes-AMG Patronas",
		"Verstappen": "Aston Martin Redbull",
		"Vettel":     "Scuderia Ferrari",
		"Norris":     "McLaren",
	}

	// Map unsorted list olduğu için iteration her zaman aynı sıra ile olmaz.
	for key, value := range drivers {
		fmt.Printf("%s drives for: %s\n", key, value)
	}
}
