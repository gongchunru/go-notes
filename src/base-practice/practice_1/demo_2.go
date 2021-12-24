package main

import "fmt"

// 数组是值传递
func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	// 长度不可变
	//arr[5] = 6

	modify(arr)
	fmt.Println(arr)
}

func modify(a [5]int) {
	a[1] = 20
}
