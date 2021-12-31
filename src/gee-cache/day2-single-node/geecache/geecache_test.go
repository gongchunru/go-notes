package geecache

import (
	"fmt"
	"testing"
	"time"
)

var set = make(map[int]bool, 0)

func PrintOne(num int) {
	if _, exist := set[num]; !exist {
		fmt.Println(num)
	}
	set[num] = true
}

func TestForLoop(t *testing.T) {
	for i := 0; i < 100; i++ {
		PrintOne(100)
	}
	time.Sleep(time.Second)
}
