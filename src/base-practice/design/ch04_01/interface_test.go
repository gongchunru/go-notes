package main

import "fmt"

type error interface {
	Error() string
}

type RPCError struct {
	Code    int64
	Message string
}

func (e *RPCError) Error() string {
	return fmt.Sprintf("%s, code=%d", e.Message, e.Code)
}

func test_run() {
	var rpcErr error = NewRPCError(400, "test error")
	err := AsErr(rpcErr)
	println(err)
}

func NewRPCError(code int64, msg string) error {
	return &RPCError{
		Code:    code,
		Message: msg,
	}
}

func AsErr(err error) error {
	return err
}
