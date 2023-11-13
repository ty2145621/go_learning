package internal

import (
	"reflect"
	"unsafe"
)

// String2Bytes no copy
// 看下这个 https://segmentfault.com/a/1190000040841232
func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	var b []byte
	pBytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pBytes.Data = sh.Data
	pBytes.Len = sh.Len
	pBytes.Cap = sh.Len

	return b
}

// String2BytesWithoutCap no copy, 返回值cap是错误的
func String2BytesWithoutCap(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// Bytes2String no copy
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
