package main

import (
	"fmt"
	"os"
)

// 闭包获取变量相当于引用传递，而非值传递。

/**

defer 调用的函数，参数的值在 defer 定义时就确定了，看下代码

defer fmt.Println(a + b)，在这时，参数的值已经确定了。

而 defer 函数内部所使用的变量的值需要在这个函数运行时才确定，看下代码

defer func() { fmt.Println(a + b) }()，a 和 b 的值在函数运行时，才能确定。

*/
func main() {
	//var a = 1
	//var b = 2
	//defer func() {
	//	fmt.Println(a + b)
	//}()
	//a = 2
	fmt.Println("main")

	fmt.Println("----t1----")
	println(t1())

	fmt.Println("----t2----")
	println(t2())

	fmt.Println("----t3----")
	fmt.Println(t3(1))

	fmt.Println("----t4----")
	fmt.Println(t4(1))
	fmt.Println("----t5----")

	defer fmt.Println("1")

	// 当os.Exit()方法退出程序时，defer不会被执行。
	os.Exit(0)

}

func t1() int {
	fmt.Println("----t1----")
	a := 1
	defer func() {
		a++
	}()
	return a
}

func t2() (a int) {
	fmt.Println("----t2----")
	defer func() {
		a++
	}()
	return 1
}

func t3(b int) int {
	fmt.Println("----t3----")
	a := 1
	defer func() {
		a++
	}()
	return 1
}

func t4(a int) int {
	fmt.Println("----t4----")
	defer func(a int) {
		a++
	}(a)
	return 1
}
