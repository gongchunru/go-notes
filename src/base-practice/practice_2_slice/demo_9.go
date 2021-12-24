package main

import "fmt"

// 结构体是将零个或多个任意类型的变量，组合在一起的聚合数据类型，也可以看做是数据的集合。

type Person struct {
	Name string
	Age  int
}

func main() {
	var p1 Person
	p1.Name = "Tom"
	p1.Age = 30
	fmt.Print("p1 = ", p1)

	var p2 = Person{Name: "Burke", Age: 30}
	fmt.Println("p2 = ", p2)

	p3 := Person{Name: "Aaron", Age: 32}
	fmt.Println("p3=", p3)

	// 匿名结构体
	P4 := struct {
		Name string
		Age  int
	}{Name: "匿名", Age: 33}

	fmt.Println("p4 = ", P4)
}
