package main

import (
	"fmt"
	"testing"
	"unsafe"
)

type MyStruct2 struct {
	i int
	j int
}

func myFunction2(ms *MyStruct2) {
	ptr := unsafe.Pointer(ms)
	for i := 0; i < 2; i++ {
		c := (*int)(unsafe.Pointer((uintptr(ptr) + uintptr(i*8))))
		*c += i + 1
		fmt.Printf("[%p] %d \n", c, *c)
	}
}

func TestSpend(t *testing.T) {
	a := &MyStruct2{40, 50}
	myFunction2(a)
	fmt.Printf("[%p] %d\n", a, *a)
}
