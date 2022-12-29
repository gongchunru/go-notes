package main

import "fmt"

func main() {
	fmt.Println(minimumMoves("XXOX"))
}

func minimumMoves(s string) int {
	n := len(s)
	res := 0
	for i := 0; i < n; i++ {
		if s[i] == 'X' {
			i += 2
			res++
		}
	}
	return res
}
