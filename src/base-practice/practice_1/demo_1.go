package main

import "fmt"

func main() {

	/**
	* 数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成，
	 一旦声明了，数组的长度就固定了，不能动态变化。
	*/
	// 一维数组
	var arr_1 [5]int
	fmt.Println(arr_1)

	var arr_2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr_2)

	arr_3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr_3)

	// 填值不足的，后面会补上0
	arr_4 := [6]int{2, 3, 4, 5, 6}
	fmt.Println(arr_4)

	// 指定位置的数组
	arr_5 := [5]int{0: 3, 3: 5, 2: 3}
	fmt.Println(arr_5)

	//
	arr_6 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr_6)

	// 二维数组
	var arr_8 = [3][5]int{{1, 2, 3, 4, 5}, {2, 3, 4}, {5, 4, 3}}
	fmt.Println(arr_8)

	arr_9 := [3][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 5, 3}, {2, 3, 4, 5, 6}}
	fmt.Println(arr_9)

	arr_10 := [...][3]int{{2, 3}, {23, 4}}
	fmt.Println(arr_10)
}
