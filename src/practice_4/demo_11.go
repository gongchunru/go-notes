package main

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func main() {
	var res Result
	res.Code = 200
	res.Message = "success"
	toJson(&res)

	setData(&res)
	toJson(&res)
}

func setData(res *Result) {
	res.Code = 500
	res.Message = "fail"
}

func toJson(res *Result) {
	marshal, err := json.Marshal(res)
	if err != nil {
		fmt.Println("json marshal error:", err)
	}
	fmt.Println("res:", string(marshal))

}
