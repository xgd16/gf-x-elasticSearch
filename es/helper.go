package es

import (
	"bytes"
	"github.com/gogf/gf/v2/encoding/gjson"
	"io"
)

// StrToIoReader 将字符串转换为 io.Reader
func StrToIoReader(s string) io.Reader {
	return bytes.NewBuffer([]byte(s))
}

// AnyToIoReader 任何类型转换为 io.Reader
func AnyToIoReader(d any) io.Reader {
	return bytes.NewBuffer(gjson.MustEncode(d))
}
