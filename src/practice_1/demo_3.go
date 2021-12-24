package main

import "fmt"

func main() {
	//在函数中传递的时候是传递的值
	var arr = [5]int{1, 2, 3, 4, 5}
	modifyArray(&arr)
	fmt.Println(arr)

	var arr_2 [5]int = arr

	fmt.Println(arr_2)
}

func modifyArray(a *[5]int) {
	a[2] = 20
}
