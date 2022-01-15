package main

import (
	"fmt"
)

func main() {

	var sentence string = "Loser"

	if sentence == "Loser" {
		fmt.Printf("Variable sentence is equal to '%v'.", sentence)
	} else {
		fmt.Printf("Variable sentence is not equal to %v", sentence)
	}
}
