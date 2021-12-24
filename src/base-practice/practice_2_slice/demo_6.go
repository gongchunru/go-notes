package main

import "fmt"

func main() {
	sli := []int{4, 5, 6}
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli), cap(sli), sli)

	sli = append(sli, 7)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli), cap(sli), sli)

	sli = append(sli, 8)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli), cap(sli), sli)

	sli = append(sli, 9)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli), cap(sli), sli)

	// append 时，容量不够需要扩容时，cap 会翻倍。
	sli = append(sli, 10)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(sli), cap(sli), sli)

}
