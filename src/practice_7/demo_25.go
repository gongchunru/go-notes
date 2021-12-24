package main

import "fmt"

// defer 函数定义的顺序 与 实际执的行顺序是相反的，也就是最先声明的最后才执行。
func main() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	fmt.Println("main")
}
