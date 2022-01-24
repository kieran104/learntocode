package main

import "fmt"

func main() {
	a := [5]int{5, 10, 15, 20, 25}
	for i := range a {
		fmt.Println(a[i])
	}
	fmt.Printf("type of array is: %T", a)
}
