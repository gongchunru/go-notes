package main

import "fmt"

/**
在声明时不会立刻去执行，而是在函数 return 后去执行的。

它的主要应用场景有异常处理、记录日志、清理数据、释放资源 等等。

这篇文章不是分享 defer 的应用场景，而是分享使用 defer 需要注意的点。
*/
func main() {
	//x := 1
	//y := 2
	//defer calc("A", x, calc("B", x, y))
	//x = 3
	//defer calc("C", x, calc("D", x, y))
	//y = 4

	// 针对的拆解一下

	x := 1
	y := 2
	tmp1 := calc("B", x, y)
	defer calc("A", x, tmp1)
	x = 3
	tmp2 := calc("D", x, y)
	defer calc("C", x, tmp2)
	y = 4

}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
