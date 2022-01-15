package main

import (
	"fmt"
)

func main() {

	var sentence string = "Loser"

	if sentence == "fame" {
		fmt.Printf("Variable sentence is equal to '%v'.", sentence)
	} else if sentence == "Winner" {
		fmt.Printf("Variable sentence is equal to %v", sentence)
	} else {
		fmt.Println("Variable sentence is not equal to 'Loser' or 'Winner'")
	}
}
