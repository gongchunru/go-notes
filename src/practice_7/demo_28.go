package main

import (
	"fmt"
	"time"
)

// defer 只对当前协程有效。

func main() {
	GoA()
	time.Sleep(1 * time.Second)
	fmt.Println("main")

}

func GoA() {
	// 只对当前协程有效
	defer (func() {
		if err := recover(); err != nil {
			fmt.Println("panic:" + fmt.Sprintf("%s", err))
		}
	})()

	go GoB()
}

func GoB() {
	panic("error")
}
