package main

import (
	"fmt"
	"go-notes/src/base-practice/practice_11"
)

func main() {
	friends, err := friend.Find("附近的人",
		friend.WithSex(1),
		friend.WithAge(30),
		friend.WithHeight(180),
		friend.WithWeight(60))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(friends)
}
