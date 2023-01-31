package main

import "fmt"

/*
Go 语言的接口类型不是任意类型
*/
type TestStruct struct {
}

func NioOrNot(v interface{}) bool {
	return v == nil
}

/*
下面的代码在 main 函数中初始化了一个 *TestStruct 类型的变量，由于指针的零值是 nil，所以变量 s 在初始化之后也是 nil
*/
func test_001() {
	var s *TestStruct
	fmt.Println(s == nil)

	// 调用 NilOrNot 函数时发生了隐式的类型转换
	// 在类型转换时，*TestStruct 类型会转换成 interface{} 类型，转换后的变量不仅包含转换前的变量，还包含变量的类型信息 TestStruct，所以转换后的变量与 nil 不相等。
	fmt.Println(NioOrNot(s))
}
