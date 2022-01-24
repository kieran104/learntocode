package main

import "fmt"

func main() {
	a := []int{5, 222, 54, 11, 69, 88, 99, 42, 2, 1}
	for i := range a {
		fmt.Println(a[i])
	}
	fmt.Printf("type of slice is: %T", a)
}
