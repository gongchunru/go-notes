package extention

import (
	"fmt"
	"testing"
	"time"
)

type items struct {
	id string
}

//var ch = make(chan string, 10) // 创建大小为 10 的缓冲信道

func TestName(t *testing.T) {
	ch := make(chan string)
	for i := 0; i < 3; i++ {
		go download("a.com/"+string(rune(i+'0')), ch)

	}

	for i := 0; i < 3; i++ {
		msg := <-ch // 等待信道返回消息。
		fmt.Println("finish", msg)
	}

	fmt.Println("Done!")

}

func download(url string, ch chan string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	ch <- url // 将 url 发送给信道
}
