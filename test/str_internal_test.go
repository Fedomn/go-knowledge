package main

//code point

//unicode 字符集是对世界上多种语言字符的通用编码 也叫万国码

//UTF-8（8-bit Unicode Transformation Format）是一种针对Unicode的可变长度字符编码
//它可以用一至四个字节对Unicode字符集中的所有有效编码点进行编码

//golang中的字符串结构是 stringStruct

//rune: unicode 字符集中，每一个字符都有一个对应的编号，我们称这个编号为 code point(码位)
//而 Go 中的rune 类型就代表一个字符的 code point
//字符集只是将每个字符给了一个唯一的编码而已
//从rune的定义 type rune = int32 也可以看出它最大4个字节，也就对应utf-8
