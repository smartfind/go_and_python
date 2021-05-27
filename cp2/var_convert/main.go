package main

import (
	"fmt"
	"strconv"
)

func main() {
	//	1、基本类型转换
	//a := int(3.0)
	//fmt.Println(a)
	//	在go语言中不支持变量间的隐式类型转换
	//	1、变量间类型转换
	//var b int = 5.0 // 5.0是常量，支持常量和变量的类型转换
	//fmt.Println(b)
	//c := 5.0
	//fmt.Printf("%T\n", c)
	////var d int = c //不支持
	//var d int = int(c)
	//fmt.Println(d)
	//	int转字符串，itoa
	var a int64 = 56
	var b int32 = int32(a)
	fmt.Println(b)
	fmt.Printf("%T\n", strconv.Itoa(int(a)))
	//	字符串转int，atoi
	fmt.Println(strconv.Atoi("20"))
	data, _ := strconv.Atoi("12")
	fmt.Println(data)
	//	parse类函数
	c, _ := strconv.ParseBool("true")
	fmt.Println(c)
	d, _ := strconv.ParseFloat("2.574357476", 64)
	fmt.Println(d)
	fmt.Printf("%T\n", d)
	f, _ := strconv.ParseInt("12", 10, 64)
	fmt.Println(f)
	g, _ := strconv.ParseUint("12", 10, 64)
	fmt.Println(g)
	// 其他类型转字符串
	h := strconv.FormatBool(true)
	fmt.Println(h)
	i := strconv.FormatFloat(2.646456, 'E', -1, 64)
	fmt.Println(i)
	j := strconv.FormatUint(42, 16)
	fmt.Println(j)
}
