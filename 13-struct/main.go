package main

import "fmt"

func main() {
	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	u := User{
		Name:  "Furkan Beğen",
		Email: "fbegen@finartz.com",
	}

	fmt.Println(u)

}
