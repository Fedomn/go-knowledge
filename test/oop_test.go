package main

import (
	"fmt"
	"testing"
	"strconv"
)

//值类型: 基础类型 array struct pointer
//引用类型: slice map channel interface
func TestValueTypeAssign(t *testing.T) {
	//slice比array少了长度定义
	testArray := [3]int{1, 2, 3}
	assignArray := testArray
	assignArray[1] = 10
	fmt.Println(testArray, assignArray)

	type myStruct struct {
		data int
	}
	testStruct := myStruct{1}
	assignStruct := testStruct
	assignStruct.data = 2
	fmt.Println(testStruct, assignStruct)

	var testPointer *int
	data1, data2 := 1, 2
	testPointer = &data1
	assignPointer := testPointer
	assignPointer = &data2
	fmt.Println(*testPointer, *assignPointer)
}

type Human struct {
	name string
	age  int
}

type Skills []string

type Student struct {
	Human
	Skills
	speciality string
}

func TestStruct(t *testing.T) {
	frank := Student{Human: Human{"frank", 20}, Skills: []string{"skill1"}, speciality: "special"}
	fmt.Println(frank.name, frank.age, frank.Skills, frank.speciality)
	frank.Skills = append(frank.Skills, "skill2")
	fmt.Println(frank.Skills)
}

//A method is a function with an implicit first argument, called a receiver
//Receiver还可以是指针, 两者的差别在于
//指针作为Receiver会对实例对象的内容发生操作
//而普通类型作为Receiver仅仅是以副本作为操作对象,并不对原实例对象发生操作
func (h Human) SayHello() {
	fmt.Printf("Hello, I am %s\n", h.name)
}

func (s Student) SayHello() {
	fmt.Printf("Hello, I am student %s\n", s.name)
}

//如果一个method的receiver是*T 你可以在一个T类型的实例变量V上面调用这个method 而不需要&V去调用这个method
//如果一个method的receiver是T  你可以在一个*T类型的变量P上面调用这个method 而不需要*P去调用这个method
func (s *Student) changeName(name string) {
	s.name = name
}

func TestMethodOverride(t *testing.T) {
	frank := Student{Human: Human{"frank", 20}}
	frank.SayHello()
	frank.changeName("Fedomn")
	frank.SayHello()
}

func TestVoidInterface(t *testing.T) {
	var any interface{}
	i := 21
	s := "string"
	any = i
	fmt.Println(any)
	any = s
	fmt.Println(any)
}

func (h Human) String() string {
	return "I am " + h.name + ",my age is " + strconv.Itoa(h.age)
}

func TestMethodString(t *testing.T) {
	frank := Student{Human: Human{"frank", 20}}
	fmt.Println(frank)
}

type Element interface{}
type List[] Element

func TestInterfaceCommaOk(t *testing.T) {
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Human{"frank", 20}

	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Human:
			fmt.Printf("list[%d] is a Human and its value is %s\n", index, value)
		default:
			fmt.Println("list[%d] is of a different type", index)
		}
	}
}
