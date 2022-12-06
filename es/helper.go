package es

import (
	"bytes"
	"io"
)

// StrToIoReader 将字符串转换为 io.Reader
func StrToIoReader(s string) io.Reader {
	return bytes.NewBuffer([]byte(s))
}
