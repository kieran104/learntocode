package main

import (
	"fmt"
)

func main() {
	x := 42 // first time you use variable you have to declare it, using the short delcaration.
	fmt.Println(x)
	x = 99 // only need to use "=" to assign a new value to a variable that's already been declared.
	fmt.Println(x)
	y := 100 + 24
	fmt.Println(y)
}
