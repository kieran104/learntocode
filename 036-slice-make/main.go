package main

import (
	"fmt"
)

func main() {
	x := make([]int, 10, 10)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 4
	x[1] = 8

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3342)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x)) // capacity is now doubled because we appended over the original cap of 10.
}
