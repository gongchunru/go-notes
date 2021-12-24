package main

import "fmt"

func main() {
	var a int
	b := "test"
	var c = 100
	fmt.Print(a, b, c)

	var i = 100
	var j = 200
	fmt.Println(i, j)
	i, j = j, i
	fmt.Println(i, j)
	const (
		iotaVariables1 = iota
		iotaVariables2 = iota
		iotaVariables3 = iota
	)
	fmt.Println(iotaVariables1, iotaVariables2, iotaVariables3)

}
