package main

import "fmt"

func swap(a int, b int) {
	//	用于交换a和b
	c := a
	a = b
	b = c
}

func swap2(d *int, e *int) {
	//	用于交换d和e
	f := *d
	*d = *e
	*e = f
}

func main() {
	//	什么是指针
	a := 10
	b := 20
	swap(a, b)
	fmt.Println(a, b)
	//	指针-对于内存来说，每一个字节其实都有地址-可以通过16进制打印出来
	fmt.Printf("%p\n", &a) // 打印变量内存地址
	//	现在有一种特殊的变量类型，这个变量只能保存地址值
	var int_pointer *int // 加*变量里面只能保存地址类型这种值
	int_pointer = &a
	fmt.Println(int_pointer)
	//	如果要修改指针指向的变量的值，用法比较特殊
	*int_pointer = 30
	fmt.Println(a)
	fmt.Println(*int_pointer)
	//	如何定义指针变量，必须定义好指针中内存中变量类型，否则通过指针去取变量的值的时候不知道应该取多大的连续内存的空间
	fmt.Printf("int_pointer所指向的内存空间的地址是：%p，内存中的值是：%d\n", int_pointer, *int_pointer)
	swap2(&a, &b)
	fmt.Println(a, b)
	//	go语言中数组是值传递，对于这种一般采用切片
	//	指针还可以指向数组，指向数组的指针，数组是值类型

	arr := [3]int{1, 2, 3}
	var ip *[3]int = &arr
	fmt.Println(&ip)
	fmt.Println(&arr)
	//	指针数组
	var ptrs [3]*int //创建能够存放三个指针变量的数组
	//  很多时候都是函数参数的时候指明的类型
	//  指针的默认值是nil
	fmt.Println(ptrs)
	//	go语言没有屏蔽指针，但是go语言在指针上做了大量限制，安全性高很多，相对于c和c++灵活性就降低了
	//	指针变量设计到两个符号&和*

}
