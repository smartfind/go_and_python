package main

import "fmt"

func main() {
	//变量的定义
	//1、基础变量定义
	//var i int
	//i = 100
	//fmt.Println(i)
	//var i int = 100
	//fmt.Println(i)

	//2、根据值自行判断变量类型（类型推断）
	//var i = 100
	//fmt.Println(i)

	//go语言可省略var
	//i := 100
	//fmt.Println(i)
	//var a int = 100
	//var b = 100
	//c := 100
	//fmt.Println(a, b, c)

	//	多变量赋值
	//var a, b, c int
	//a, b, c = 10, 20, 30
	//fmt.Println(a, b, c)
	//var a, b, c int = 10, 20, 30
	//fmt.Println(a, b, c)

	//	集合类型,多变量不同类型
	//var (
	//	a    int
	//	name string
	//)
	//a = 100
	//name = "xiaoming"
	//fmt.Println(a, name)
	var i int = 100
	i = 10
	fmt.Println(i)
	//	匿名变量，变量一旦被定义，不使用会报错，匿名变量解决该问题 _

	// 常量
	//const pi = 3.1415926
	//r := 2.0
	//c := 2 * pi * r
	//fmt.Println(c)
	//	多个常量一起定义
	const (
		Unknown = 0
		female  = 1
		age     = 20
	)
	fmt.Println(Unknown, female, age)
	//	常量组如不指定类型和初始值，该类型和值与上一行类型一致
	//  常量的数据类型值可以是布尔、数字和字符串
	//  没有被使用的常量在编译时不会报错，与变量不同
	const (
		w int = 20
		x
		y = "abc"
		z
	)
	fmt.Println(w, x, y, z)
	fmt.Println(w, x, y)
	// const常量的iota，常量计数器
	// iota 该常量的值等于上一个常量的表达式
	// 不同的const常量组互不干扰
	const (
		a = iota // 计数器
		b        // 行
		c
	)
	fmt.Println(a, b, c)
	//	没有表达式的常量定义复用上一行的表达式
	// 从第一行开始，iota逐行加1
	const (
		d = iota
		e = 100
		f
		g = iota
	)
	fmt.Println(d, e, f, g)

	const (
		// iota代表的是行数
		h = iota
		j = 200
		k
		l, m = iota, iota
		n    = iota
	)
	fmt.Println(h, j, k, l, m, n)
}
