package main

import (
	"fmt"
)

func main() {
	var x [10]int // how to make an array NOT using the short delcaration with empty values
	var y [8]int  // different TYPE than x

	fmt.Println(x)

	x[5] = 244 // changing the value at index 5

	fmt.Println(x)
	fmt.Println(y)
}
