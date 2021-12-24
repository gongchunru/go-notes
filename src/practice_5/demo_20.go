package main

import "fmt"

/**
默认每个 case 带有 break
case 中可以有多个选项
fallthrough 不跳出，并执行下一个 case
*/
func main() {
	i := 3
	fmt.Printf("当i=%d 时:\n", i)
	switch i {
	case 1:
		fmt.Println("输出 i=", 1)
	case 2:
		fmt.Println("输出 i=", 2)
	case 3:
		fmt.Println("输出 i=", 3)
		fallthrough
	case 4, 5, 6:
		fmt.Println("输出 i=", "4,5,6")
	default:
		fmt.Println("输出 i=", "default")
	}
}
