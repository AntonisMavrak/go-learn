package main

import "fmt"

func main() {

	x, y := 5, 6

	// Arithmetic
	fmt.Println("x + y = ", x+y)
	fmt.Println("x - y = ", x-y)
	fmt.Println("x * y = ", x*y)
	fmt.Println("x / y = ", x/y)
	fmt.Println("x mod y = ", x%y)

	// Booleans
	isBool := true
	hate := false

	fmt.Println(isBool && hate)
	fmt.Println(isBool || hate)
	fmt.Println(!isBool)
}
