package interview

import (
	"fmt"
	"reflect"
	"unsafe"
)

func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func String2Byte2(a string) []byte {
	ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a))
	b := *(*[]byte)(unsafe.Pointer(&ssh))
	return b
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func test2() {
	s := "aaaaa"
	a1 := String2Bytes(s)
	a2 := String2Byte2(s)

	aa1 := (*reflect.SliceHeader)(unsafe.Pointer(&a1))
	aa2 := (*reflect.SliceHeader)(unsafe.Pointer(&a2))

	fmt.Printf("%+v, %+v\n", aa1, aa2)

}

func test1() {
	var s1 []int
	s2 := make([]int, 0)
	s4 := make([]int, 0)

	fmt.Printf("s1 pointer:%+v, s2 pointer:%+v, s4 pointer:%+v, \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s1)), *(*reflect.SliceHeader)(unsafe.Pointer(&s2)), *(*reflect.SliceHeader)(unsafe.Pointer(&s4)))
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s1))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data)
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&s4))).Data)
}
