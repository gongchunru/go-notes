package main

import "fmt"

func main() {
	person := map[int]string{
		1: "Tom",
		2: "Aaron",
		3: "John",
	}
	fmt.Println("person:", person)

	delete(person, 2)
	fmt.Println("person:", person)

	person[2] = "Jack"
	person[3] = "Kevin"
	fmt.Println("person:", person)
}
