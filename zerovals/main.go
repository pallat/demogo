package main

import "fmt"

func main() {
	var s string
	var i int
	var b bool

	fmt.Printf("type of s: %6T, zero value of s: %q\n", s, s)
	fmt.Printf("type of i: %6T, zero value of i: %d\n", i, i)
	fmt.Printf("type of b: %6T, zero value of b: %t\n", b, b)
}
