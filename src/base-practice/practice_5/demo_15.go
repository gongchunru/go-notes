package main

import (
	"fmt"
)

func main() {
	person := [3]string{"Tom", "Aaron", "John"}
	fmt.Printf("len=%d,cap=%d,array=%v\n", len(person), cap(person), person)

	fmt.Println("")

	// 循环
	for k, v := range person {
		fmt.Printf("person[%d]:%s\n", k, v)
	}
	fmt.Println("")

	for i := 0; i < len(person); i++ {
		fmt.Printf("person[%d]:%s\n", i, person[i])
	}
	fmt.Println("")

	for _, name := range person {
		fmt.Println("name:", name)
	}

	//
	for k := range person {
		fmt.Println("k:", k)
	}
}
