package main

func main() {

	age := 18

	// You can also use the following syntax:
	if age >= 18 {
		println("You can vote!")
	} else if age >= 16 {
		println("You can drive!")
	} else {
		println("You can't vote or drive!")
	}

	switch age {
	case 16:
		println("You can drive!")
		break
	case 18:
		println("You can vote!")
		break
	case 21:
		println("You can drink!")
		break
	default:
		println("You can't do anything!")
		break
	}
}
