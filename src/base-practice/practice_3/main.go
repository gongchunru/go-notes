package main

import "fmt"

/*
 int8  1      -2^7 ~ 2^7-1
 int16 2
 int32 4
 int64 8

uint8
uint16
uint32
uint64

*/

func main() {
	var a = 100
	var intb int32

	intb = int32(a)

	fmt.Printf("a---%t", a, intb)
}
