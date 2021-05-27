package main

import "fmt"

func main() {
	//	算术运算符
	//a := 12
	//b := 22
	//fmt.Println(a / b) // 向上取整
	//a++
	//fmt.Println(a)
	var a bool = true
	var b bool = false
	//  加不加（）均可
	if a && b {
		fmt.Println("两个条件均满足")
	}
	if a || b {
		fmt.Println("有一个条件满足")
	}
	if !(a && b) {
		fmt.Println("反转为真")
	}
}
