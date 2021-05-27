package main

import (
	"errors"
	"fmt"
)

//  相比其他静态语言，go语言的函数有很多亮点
// 函数的几个要素：1、函数名；2、参数；3、返回值
//  1、函数的第一种定义
func add(a int, b int) int {
	var sum int
	sum = a + b
	return sum
}

//  2、函数的第二种定义方式
func add2(a, b int) (sum int) {
	sum = a + b
	return sum
}

//  3、函数的第三种定义方式
func add3(a, b int) (sum int) {
	sum = a + b
	return // 默认返回预定义返回的参数值
}

//  4、函数的第四种定义方式
//  被除数等于0，要返回多个值，一个特性
func div(a, b int) (int, error) {
	var err error
	var result int
	if b == 0 {
		err = errors.New("被除数不能为0")
	} else {
		result = a / b
	}

	return result, err
}

func div2(a, b int) (result int, err error) {
	if b == 0 {
		err = errors.New("被除数不能为0")
	} else {
		result = a / b
	}

	return result, err
}

func main() {
	result, err := div(12, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
