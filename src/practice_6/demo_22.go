package main

import (
	"fmt"
	"time"
)

/**
chan 可以理解为队列，遵循先进先出的原则
go关键字,go关键字后加一个函数，就可以创建一个线程，函数可以为已经写好的函数，也可以是匿名函数

并发机制，不会输出 goroutine
*/
func main() {

	fmt.Println("main start")

	go func() {
		fmt.Println("goroutine")
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("main end")

	//生命不带缓冲的通道 进和出都会阻塞。
	//ch1 := make(chan string)
	//
	//ch1 <- "a"
	//val, ok := <-ch1
	//val := <-ch1
	//close(ch1)

	//close 以后不能再写入，写入会出现 panic
	//重复 close 会出现 panic
	//只读的 chan 不能 close
	//close 以后还可以读取数据

	////声明带10个缓冲的通道  进一次长度 +1，出一次长度 -1，如果长度等于缓冲长度时，再进就会阻塞。
	//ch2 := make(chan string, 10)
	//
	////声明只读通道
	//ch3 := make(<-chan string)
	//
	//// 声明只写通道
	//ch4 := make(chan<- string)

}
