package day01_codec

import (
	codec "geerpc01/codec"
)

const MagicNumber = 0x3bef5c

type Option struct {
	MagicNumber int
	CodecType   codec.Codec
}

var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodecType:   codec.GobType,
}
