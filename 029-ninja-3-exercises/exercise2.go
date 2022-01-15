package main

import (
	"fmt"
)

func main() {

	for a := 'A'; a <= 'Z'; a++ {
		fmt.Println(" ")
		for j := 0; j < 3; j++ {
			fmt.Printf("%#U\n", a)
		}
	}
}
