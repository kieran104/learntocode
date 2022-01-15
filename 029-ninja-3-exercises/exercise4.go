package main

import (
	"fmt"
)

func main() {

	years := 1993

	for {
		if years > 2022 {
			break
		}
		fmt.Println(years)
		years++
	}
}
