package main

import "fmt"

func main() {
	var maxUint32 uint32 = 4294967295
	//var maxUint32Plus uint32 = 4294967296
	fmt.Println(maxUint32)

	fmt.Println(maxUint32 + 1) //wraparound
}
