package main

import (
	"fmt"
)

func main() {

	years := 1993

	for years <= 2022 {
		fmt.Println(years)
		years++ // years ++ needs to be at the end or the original number wont print
	}
}
