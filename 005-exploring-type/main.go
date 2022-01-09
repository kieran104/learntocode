package main

import "fmt"

var y = 42

// DECLARED the VARIABLE with the IDENTIFIER "z" is TYPE string
// and I ASSIGN the VALUE "shaken, not stirred"

var z string = "Shaken not stirred"

// this is a STATIC programming language
// a VARIABLE is DECLARED to hold a certain TYPE
// not a DYNAMIC programming language

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)

}
