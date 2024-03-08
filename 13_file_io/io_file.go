package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Create("./13_file_io/test.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString("This is a test file")
	if err != nil {
		log.Fatal(err)
		return
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
		return
	}

	stream, err := os.ReadFile("./13_file_io/test.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	string := string(stream)
	fmt.Println(string)

}
