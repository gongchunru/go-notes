package codec

import "io"

type Header struct {
	ServiceMethod string
	Seq           uint64
	Error         string
}

// Codec 抽象出对消息体进行编解码的接口 Codec，抽象出接口是为了实现不同的 Codec 实例：
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

type NewCodeFunc func(io.ReadWriteCloser) Codec

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json"
)

var NewCodecFuncMap map[Type]NewCodeFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodeFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
