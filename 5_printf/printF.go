package main

import "fmt"

func main() {

	var name string = "John"
	const pi float64 = 3.14159265
	win := true
	x := 5

	fmt.Println(len(name))
	fmt.Println(name + " is a good guy")

	fmt.Printf("%f \n", pi)   // %f is a format specifier for float
	fmt.Printf("%.3f \n", pi) // %.3f is a format specifier that specifies the number of decimal places
	fmt.Printf("%T \n", pi)   // %T is a format specifier for the data type of value
	fmt.Printf("%t \n", win)  // %t is a format specifier for boolean values
	fmt.Printf("%d \n", x)    // %d is a format specifier for integers
	fmt.Printf("%b \n", 25)   // %b is a format specifier for binary representation
	fmt.Printf("%c \n", 33)   // %c is a format specifier for character representation ASCII
	fmt.Printf("%x \n", 15)   // %x is a format specifier for hexadecimal representation
	fmt.Printf("%e \n", pi)   // %e is a format specifier for scientific notation
}
