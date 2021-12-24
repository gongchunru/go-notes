package main

import (
	"fmt"
	"go-notes/src/base-practice/practice_9"
)

func main() {
	name := "Tom"
	s, err := study.New(name)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s.Write("English"))
	fmt.Println(s.Listen("English"))
	fmt.Println(s.Read("English"))
	fmt.Println(s.Speak("English"))

}
