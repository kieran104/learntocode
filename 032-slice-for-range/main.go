package main

import (
	"fmt"
)

func main() {
	x := []int{70, 20, 44, 51, 22}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	for i, v := range x {
		if v == 44 {
			fmt.Println("")
			fmt.Println(i, v)
		} else {
			fmt.Println(v)
		}
	}
}
