package main

import "fmt"

func main() {

	a := (42 == 42)
	b := (42 <= 42)
	c := (42 >= 43)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 43)

	fmt.Println(a, b, c, d, e, f)

	// if a != b {
	// 	fmt.Println("a is not equal to b")
	// } else {
	// 	fmt.Println("a is equal to b")
	// }
	// if a < b || b > a {
	// 	fmt.Println("a is greater than or less than b")
	// } else {
	// 	fmt.Println("a is not greater or less than b")
	// }
	// if a == b {
	// 	fmt.Println("a is equal to b")
	// }
}
