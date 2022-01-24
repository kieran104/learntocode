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

	m["todd"] = 33

	v, ok := m["Barnabus"]
	if v == 0 || ok == false {
		fmt.Println("key not found")
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	ii := []int{9, 24, 11, 2}
	for i, v := range ii {
		fmt.Println(i, v)
	}
}
