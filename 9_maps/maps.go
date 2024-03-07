package main

import "fmt"

func main() {

	StudentAge := make(map[string]int)
	StudentAge["John"] = 32
	StudentAge["Bob"] = 31
	StudentAge["Marry"] = 28
	StudentAge["Todd"] = 33

	fmt.Println(StudentAge)
	fmt.Println(len(StudentAge))

	superhero := map[string]map[string]string{
		"Superman": map[string]string{
			"RealName": "Clark Kent",
			"City":     "Metropolis",
		},
		"Batman": map[string]string{
			"RealName": "Bruce Wayne",
			"City":     "Gotham",
		},
	}

	if temp, hero := superhero["Superman"]; hero {
		fmt.Println(temp["RealName"], temp["City"])
	}
}
