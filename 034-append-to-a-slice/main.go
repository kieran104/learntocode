package main

import (
	"fmt"
)

func main() {
	x := []int{70, 20, 44, 51, 22}
	fmt.Println(x)
	x = append(x, 22222, 44444, 55555, 66666)
	fmt.Println(x)

	y := []int{21, 768, 999}
	x = append(x, y...)
	fmt.Println(x)
}
