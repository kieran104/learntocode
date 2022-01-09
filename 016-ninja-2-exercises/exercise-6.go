package main

import "fmt"

func main() {

	const (
		a = 2019 + iota
		b
		c
		d
	)

	fmt.Println(a, b, c, d)
}
