package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	var arr = []int{-2, -4, 3, 6}

	res := canReorderDoubled(arr)
	fmt.Println(res)
}

func canReorderDoubled(arr []int) bool {
	value := map[int]int{}
	for i := 0; i < len(arr); i++ {
		value[arr[i]]++
	}

	// 按照数的绝对值排序
	sort.Slice(arr, func(a, b int) bool {
		return int(math.Abs(float64(arr[a]))) < int(math.Abs(float64(arr[b])))
	})

	cnt := 0

	for i := 0; i < len(arr) && cnt < len(arr)/2; i++ {
		//这个数已经被使用了
		if value[arr[i]] == 0 {
			continue
		}
		value[arr[i]]--

		if value[arr[i]*2] == 0 {
			return false
		} else {
			value[arr[i]*2]--
		}
		cnt++
	}
	return true
}
