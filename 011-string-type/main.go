package main

import (
	"fmt"
)

func main() {

	s := "Hello Playground"
	fmt.Println(s)
	fmt.Printf("%T", s)

	bs := []byte(s)
	fmt.Printf("%T\n", bs)
	fmt.Println(bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}
	for i, v := range s {
		fmt.Println(i, v)
	}

}
