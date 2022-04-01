package main

import "fmt"

func main() {
	numbers := selfDividingNumbers(1, 22)
	fmt.Println(numbers)
}

func selfDividingNumbers(left int, right int) (ans []int) {
	check := func(num int) bool {
		for x := num; x > 0; x /= 10 {

			t := x % 10
			// 包含0 或者不能整除
			if t == 0 || num%t != 0 {
				return false
			}
		}
		return true
	}
	for i := left; i <= right; i++ {
		if check(i) {
			ans = append(ans, i)
		}
	}
	return ans
}
