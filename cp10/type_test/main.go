package main

import "fmt"

func main() {
	//	go语言中关键词type
	//	1、给一个类型定义别名，例如rune，byte
	//  为什么会有byte，就是为了强调我们现在处理的对象是字节类型
	//  2、这种别名实际上还是为了代码的可读性，这个实际上本质上仍然是uint8，无非是在代码编码阶段可读性强而已
	type mybyte = byte
	var a mybyte
	fmt.Printf("%T\n", a)
	//	第二种，基于一个已有的类型定义一个新的类型
	type myint int // 没有等号
	var i myint
	fmt.Printf("%T\n", i)

	//	第三种，定义结构体
	type Course struct {
	}
	//  第四种，定义接口
	//  第五中，定义函数别名

}
