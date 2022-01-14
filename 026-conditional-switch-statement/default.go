package main

import (
	"fmt"
)

func main() {

	n := "Bond"
	switch n {
	case "Miss Moneypenny", n, "Do No":
		fmt.Println("Money Penny")
	case "M": // "Bond" == true in this case
		fmt.Println("M")
	case "Q":
		fmt.Println("Q")
	default:
		fmt.Println("this is default")
	}
}
