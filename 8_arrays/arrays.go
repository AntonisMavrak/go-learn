package main

import "fmt"

func main() {

	//var EvenNum [5]int
	//
	//EvenNum[0] = 0
	//EvenNum[1] = 2
	//EvenNum[2] = 4
	//EvenNum[3] = 6
	//EvenNum[4] = 8

	EvenNum := [5]int{0, 2, 4, 6, 8}

	for _, value := range EvenNum {
		fmt.Println(value)
	}

	// Slice array
	numSlice := []int{5, 4, 3, 2, 1}

	sliced := numSlice[3:5]
	//sliced := numSlice[:5]
	//sliced := numSlice[0:]
	fmt.Println(sliced)

	slice2 := make([]int, 5, 10)
	copy(slice2, numSlice)
	fmt.Println(numSlice)

	slice3 := append(slice2, 0, -1)
	fmt.Println(slice3)
}
