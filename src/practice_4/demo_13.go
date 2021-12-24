package main

import (
	"encoding/json"
	"fmt"
)

// map 初始化
func main() {

	res := make(map[string]interface{})
	res["code"] = 200
	res["message"] = "success"
	res["data"] = map[string]interface{}{
		"username": "Tom",
		"age":      "age",
		"hobby":    []string{"reading", "Climbing"},
	}
	fmt.Println("res", res)

	//序列化
	marshal, err := json.Marshal(res)
	if err != nil {
		fmt.Println("json marshal error:", err)
	}
	fmt.Println("")
	fmt.Println("marshal:", string(marshal))

	//反序列化
	res2 := make(map[string]interface{})
	err = json.Unmarshal(marshal, &res2)
	if err != nil {
		fmt.Println("json unmarshal error:", err)
	}
	fmt.Println("map data:", res2)

}
