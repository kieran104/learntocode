package main

import (
	"fmt"
)

func main() {
	// x := 43 % 40
	// y := 43 / 40
	// fmt.Println(x, y)
	x := 1
	for {
		x++
		if x > 100 {
			break
		}
		if x%2 == 0 { // != would give you all even numbers
			continue
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("done")
}
