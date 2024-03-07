package main

import "fmt"

func main() {

	// Different ways to declare variables
	var a int = 5
	var b float32 = 4.32
	const pi float64 = 3.14159265

	var (
		c = 8
		d = 9
	)

	x, y := 14, 15

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c, d, x, y)
	fmt.Println(pi)

}
