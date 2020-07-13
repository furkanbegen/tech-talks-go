package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello, world"
	a := "Hello, 世界"

	fmt.Println(len(s))
	fmt.Println(len(a))                    // "13"
	fmt.Println(utf8.RuneCountInString(a)) // "9"

	fmt.Println(s[0], s[7]) // "104 119"  ('h' ve 'w')
	// string in boyutundan daha büyük bir değere erişmek istediğimiz için hata (panic) verir
	fmt.Println(a[14])

	fmt.Println(s[0:5])
	fmt.Println(s[5:])
}
