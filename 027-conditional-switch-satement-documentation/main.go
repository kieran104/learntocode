package main

import (
	"fmt"
)

func main() {

	switch { // if no expression is in the switch statement then the default value is true
	case false:
		fmt.Println("will not print because case == false and the switch statement is looking for something that evaluates to true")
	case true:
		fmt.Println("this will print")
	}
}
