package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 21,
	}

	fmt.Println(m)

	fmt.Println(m["James"])

	fmt.Println(m["Barnabus"])

	v, ok := m["Barnabus"]
	if v == 0 || ok == false {
		fmt.Println("key not found")
	}

}
