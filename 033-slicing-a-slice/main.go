package main

import (
	"fmt"
)

func main() {
	x := []int{70, 20, 44, 51, 22}
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[1:3]) // gives the values between these index positions

	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i <= 4; i++ { // doing the same thing as the range loop
		fmt.Println(i, x[i]) // loop increases the index and prints the value at the index
	}
}
