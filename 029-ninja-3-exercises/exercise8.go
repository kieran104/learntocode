package main

import (
	"fmt"
)

func main() {
	h := 100

	switch {
	case 100 == 200:
		fmt.Println(h)
	case true:
		fmt.Println("true")
	default:
		fmt.Println("default answer")
	}
}
