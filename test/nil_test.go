package main

import (
	"fmt"
	"testing"
	"unsafe"
)

// nil定义 src/builtin/builtin.go

// nil is a predeclared identifier representing the zero value for a
// pointer, channel, func, interface, map, or slice type.
// var nil Type // Type must be a pointer, channel, func, interface, map, or slice type

// Type is here for the purposes of documentation only. It is a stand-in
// for any Go type, but represents the same type for any given function
// invocation.
// type Type int

// 上面看出nil是一个值，而不是一个类型
// Go语言规范 任何类型在未初始化时都对应一个零值
// 布尔类型是false，整型是0，字符串是""，而pointer、func、interface、slice、channel和map的零值都是nil。
// struct的零值与其属性有关

// nil没有默认的类型，尽管它是多个类型的零值，必须显式或隐式 指定每个nil的明确类型
func TestCheckNil(t *testing.T) {
	// 明确.
	_ = (*struct{})(nil)
	_ = []int(nil)
	_ = map[int]bool(nil)
	_ = chan string(nil)
	_ = (func())(nil)
	_ = interface{}(nil)

	// 隐式.
	var _ *struct{} = nil
	var _ []int = nil
	var _ map[int]bool = nil
	var _ chan string = nil
	var _ func() = nil
	var _ interface{} = nil
}

// nil不是关键字 可以在代码中定义nil 那么nil就会被隐藏
func TestAssignNil(t *testing.T) {
	nil := 1
	fmt.Println(nil)
	a := 2
	if a == nil {
		fmt.Println("ok")
	}
}

// 不同类型nil的内存地址是一样的
func TestNilAddr(t *testing.T) {
	var m map[int]string
	var ptr *int
	var sl []int
	fmt.Printf("%p\n", m)   //0x0
	fmt.Printf("%p\n", ptr) //0x0
	fmt.Printf("%p\n", sl)  //0x0
}

// nil值的大小 与其 类型non-nil值的大小 相同
// 即：不同零值的nil标识符可能具有不同的大小
func TestNilSize(t *testing.T) {
	var p *struct{}
	fmt.Println(unsafe.Sizeof(p)) // 8

	var s []int
	fmt.Println(unsafe.Sizeof(s)) // 24

	var m map[int]bool = nil
	fmt.Println(unsafe.Sizeof(m)) // 8

	var c chan string = nil
	fmt.Println(unsafe.Sizeof(c)) // 8

	var f func() = nil
	fmt.Println(unsafe.Sizeof(f)) // 8

	var i interface{} = nil
	fmt.Println(unsafe.Sizeof(i)) // 16
}

// Go里 两个值比较，只能在两个值 之间可以隐式转换的情况下 可以比较
// nil值比较并没有脱离上述规则。
func TestNilCmp(t *testing.T) {
	//1、不同类型的nil是不能比较的

	//var m map[int]string
	//var ptr *int
	//fmt.Printf(m == ptr) //invalid operation: m == ptr (mismatched types map[int]string and *int)

	//2、同一类型的两个nil值可能无法比较
	//map、slice、func类型 是不可比较类型。所以比较它们的nil也是非法的。

	//var v1 []int = nil
	//var v2 []int = nil
	//fmt.Println(v1 == v2) //invalid operation: v1 == v2 (operator == not defined on []int)
	//fmt.Println((map[string]int)(nil) == (map[string]int)(nil)) //invalid operation
	//fmt.Println((func())(nil) == (func())(nil)) //invalid operation

	//不可比较的类型的值 是可以与nil 比较的。常见判断初始化
	fmt.Println(map[string]int(nil) == nil) //true
	fmt.Println((func())(nil) == nil)       //true

	//3、两nil值可能不相等
	//接口值将转换为接口值的类型。转换后的接口值具有具体的动态类型, 但其他接口值没有
	fmt.Println((interface{})(nil) == (*int)(nil)) //false
}

// map的key为指针、函数、interface、slice、channel和map，则key可以为nil。
func TestMapNil(t *testing.T) {
	mmap := make(map[*string]int, 4)
	a := "a"
	mmap[&a] = 1
	mmap[nil] = 99
	fmt.Println(mmap) //map[0xc042008220:1 <nil>:99]
}

// 总结
// nil是一个值，而不是一个类型
