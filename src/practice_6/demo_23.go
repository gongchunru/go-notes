package main

import (
	"fmt"
	"time"
)

func main() {

	//fmt.Println("main start")
	//ch := make(chan string)
	//// 定义的是一个无缓冲的 chan，赋值后就陷入了阻塞。
	//ch <- "a" // 入 chan
	//go func() {
	//	val := <-ch // 出 chan
	//	fmt.Println(val)
	//}()
	//fmt.Println("main end")

	fmt.Println("main start")
	ch := make(chan string, 1)
	ch <- "1" // 入 chan
	go func() {
		val := <-ch // 出chan
		fmt.Println(val)
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("main end")

	demo2()
	fmt.Println("--------demo3---------")

	demo3()
	fmt.Println("---------demo4--------")

	demo4()
	fmt.Println("---------demo5---------")

	demo5()

}

func demo2() {
	fmt.Println("main start")
	ch := make(chan string)
	go func() {
		ch <- "a"
	}()
	go func() {
		val := <-ch // 出chan
		fmt.Println(val)
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("main end")
}

func demo3() {
	fmt.Println("main start")

	ch := make(chan string)
	go func() {
		ch <- "a"
	}()

	go func() {
		val := <-ch
		fmt.Println(val)
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}

// 带缓冲的通道，如果长度等于缓冲长度时，再进就会阻塞。
func demo4() {
	fmt.Println("main start")

	ch := make(chan string, 3)
	go producer(ch)

	time.Sleep(1 * time.Second)
	fmt.Println("main end")
}

func demo5() {
	fmt.Println("main start")
	ch := make(chan string, 3)
	go producer(ch)
	go consumer(ch)
	time.Sleep(2 * time.Second)
	fmt.Println("main end")
}

func producer(ch chan string) {
	fmt.Println("producer start")
	ch <- "a"
	ch <- "b"
	ch <- "c"
	ch <- "d"
	fmt.Println("producer end")
}

func consumer(ch chan string) {
	for {
		msg := <-ch
		fmt.Println(msg)
	}
}
