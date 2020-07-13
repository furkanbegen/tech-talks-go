package main

import "fmt"

// Driver represents a Formula 1 driver
type Driver struct {
	Name string
	Team string
}

func (d Driver) String() string {
	return fmt.Sprintf("%s drives for %s", d.Name, d.Team)
}

func main() {
	d := Driver{Name: "Hamilton", Team: "Mercedes-Amg Patronas"}
	fmt.Println(d.String())
}
