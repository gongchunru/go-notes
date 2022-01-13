package main

import "fmt"

// 切片是一种动态数组，比数组操作灵活，长度不是固定的，可以进行追加和删除。

func main() {
	// nil 切片
	var sli_1 []int
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli_1), cap(sli_1), sli_1)

	// 空切片
	var sli_2 = []int{}
	fmt.Printf("len=%d cap=%d slice=%v \n", len(sli_2), cap(sli_2), sli_2)

	var sli_3 = []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slicen=%v\n", len(sli_3), cap(sli_3), sli_3)

	sli_4 := []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d cap=%d slicen=%v\n", len(sli_4), cap(sli_4), sli_4)

	var sli_5 = make([]int, 5, 8)
	fmt.Printf("len=%d cap=%d slicen=%v\n", len(sli_5), cap(sli_5), sli_5)

	sli_6 := make([]int, 5, 9)
	fmt.Printf("len=%d cap=%d slicen=%v\n", len(sli_6), cap(sli_6), sli_6)
}
