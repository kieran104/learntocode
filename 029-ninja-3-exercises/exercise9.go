package main

import (
	"fmt"
)

func main() {

	var favSport string = "cricket"

	switch favSport {
	case "basketball":
		fmt.Println("favorite sport is basketball")
	case "cricket":
		fmt.Println("favorite sport is cricket")
	case "football":
		fmt.Println("favorite sport is football")
	}
}
