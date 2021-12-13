package main

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

//https://go101.org/article/unsafe.html

//https://frankma.netlify.app/posts/database/boltdb/#%E5%86%85%E5%AD%98%E5%AF%B9%E9%BD%90
func TestAlign(t *testing.T) {
	var x struct {
		a int64
		b bool
		c string
	}

	fmt.Println(unsafe.Sizeof(x.a), unsafe.Sizeof(x.b), unsafe.Sizeof(x.c), unsafe.Sizeof(x)) // 8 1 16 32

	fmt.Println(unsafe.Alignof(x.a)) // 8
	fmt.Println(unsafe.Alignof(x.b)) // 1
	fmt.Println(unsafe.Alignof(x.c)) // 8

	fmt.Println(unsafe.Offsetof(x.a)) // 0
	fmt.Println(unsafe.Offsetof(x.b)) // 8
	fmt.Println(unsafe.Offsetof(x.c)) // 16
}

//reference: https://github.com/valyala/fasthttp/blob/master/bytesconv.go#L328
func TestByteString(t *testing.T) {
	var b2s = func(bs []byte) string {
		return *(*string)(unsafe.Pointer(&bs))
	}
	var s2b = func(s string) (b []byte) {
		bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
		sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
		bh.Data = sh.Data
		bh.Cap = sh.Len
		bh.Len = sh.Len
		return b
	}
	fmt.Println(b2s([]byte{'b', 'y', 't', 'e'}))
	fmt.Println(s2b("byte"))
}
