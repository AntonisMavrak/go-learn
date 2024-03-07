package main

import "fmt"

func main() {
	x := 10

	fmt.Println(x)  // Value of x
	fmt.Println(&x) // Address where x is stored

	changeValue(&x)
}

func changeValue(x *int) {
	*x = 7
}
