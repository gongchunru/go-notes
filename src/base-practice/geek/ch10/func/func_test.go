package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestReturnMultiValues(t *testing.T) {

}

func TestSpend(t *testing.T) {
	tsSF := timeSpend(slowFunc)
	t.Log(tsSF(10))
}

func timeSpend(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spend:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
}

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}
