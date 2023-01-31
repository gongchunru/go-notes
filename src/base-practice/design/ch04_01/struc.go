package main

import "fmt"

type MyStruct struct {
	i int
}

func myFunction(a MyStruct, b *MyStruct) {
	a.i = 31
	b.i = 41
	fmt.Printf("in my_function - a =(%d,%p) b=(%v, %p)\n", a, &a, b, &b)
}

func main() {
	a := MyStruct{30}
	b := MyStruct{40}
	fmt.Printf("before - a =(%d,%p) b=(%v, %p)\n", a, &a, b, &b)
	myFunction(a, &b)
	fmt.Printf("after -  a=(%d,%p) b = (%v,%p)\n", a, &a, b, &b)
}
