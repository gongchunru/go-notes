package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 1))

	//fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))
	//z := strings.SplitN("a,b,c", ",", 0)
	//fmt.Println(z)
}
