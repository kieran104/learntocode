package main

import (
	"fmt"
)

func main() {
	s := "H"

	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)

	n := bs[0]
	fmt.Println(n)
	fmt.Printf("type is : %T\n", n)         // type
	fmt.Printf("binary is : %b\n", n)       // binary
	fmt.Printf("hexa decimal is: %#x\n", n) // hexa decimal
}
