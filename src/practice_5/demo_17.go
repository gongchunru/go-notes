package main

import "fmt"

// 循环map
func main() {

	person := map[int]string{
		1: "Tom",
		2: "Aaron",
		3: "John",
	}

	fmt.Printf("len=%d map=%v\n", len(person), person)

	fmt.Println("")

	for k, v := range person {
		fmt.Printf("person[%d]:%s\n", k, v)
	}

	fmt.Println("")
	for _, name := range person {
		fmt.Printf("name: %s \n", name)
	}

	fmt.Println("")
	for i := 1; i <= len(person); i++ {
		fmt.Printf("person[%d]:%s\n", i, person[i])
	}

	fmt.Println("")

	for i := range person {
		fmt.Printf("person[%d]:%s\n", i, person[i])
	}
}
