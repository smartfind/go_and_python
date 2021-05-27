package main

import "fmt"

// 全局变量只能用var定义
// a := 10 是错误的
// 全局变量为被引用编译会通过
var a = 100

func main() {
	//	变量的作用域
	//	局部变量
	c := 100
	fmt.Println(a, c)
	//	在go语言中字符和字符串不是一种类型，字符类型是单引号，字符串时是双引号
	fmt.Printf("%c\n", 97)
	fmt.Printf("%d\n", 'a')
	fmt.Printf("%d\n", '河')
	fmt.Printf("%T", "河")
}
