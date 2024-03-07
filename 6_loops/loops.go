package main

func main() {

	// For loop
	for i := 1; i <= 10; i++ {
		println(i)
	}

	// While loop
	j := 1
	for j <= 10 {
		println(j)
		j++
	}

	// Infinite loop
	// for {
	// 	println("Infinite loop")
	// }

	// Break and continue
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		if i < 3 {
			continue
		}
		println(i)
	}

}
