package main

import "fmt"

// 闭包
func main() {
	var a = 1
	var b = 2
	defer fmt.Println(a + b)
	a = 2
	fmt.Println("main")
}
