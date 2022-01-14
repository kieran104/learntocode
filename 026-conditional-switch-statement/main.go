package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print")
	case (3 == 3):
		fmt.Println("this is true")
		fallthrough
	case (4 == 4):
		fmt.Println("also true, but does it print?") // will not print without fallthrough even if the case is true
		fallthrough
	case (7 == 9):
		fmt.Println("not true 1")
		fallthrough
	case (11 == 14):
		fmt.Println("not true 2")
		fallthrough
	case (15 == 15):
		fmt.Println("true 15")
		fallthrough
		// fallthrough - don't need a fallthrough if at the last case, because theres nothing to fallthrough to.
	default: // default just prints regardless without any case?
		fmt.Println("this is default")
	}
}
