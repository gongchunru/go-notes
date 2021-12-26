package main

import (
	"gee"
	"net/http"
)

/*
panic 会导致程序被中止，但是在退出前，会先处理完当前协程上已经defer 的任务，执行完成后再退出。效果类似于 java 语言的 try...catch。

Go 语言还提供了 recover 函数，可以避免因为 panic 发生而导致整个程序终止，recover 函数只在 defer 中生效。
recover 捕获了 panic，程序正常结束。当 panic 被触发时，控制权就被交给了 defer



curl "http://localhost:9999/panic"

curl "http://localhost:9999"

*/

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "hello Jacky\n")
	})

	r.GET("/panic", func(c *gee.Context) {
		names := []string{"Jacky"}
		c.String(http.StatusOK, names[100])
	})

	r.Run("localhost:9999")
}
