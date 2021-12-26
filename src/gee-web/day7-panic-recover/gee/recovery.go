package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		c.Next()
	}
}

func trace(message string) string {
	var pcs [32]uintptr
	// return program counters of function invocations on the calling goroutine's stack
	// 跳过前3个Caller
	n := runtime.Callers(3, pcs[:])

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")

	for _, pc := range pcs[:n] {
		// 获取对应的函数，
		fn := runtime.FuncForPC(pc)
		// 获取对应的文件名和行号打印出来
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}
