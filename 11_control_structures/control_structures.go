// Defer
/*
	- Defer statement defers the execution of a function until the surrounding function returns.
	- multiple defers are pushed into stack and executes in Last In First Out (LIFO) order.
	- Defer generally used to cleanup resources like file database connection etc.
*/

// Recover
/*
	- Recover is a built-in function that regains control of a panicking goroutine.
	- It helps to regain the normal flow of execution after a panic.
	- Generally it used with defer statement to recover panic in goroutine.
*/

// Panic
/*
	- Panic is a built-in function that stops the ordinary flow of control and begins panicking.
	- Panic is similar to throwing an exception in other languages.
	- Generally when a panic is called then the normal execution flow stops immediately.
	- The deffered functions are executed normally.
*/

package main

import "fmt"

func main() {

	fmt.Println(div(3, 0))
	fmt.Println(div(3, 5))
	demPanic()

}

func div(num1 int, num2 int) int {

	defer func() {
		fmt.Println(recover())
	}()
	solution := num1 / num2
	return solution
}

func demPanic() {
	defer func() {
		fmt.Println(recover())
	}()
	panic("PANIC")
}
