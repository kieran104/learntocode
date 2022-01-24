package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 21,
	}

	fmt.Println(m)

	delete(m, "James")

	fmt.Println(m)
}
