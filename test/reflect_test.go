package main

import (
	"fmt"
	"reflect"
	"testing"
)

// 接口 -> 反射对象
// reflect.TypeOf reflect.ValueOf
// Type Value 的 kind方法 返回基础类型
func TestI2O(t *testing.T) {
	type MInt int
	var a MInt = 1

	av := reflect.ValueOf(a)
	ak := av.Kind()
	fmt.Printf("ValueOf: %+v, Kind: %+v, Is int %v\n", av, ak, ak == reflect.Int)

	at := reflect.TypeOf(a)
	atk := at.Kind()
	fmt.Printf("TypeOf: %+v, Kind: %+v\n", at, atk)
}

// 反射对象 -> 接口
// ValueOf的逆过程
func TestO2I(t *testing.T) {
	v := 1.2
	vf := reflect.ValueOf(v)
	f, ok := vf.Interface().(float64)
	fmt.Println(f, ok)
}

// 若要修改反射对象，值必须可设置，即传入ValueOf的是指针
func TestSet(t *testing.T) {
	a := 1.1
	fmt.Println("settability of a: ", reflect.ValueOf(a).CanSet())

	b := 3.4
	fmt.Println("settability of b:", reflect.ValueOf(&b).Elem().CanSet())

	// 通过elem返回指向value的指针
	fmt.Println("before set b: ", b)
	elem := reflect.ValueOf(&b).Elem()
	elem.SetFloat(2.2)
	fmt.Println("after set b: ", b)
}

// 通过反射修改一个struct
func TestSetStruct(t *testing.T) {
	type S struct {
		A int
		B string
	}
	model := S{23, "F"}

	s := reflect.ValueOf(&model).Elem()

	sType := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, sType.Field(i).Name, f.Type(), f.Interface())
	}

	if s.Field(0).CanSet() {
		if s.Field(0).Type().Kind() == reflect.Int {
			s.Field(0).SetInt(111)
		}
	}

	fmt.Printf("modify struct: %+v\n", s)
}

// 通过反射修改一个方法，泛型的装饰器
func Decorator(decoPtr, fn interface{}) (err error) {
	decoratedFunc := reflect.ValueOf(decoPtr).Elem()
	targetFunc := reflect.ValueOf(fn)

	wrapFunc := func(in []reflect.Value) (out []reflect.Value) {
		fmt.Println("before")
		out = targetFunc.Call(in)
		fmt.Println("after")
		return
	}
	v := reflect.MakeFunc(targetFunc.Type(), wrapFunc)

	decoratedFunc.Set(v)
	return
}

func fdd(a, b int) int {
	fmt.Printf("fdd %d + %d\n", a, b)
	return a + b
}

func TestDecorator(t *testing.T) {
	var myFn func(int, int) int
	Decorator(&myFn, fdd)
	myFn(1, 2)
}

// 参考 https://github.com/williamhng/The-Laws-of-Reflection
// https://github.com/astaxie/gopkg/blob/master/reflect/MakeFunc.md
