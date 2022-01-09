package main

import (
	"fmt"
)

func main() {

	type hungry int
	var x hungry

	fmt.Printf("%v\t%T\n", x, x)

	x = 42

	fmt.Printf("%v\n", x)

	var y int
	y = int(x)

	fmt.Printf("%v\n%T", y, y)
}
