package main

import (
	"fmt"
	"time"
)

var (
	cst *time.Location
)

const CSTLayout = "2006-01-02 15:04:05"

func init() {
	var err error
	if cst, err = time.LoadLocation("Asia/Shanghai"); err != nil {
		panic(err)
	}
}

func RFC3339ToCSTLayout(value string) (string, error) {
	ts, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return "", err
	}
	return ts.In(cst).Format(CSTLayout), nil
}

func main() {

	RFC3339Str := "2020-11-08T08:18:46+08:00"
	cst, err := RFC3339ToCSTLayout(RFC3339Str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cst)

}
